-- name: GetPraticien :one
SELECT * FROM Praticiens WHERE id = ?;

-- name: GetPraticiens :many
SELECT * FROM Praticiens;
