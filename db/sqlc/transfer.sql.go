// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO
    transfers (
        from_account_id,
        to_account_id,
        amount
    )
VALUES ($1, $2, $3)
RETURNING
    id, created_at, from_account_id, to_account_id, amount
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, created_at, from_account_id, to_account_id, amount FROM transfers WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
	)
	return i, err
}

const listTransfersOfDeposit = `-- name: ListTransfersOfDeposit :many
SELECT id, created_at, from_account_id, to_account_id, amount
FROM transfers
WHERE
    to_account_id = $1
ORDER BY id
LIMIT $2
OFFSET
    $3
`

type ListTransfersOfDepositParams struct {
	ToAccountID int64 `json:"to_account_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListTransfersOfDeposit(ctx context.Context, arg ListTransfersOfDepositParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersOfDeposit, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersOfWithdraw = `-- name: ListTransfersOfWithdraw :many
SELECT id, created_at, from_account_id, to_account_id, amount
FROM transfers
WHERE
    from_account_id = $1
ORDER BY id
LIMIT $2
OFFSET
    $3
`

type ListTransfersOfWithdrawParams struct {
	FromAccountID int64 `json:"from_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfersOfWithdraw(ctx context.Context, arg ListTransfersOfWithdrawParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersOfWithdraw, arg.FromAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}