// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: submission.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createSubmission = `-- name: CreateSubmission :exec
INSERT INTO submissions (id, user_id, question_id, language_id)
VALUES ($1, $2, $3, $4)
`

type CreateSubmissionParams struct {
	ID         uuid.UUID
	UserID     uuid.NullUUID
	QuestionID uuid.UUID
	LanguageID int32
}

func (q *Queries) CreateSubmission(ctx context.Context, arg CreateSubmissionParams) error {
	_, err := q.db.Exec(ctx, createSubmission,
		arg.ID,
		arg.UserID,
		arg.QuestionID,
		arg.LanguageID,
	)
	return err
}

const getSubmission = `-- name: GetSubmission :one
SELECT 
    testcases_passed, 
    testcases_failed 
FROM 
    submissions 
WHERE 
    id = $1
`

type GetSubmissionRow struct {
	TestcasesPassed pgtype.Int4
	TestcasesFailed pgtype.Int4
}

func (q *Queries) GetSubmission(ctx context.Context, id uuid.UUID) (GetSubmissionRow, error) {
	row := q.db.QueryRow(ctx, getSubmission, id)
	var i GetSubmissionRow
	err := row.Scan(&i.TestcasesPassed, &i.TestcasesFailed)
	return i, err
}

const getSubmissionsByUserId = `-- name: GetSubmissionsByUserId :many
SELECT id, question_id, testcases_passed, testcases_failed, runtime, submission_time, testcase_id, language_id, description, memory, user_id, status
FROM submissions
WHERE user_id = $1
`

func (q *Queries) GetSubmissionsByUserId(ctx context.Context, userID uuid.NullUUID) ([]Submission, error) {
	rows, err := q.db.Query(ctx, getSubmissionsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Submission
	for rows.Next() {
		var i Submission
		if err := rows.Scan(
			&i.ID,
			&i.QuestionID,
			&i.TestcasesPassed,
			&i.TestcasesFailed,
			&i.Runtime,
			&i.SubmissionTime,
			&i.TestcaseID,
			&i.LanguageID,
			&i.Description,
			&i.Memory,
			&i.UserID,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTestCases = `-- name: GetTestCases :many
SELECT id, expected_output, memory, input, hidden, runtime, question_id 
FROM testcases
WHERE question_id = $1
  AND (CASE WHEN $2 = TRUE THEN hidden = FALSE ELSE TRUE END)
`

type GetTestCasesParams struct {
	QuestionID uuid.UUID
	Column2    interface{}
}

func (q *Queries) GetTestCases(ctx context.Context, arg GetTestCasesParams) ([]Testcase, error) {
	rows, err := q.db.Query(ctx, getTestCases, arg.QuestionID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Testcase
	for rows.Next() {
		var i Testcase
		if err := rows.Scan(
			&i.ID,
			&i.ExpectedOutput,
			&i.Memory,
			&i.Input,
			&i.Hidden,
			&i.Runtime,
			&i.QuestionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSubmission = `-- name: UpdateSubmission :exec
UPDATE submissions
SET testcases_passed = $1, testcases_failed = $2, runtime = $3, memory = $4
WHERE id = $5
`

type UpdateSubmissionParams struct {
	TestcasesPassed pgtype.Int4
	TestcasesFailed pgtype.Int4
	Runtime         pgtype.Numeric
	Memory          pgtype.Int4
	ID              uuid.UUID
}

func (q *Queries) UpdateSubmission(ctx context.Context, arg UpdateSubmissionParams) error {
	_, err := q.db.Exec(ctx, updateSubmission,
		arg.TestcasesPassed,
		arg.TestcasesFailed,
		arg.Runtime,
		arg.Memory,
		arg.ID,
	)
	return err
}
