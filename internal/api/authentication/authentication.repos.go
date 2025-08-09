package authentication

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/longnh462/go-gin-boilerplate/infra/database/postgres/entitys"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) GetUserByEmail(email string) (*entitys.UserEntity, error) {
	user := &entitys.UserEntity{}
	query := `SELECT user_id, username, email, password FROM users WHERE email = $1`

	err := ar.db.QueryRow(query, email).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (ar *AuthRepository) CreateUser(user *entitys.UserEntity) error {
	query := `INSERT INTO users (user_id, username, email, password) VALUES ($1, $2, $3, $4)`

	_, err := ar.db.Exec(query, user.UserId, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (ar *AuthRepository) GetUserRoles(userId uuid.UUID) ([]string, error) {
	query := `
		SELECT r.role_name 
		FROM roles r 
		JOIN user_roles ur ON r.role_id = ur.role_id 
		WHERE ur.user_id = $1
	`

	rows, err := ar.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err != nil {
			return nil, fmt.Errorf("failed to scan role: %w", err)
		}
		roles = append(roles, roleName)
	}

	return roles, nil
}
