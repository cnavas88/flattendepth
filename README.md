# flattendepth
Repository to flatten array and return its depth

## Description
This project is an API Rest, you can call to API with an array as parameter and return this nested array convert to flatten array. get the depth of the nested array too.

# How to use

1- clone the repository
2- run the project with go run main.go (listen to :8080)
3- open postman and generate the next request:
* method: POST
* url: localhost:8080
* body:
```json
{
  "array": [[10, 20, 30], 40]
}
```



