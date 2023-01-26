package repository

import (
	"context"

	"github.com/Egor-Tihonov/SandMailLogic/internal/models"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

func (r *PostgresDB) GetEmailFromDB(ctx context.Context, email string) (string, error) {
	name := ""
	err := r.Pool.QueryRow(ctx, "select name from users where email=$1", email).Scan(&name)
	if err != nil {
		if name == "" {
			return "", models.ErrorUserDoesntExist
		}
		logrus.Error("database error: %e", err.Error())
		return "", err
	}

	return name, nil
}

func (r *PostgresDB) UpdatePassword(ctx context.Context, email string, password string) error {
	_, err := r.Pool.Exec(ctx, "update users set password = $1 where email = $2", &password, &email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.ErrorUserDoesntExist
		}
		return err
	}
	return nil
}
