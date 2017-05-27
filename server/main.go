package main

const (
	port: ":1234"
)

type server struct{
	store.database.Database
}

fun (s *server) createTodo()