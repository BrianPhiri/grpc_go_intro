package main

import(
	pb "github.com/BrianPhiri/grpc_go_intro/api"
)

const (
	port = ":1234"
)

type server struct{
	store database.Database
}

func (s *server) createTodo(ctx context.Context, in *pb.createTodoRequest) (*pb.createTodoResponse, error){
	todo, err := s.store.createTodo(in.Todo)
	if err != nil{
		return nil, err
	}
	return &pb.createTodoResponse{
		Todo: todo
	}, nil
}

func (s *server) listTodo(ctx contex.Context, in *pb.listTodoRequest) (*pb.listTodoResponse, error){
	todo, err := s.store.listTodo()
	if(err !=nil){
		return nil, err
	}
	return &pb.listTodoResponse{
		Todo: todo
	}, nil
}

func (s *server) deleteTodo(ctx context.Context, in *pb.deleteTodoRequest) (*pb.deleteTodoResponse, error){
	todo, err := s.store.deleteTodo(in.id)
	if(err != nil){
		return nil, err
	}
	return &pb.deleteTodoResponse{
		Todo: todo
	}, nil
}