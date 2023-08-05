package repository

import (
	"context"
)

func (r repository) GetUserID(ctx context.Context, token string) (int, error) {
	const query = "select ID from users where Token=$1"

	var id int
	row := r.pool.QueryRow(ctx, query, token)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}
