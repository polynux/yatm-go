// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const getPraticien = `-- name: GetPraticien :one
SELECT id, name, firstname, address, zip, city, description, profession FROM Praticiens WHERE id = ?
`

func (q *Queries) GetPraticien(ctx context.Context, id interface{}) (Praticien, error) {
	row := q.db.QueryRowContext(ctx, getPraticien, id)
	var i Praticien
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Firstname,
		&i.Address,
		&i.Zip,
		&i.City,
		&i.Description,
		&i.Profession,
	)
	return i, err
}

const getPraticiens = `-- name: GetPraticiens :many
SELECT id, name, firstname, address, zip, city, description, profession FROM Praticiens
`

func (q *Queries) GetPraticiens(ctx context.Context) ([]Praticien, error) {
	rows, err := q.db.QueryContext(ctx, getPraticiens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Praticien
	for rows.Next() {
		var i Praticien
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Firstname,
			&i.Address,
			&i.Zip,
			&i.City,
			&i.Description,
			&i.Profession,
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
