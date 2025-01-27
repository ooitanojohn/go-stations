package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject string, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	fmt.Println("Prepare:", subject)
	// Prepare the insert statement
	stmt, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	fmt.Println("Execute:", subject)
	// Execute the insert statement
	result, err := stmt.ExecContext(ctx, subject, description)
	if err != nil {
		return nil, err
	}

	fmt.Println("LastGet:", subject)
	// Get the ID of the inserted TODO
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	fmt.Println("lastInsertId:", id)

	fmt.Println("Prepare:", subject)
	// Prepare the confirm statement
	confirmStmt, err := s.db.PrepareContext(ctx, confirm)
	if err != nil {
		return nil, err
	}
	defer confirmStmt.Close()

	fmt.Println("Query:", subject)
	// Query the inserted TODO by ID
	row := confirmStmt.QueryRowContext(ctx, id)

	// Create a TODO object to return
	todo := &model.TODO{}
	err = row.Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
