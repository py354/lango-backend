package repository

import (
	"context"
)

func (r repository) SetUserSettings(ctx context.Context, token string, mode int, complexity int) (err error) {
	const query = `update users set EducationMode=$1, EducationComplexity=$2 where Token=$3`
	_, err = r.pool.Exec(ctx, query, mode, complexity, token)
	return
}
