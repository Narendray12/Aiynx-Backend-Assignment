package repository

import (
	"context"
	"database/sql"
	"time"

	db "aiynx/db/migrations"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		q: db.New(dbConn),
	}
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (db.User, error) {
	return r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *UserRepository) List(ctx context.Context, limit, offset int32) ([]db.User, error) {
	return r.q.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (r *UserRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	return r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}
