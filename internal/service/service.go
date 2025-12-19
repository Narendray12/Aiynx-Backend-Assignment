package service

import (
	"context"
	"time"

	"aiynx/internal/models"
	"aiynx/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

/*
========================
Business helpers
========================
*/
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

/*
========================
GET /users/:id
========================
*/
func (s *UserService) GetUser(ctx context.Context, id int32) (models.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob,
		Age:  CalculateAge(u.Dob),
	}, nil
}

/*
========================
POST /users
========================
*/
func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (models.User, error) {

	u, err := s.repo.Create(ctx, name, dob)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob,
		Age:  CalculateAge(u.Dob),
	}, nil
}

/*
========================
GET /users
========================
*/
func (s *UserService) ListUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]models.User, error) {

	users, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	result := make([]models.User, 0, len(users))
	for _, u := range users {
		result = append(result, models.User{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob,
			Age:  CalculateAge(u.Dob),
		})
	}

	return result, nil
}

/*
========================
PUT /users/:id
========================
*/
func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob time.Time,
) (models.User, error) {

	u, err := s.repo.Update(ctx, id, name, dob)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob,
		Age:  CalculateAge(u.Dob),
	}, nil
}

/*
========================
DELETE /users/:id
========================
*/
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
