-- name: GetPraticien :one
SELECT * FROM Praticiens WHERE id = ?;

-- name: GetAllPraticiens :many
SELECT * FROM Praticiens ORDER BY name;

-- name: GetPraticiens :many
SELECT * FROM Praticiens ORDER BY name LIMIT ? OFFSET ?;

-- name: GetPraticiensByProfession :many
SELECT * FROM Praticiens WHERE profession = ? ORDER BY name LIMIT ? OFFSET ?;

-- name: GetAllProfessions :many
SELECT DISTINCT profession FROM Praticiens;
