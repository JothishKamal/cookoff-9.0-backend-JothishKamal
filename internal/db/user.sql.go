// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const banUser = `-- name: BanUser :exec
UPDATE users
SET is_banned = TRUE
WHERE id = $1
`

func (q *Queries) BanUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, banUser, id)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, email, reg_no, password, role, round_qualified, score, name)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, email, reg_no, password, role, round_qualified, score, name, is_banned
`

type CreateUserParams struct {
	ID             uuid.UUID
	Email          string
	RegNo          string
	Password       string
	Role           string
	RoundQualified int32
	Score          pgtype.Int4
	Name           string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.RegNo,
		arg.Password,
		arg.Role,
		arg.RoundQualified,
		arg.Score,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.RegNo,
		&i.Password,
		&i.Role,
		&i.RoundQualified,
		&i.Score,
		&i.Name,
		&i.IsBanned,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, reg_no, password, role, round_qualified, score, name, is_banned
FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.RegNo,
			&i.Password,
			&i.Role,
			&i.RoundQualified,
			&i.Score,
			&i.Name,
			&i.IsBanned,
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, reg_no, password, role, round_qualified, score, name, is_banned
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.RegNo,
		&i.Password,
		&i.Role,
		&i.RoundQualified,
		&i.Score,
		&i.Name,
		&i.IsBanned,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, reg_no, password, role, round_qualified, score, name, is_banned
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.RegNo,
		&i.Password,
		&i.Role,
		&i.RoundQualified,
		&i.Score,
		&i.Name,
		&i.IsBanned,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, email, reg_no, password, role, round_qualified, score, name, is_banned
FROM users
WHERE name = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.RegNo,
		&i.Password,
		&i.Role,
		&i.RoundQualified,
		&i.Score,
		&i.Name,
		&i.IsBanned,
	)
	return i, err
}

const unbanUser = `-- name: UnbanUser :exec
UPDATE users
SET is_banned = FALSE
WHERE id = $1
`

func (q *Queries) UnbanUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, unbanUser, id)
	return err
}

const upgradeUsersToRound = `-- name: UpgradeUsersToRound :exec
UPDATE users
SET round_qualified = GREATEST(round_qualified, $2)
WHERE id::TEXT = ANY($1::TEXT[])
`

type UpgradeUsersToRoundParams struct {
	Column1        []string
	RoundQualified int32
}

func (q *Queries) UpgradeUsersToRound(ctx context.Context, arg UpgradeUsersToRoundParams) error {
	_, err := q.db.Exec(ctx, upgradeUsersToRound, arg.Column1, arg.RoundQualified)
	return err
}
