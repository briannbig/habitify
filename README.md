This is a golang backend for a project to help users track their habits

### Getting started

```shell 
go mod install
go run .
```
Visit `:5050` to access the REST API
## Rest API Endpoints

| Method | Endpoint    | Description                  |
| ------ | ----------- | ---------------------------- |
| POST   | /register   | Register(sign up) a new user |
| POST   | /users      | Create a new user            |
| GET    | /users      | Get all users                |
| PUT    | /users/:id  | Update a user by ID          |
| DELETE | /users/:id  | Delete a user by ID          |
| POST   | /habits     | Create a new habit           |
| GET    | /habits     | Get all habits               |
| PUT    | /habits/:id | Update a habit by ID         |
| DELETE | /habits/:id | Delete a habit by ID         |
