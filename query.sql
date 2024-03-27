-- name: GetPraticien :one
SELECT * FROM Praticiens WHERE id = ?;

-- name: GetAllPraticiens :many
SELECT * FROM Praticiens;

-- name: GetPraticiens :many
SELECT * FROM Praticiens LIMIT ? OFFSET ?;

-- name: GetPraticiensByProfession :many
SELECT * FROM Praticiens WHERE profession = ?;

-- name: GetAllProfessions :many
SELECT DISTINCT profession FROM Praticiens;
