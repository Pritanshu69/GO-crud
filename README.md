#  Created a CRUD Operation code using go lang!

1.Data Model: Defined within the same file using the Item struct.
2.In-Memory Store: Uses a global map and a mutex for concurrent access.
3.Handlers:
- createItem: Adds a new item.
- readItem: Retrieves an item by ID.
- updateItem: Updates an existing item.
- deleteItem: Deletes an item.
4.Routing: Uses gorilla/mux to handle HTTP routes and methods.
5.Server Setup: Starts an HTTP server on port 8080.
 
### Structure of the directory: 
```bash
  /your-project-directory
|-- main.go
|-- go.mod
|-- go.sum

```
### You can then use tools like curl to check the CRUD operation
- Create
```bash
  curl -X POST http://localhost:8080/items -d '{"id":"1","name":"Item 1"}' -H "Content-Type: application/json"
```
- Read
```bash
  curl http://localhost:8080/items/1
```
- Update
```bash
curl -X PUT http://localhost:8080/items/1 -d '{"id":"1","name":"Updated Item"}' -H "Content-Type: application/json"
```
- Delete
```bash
  curl -X DELETE http://localhost:8080/items/1
```
