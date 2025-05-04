package postgre

import (
	"context"
	"em/internal/model"
	"em/internal/storage"
	"em/pkg/utils"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (this *Storage) SaveUser(ctx context.Context, user *model.User) error {
	const op string = "storage.postgre.user.SaveUser"

	return utils.WrapError(this.db.QueryRow(ctx, `
		INSERT INTO users(name, surname, patronymic, age, nationality, gender) 
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id`,
		user.Name, user.Surname, user.Patronymic, user.Age, user.Nationality, user.Gender,
	).Scan(&user.Id), op, "db.QueryRow", "Scan")
}

func (this *Storage) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	const op string = "storage.postgre.user.GetUserById"

	var user model.User
	if err := this.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id).Scan(
		&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Nationality, &user.Gender,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: db.QueryRow: Scan: %w", op, err)
	}

	return &user, nil
}

func (this *Storage) UpdateUser(ctx context.Context, user *model.User) error {
	const op string = "storage.postgre.user.UpdateUser"

	tag, err := this.db.Exec(ctx, `
		UPDATE users SET 
			name = $1, surname = $2, patronymic = $3, nationality = $4, gender = $5, age = $6
		WHERE id = $7`,
		user.Name, user.Surname, user.Patronymic, user.Nationality, user.Gender, user.Age, user.Id,
	)
	if err != nil {
		return fmt.Errorf("%s: db.Exec: %w", op, err)
	}

	if tag.RowsAffected() == 0 {
		return storage.ErrUserNotFound
	}

	return nil
}

func (this *Storage) DeleteUserById(ctx context.Context, id int64) error {
	const op string = "storage.postgres.DeleteUserById"

	tag, err := this.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("%s: db.Exec: %w", op, err)
	}

	if tag.RowsAffected() == 0 {
		return storage.ErrUserNotFound
	}

	return nil
}
