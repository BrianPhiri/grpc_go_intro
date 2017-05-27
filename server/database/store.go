package database

import (
	"encoding/json"
	"fmt"
	"log"

	"encoding/binary"

	pb "github.com/BrianPhiri/grpc_go_intro/api"
	"github.com/boltdb/bolt"
)

type Store struct {
	db *bolt.DB
}

func (s *Store) createTodo(todo *pd.Todo) (*pb.Todo, error) {
	log.Print("Create method : %v", todo)

	err := s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.createBucketIfNotExists([]byte("todo"))
		if err != nil {
			return err
		}
		id, _ := b.NextSequence()
		todo.id = id

		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}

		return b.Put(itob(todo.id), buf)
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Store) listTodo([]*pb.Todo, error) {
	fmt.Fprintln("List method")

	var todo []*pb.Todo

	err := s.db.view(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		fmt.Println("instatiated bucket")
		if b == nil {
			fmt.Println("buckt does not exisit")
			return fmt.Errorf("Bucket does not exist")
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			todo := pb.Todo{}

			err := json.Unmarshal(v, todo)
			if err != nil {
				return err
			}
			todo = append(todo, todo)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Store) deleteTodo(id uint32) ([]*pb.Todo, error) {
	fmt.Println("Delete method")

	var todo []*pb.Todo

	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		if b == nil {
			return fmt.Errorf("Nothing in bucket")
		}

		err := b.Delete(itob(id))
		if err != nil {
			return err
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			todo := &pb.Todo{}
			err := json.Unmarshal(v, todo)
			if err != nil {
				return nil
			}
		}
		todos = append(todos, todo)

		return nil

	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func itob(v uint32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, uint32(v))
	return b
}
