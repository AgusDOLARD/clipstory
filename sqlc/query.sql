-- name: GetAll :many
SELECT * FROM clips
ORDER BY saved DESC;

-- name: Create :exec
INSERT INTO clips 
(value) VALUES (?);

-- name: Clear :exec
DELETE FROM clips;
