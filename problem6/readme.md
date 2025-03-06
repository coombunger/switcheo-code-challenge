# Transaction Broadcaster Service

The transaction broadcaster service consists of the service itself, and the UI which calls the service's HTTP API.

## UI

The UI allows all users to view the list of successful, failed, or pending transactions,
and allows admins to retry failed transactions.

- To get the list of transactions, the UI calls the GET /transaction endpoint.
- To retry failed transactions, the UI calls the POST /transaction/retry endpoint.
- To sign in and out, the UI calls the /auth endpoints.

## Transaction Broadcaster Service

The transaction broadcaster service employs a microservice architecture, with an auth, transaction, and broadcasting service.
Different services can be scaled individually for scalability, and multiple instances for each service ensures fault tolerance.

### API

The service exposes a HTTP API for the UI and other services to broadcast, retry, and view transactions.

All requests are routed through a gateway, which directs auth requests to the auth service and transaction requests to the transaction service.
The gateway is responsible for validating the requests (e.g. auth) and load balancing between different instances of the services.
If the request is invalid, it returns HTTP `4XX`. If there is an internal server error, it returns HTTP `500`. Else, it returns the response from the corresponding service.

### Auth Service

The service uses JWT authentication with RBAC.
All users are authorized for the `GET /transaction` and `POST /transaction/broadcast` endpoints, but only admins are authorized for the `POST /transaction/retry` endpoint.

### Transaction Service

The transaction service handles transaction requests.

This service is also responsible for fault recovery.
When the service restarts, it queries the database for any `PENDING` transactions, and publishes `check and broadcast` messages with their transaction ids to the message queue.

**GET /transaction**

This endpoint takes a `state` query parameter, which specifies which state (`SUCCESS`, `FAILURE`, `PENDING`) to filter the transactions by.

The service queries the database and returns the list of filter queries with HTTP `200`.

**POST /transaction/broadcast**

The service validates the transaction, and returns HTTP `400` if invalid.
Else, it signs the transaction, creates a new transaction record in the database with state `PENDING`,
publishes a `broadcast` message with the transaction id to the message queue, and returns HTTP `200`.

**POST /transaction/retry**

This endpoint takes a `transaction_id` body parameter.

The service queries the database by transaction id and returns HTTP `400` if the transaction does not exist or is not in a `FAILED` state.
Else, it updates the transaction record to state `PENDING` in the database and publishes a `broadcast` message with the transaction id to the message queue and returns HTTP `200`.

### Broadcasting Service

The broadcasting service consumes messages with transaction ids from the message queue and queries the database by transaction id.
If the message is a `check and broadcast` message, it queries the blockchain node to check if the transaction was successful.
If so, it updates the transaction record to `SUCCESS` in the database.
Else, or if the message was a `broadcast` message, it broadcasts the transaction to the blockchain node using a RPC request and proceeds based on the response.

**Success**

If the RPC request succeeds, the service updates the transaction record to `SUCCESS` in the database.

**Failure / Time Out**

If the RPC request fails or times out, the service uses an exponential backoff strategy with maximum retries.
If the maximum number of retries is exceeded, the service updates the transaction record to `FAILURE` in the database.
Else, the service waits before publishing a message with the transaction id and an incremented retry count back on the message queue.
If the request failed, it uses a `broadcast` message, and if the request timed out, it uses a `check and broadcast` message.

### Database

The database stores the single source of truth on the state of all transactions.

The database is a relational database, and for each transaction record it stores the transaction ID, data, and state.

The transaction state can take the values `SUCCESS`, `FAILURE` , and `PENDING`.
The database is updated immediately by the transaction manager upon any change to the transaction state.
Hence, in the event of a system fault, the state of each transaction can always be recovered from the database and pending transactions can be checked for success and retried if necessary.

The database can use sharding by transaction id for horizontal scalability.
The database can also use master-slave asynchronous replication for scalability and fault-tolerance.


