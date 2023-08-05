package repository

import (
	"atom-project/internal/models"
	"context"
)

func (r repository) GetAllLessons(ctx context.Context) ([]models.Lesson, error) {
	const query = `select ID, Title, BlockingLesson, BlockingText from lessons`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lessons := make([]models.Lesson, 0)
	for rows.Next() {
		l := models.Lesson{}
		err = rows.Scan(&l.ID, &l.Title, &l.BlockingLesson, &l.BlockingText)
		if err != nil {
			return nil, err
		}

		tags, err := r.GetLessonTags(ctx, l.ID)
		if err != nil {
			return nil, err
		}
		l.Tags = tags

		lessons = append(lessons, l)
	}

	return lessons, nil
}

func (r repository) GetLessonTags(ctx context.Context, lessonID int) ([]string, error) {
	const query = `select Tag from lessons_tags where LessonID=$1`

	rows, err := r.pool.Query(ctx, query, lessonID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r repository) GetUserLessons(ctx context.Context, token string) ([]models.ListLesson, error) {
	userID, err := r.GetUserID(ctx, token)
	if err != nil {
		return nil, err
	}

	lessons, err := r.GetAllLessons(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]models.ListLesson, 0)
	const query = `select exists(select 1 from completed_lessons where UserID=$1 and LessonID=$2)`

	for _, l := range lessons {
		var IsCompleted, IsOpen bool
		row := r.pool.QueryRow(ctx, query, userID, l.ID)
		err := row.Scan(&IsCompleted)
		if err != nil {
			return nil, err
		}

		if l.BlockingLesson == 0 {
			IsOpen = true
		} else {
			row := r.pool.QueryRow(ctx, query, userID, l.BlockingLesson)
			err := row.Scan(&IsOpen)
			if err != nil {
				return nil, err
			}
		}

		result = append(result, models.ListLesson{
			Lesson:      l,
			IsCompleted: IsCompleted,
			IsOpen:      IsOpen,
		})
	}

	return result, err
}
