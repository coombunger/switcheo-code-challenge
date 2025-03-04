# crude

Use `ignite chain serve` to start the blockchain.

### CRUD Interface

Users can interact with the blockchain using the cli with any `account`, such as the generated `alice` and `bob` accounts.

Users can create, update, delete, show, and list resources.
Each resource has an id, name, value, and creator.

**Create Resource** 

```
cruded tx crude create <name> <value> --from <account> --chain-id crude
```



**Update Resource**

All parameters are mandatory, i.e. using `""` for `name` or `value` will update that property to become the empty string.

Resources can only be updated by their creator.

```
cruded tx crude update <name> <value> <id> --from <account> --chain-id crude
```



**Delete Resource**

Resources can only be deleted by their creator.

```
cruded tx crude delete <id> --from <account> --chain-id crude
```



**Show Resource**

```
cruded q crude show <id>
```



**List Resources**

Returns all resources with name equal to `name_filter` and value equal to `value_filter`.
Use `""` for either filter to not filter on that property.

Results are paginated by `limit`, and `offset` or `page_key`. 
The next page key is returned if there are more resources.

```
cruded q crude list <name_filter> <value_filter> --page-limit=<limit> [--page-offset=<offset>] [--page-key=<page_key>]
```