package services

import (
	"log"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/models/web"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServicesImpl struct {
	UserRepo repository.UserRepo
	db       *pgxpool.Pool
}

func NewUserServicesImpl(db *pgxpool.Pool, UserRepo repository.UserRepo) UserServices {
	return &UserServicesImpl{UserRepo: repository.NewUserRepoImpl(db), db: db}
}

func (u *UserServicesImpl) CreateUser(userRequest *web.UserRequest) (*web.UserResponse, *helpers.CustomError) {
	user := &models.User{
		GoogleId:  userRequest.GoogleId,
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		AvatarUrl: userRequest.AvatarUrl,
	}

	log.Println("[SERVICE] Checking if user exists by email:", user.Email)

	existingUser, err := u.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		log.Println("[SERVICE] Error getting user:", err)
		return nil, err
	}

	if existingUser != nil {
		log.Println("[SERVICE] User already exists, returning existing user:", existingUser.Email)
	} else {
		log.Println("[SERVICE] User not found, creating new user:", user.Email)
		existingUser, err = u.UserRepo.CreateUser(user)
		if err != nil {
			log.Println("[SERVICE] Error creating user:", err)
			return nil, err
		}
		log.Println("[SERVICE] Successfully created new user:", existingUser.Email)
	}

	return &web.UserResponse{
		ID:        existingUser.ID,
		GoogleId:  existingUser.GoogleId,
		Name:      existingUser.Name,
		Email:     existingUser.Email,
		AvatarUrl: existingUser.AvatarUrl,
	}, nil
}
