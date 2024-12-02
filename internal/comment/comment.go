package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplimented  = errors.New("not implemented")
)

// Comment - a representation of the comment
// struct for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all of
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new Service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	// Implement logic to fetch a
	fmt.Println("Retrieving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println((err))
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplimented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplimented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplimented
}
