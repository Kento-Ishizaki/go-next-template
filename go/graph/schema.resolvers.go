package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/Kento-Ishizaki/go-next-template/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "1",
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "user" + input.UserID,
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "1",
			Text: "todo1",
			User: &model.User{
				ID:   "1",
				Name: "user1",
			},
			Done: false,
		},
		{
			ID:   "2",
			Text: "todo2",
			User: &model.User{
				ID:   "2",
				Name: "user2",
			},
			Done: false,
		},
		{
			ID:   "3",
			Text: "todo3",
			User: &model.User{
				ID:   "3",
				Name: "user3",
			},
			Done: true,
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
