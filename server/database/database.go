package database
// package store

import (
	pb "github.com/BrianPhiri/grpc_go_intro/api"
)

type Database interface{
	createTodo(todo *pb.Todo, error)
	listTodo() ([]*pb.Todo, error)
	deleteTodo(id uint32)([]*pb.Todo, error)
}