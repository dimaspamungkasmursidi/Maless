package repository

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	models "github.com/Bobby-P-dev/todo-listgo.git/internal/models/domains"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoImpl struct {
	db *pgxpool.Pool
}

func NewUserRepoImpl(db *pgxpool.Pool) UserRepo {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) GetUserByEmail(email string) (*models.User, *helpers.CustomError) {
	ctx := context.Background()
	var user models.User

	queryGet := `SELECT id, google_id, name, email, avatar_url FROM "user" WHERE email = $1`

	err := u.db.QueryRow(ctx, queryGet, email).Scan(&user.ID, &user.GoogleId, &user.Name, &user.Email, &user.AvatarUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("[REPOSITORY] User not found with email:", email)
			return nil, nil
		}
		log.Println("[REPOSITORY] Error scanning user:", err)
		return nil, helpers.NewCustomError("Failed to retrieve user", http.StatusInternalServerError)
	}

	log.Println("[REPOSITORY] Found existing user:", user.Email)
	return &user, nil
}

func (u *UserRepoImpl) CreateUser(user *models.User) (*models.User, *helpers.CustomError) {

	ctx := context.Background()
	query := `
        INSERT INTO "user"
            (google_id, name, email, avatar_url)
        VALUES
            ($1, $2, $3, $4)
        RETURNING
            id, google_id, name, email, avatar_url`

	var newUser models.User

	err := u.db.QueryRow(
		ctx,
		query,
		user.GoogleId,
		user.Name,
		user.Email,
		user.AvatarUrl,
	).Scan(
		&newUser.ID,
		&newUser.GoogleId,
		&newUser.Name,
		&newUser.Email,
		&newUser.AvatarUrl,
	)
	if err != nil {
		return nil, helpers.NewCustomError("Error creating user", http.StatusInternalServerError)
	}

	return &newUser, nil
}

func (u *UserRepoImpl) GetAllUsers() ([]*models.User, *helpers.CustomError) {

	ctx := context.Background()
	var users []*models.User

	query := "SELECT id, google_id, name, email, avatar_url FROM users"

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, helpers.NewCustomError("Error getting users", 500)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.GoogleId, &user.Name, &user.Email, &user.AvatarUrl)
		if err != nil {
			return nil, helpers.NewCustomError("Error getting users", 500)
		}
		users = append(users, &user)
	}

	return users, nil
}
