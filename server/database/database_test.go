package database

import "testing"

func testCreatetodo(t *testing.T) {
	db, teardown, err := setupTestDatabase()
	if err != nil {
		t.Errorf("Error creating database : %v", err)
	}
	defer teardown()

	store := NewStore(db)
	mock := &pb.Todo{
		title:       "test 1",
		description: "Testing create",
		done:        false,
	}
}
