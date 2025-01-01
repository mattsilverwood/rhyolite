-- name: GetNoteById :one
SELECT * FROM note WHERE id = ? LIMIT 1;

-- name: ListAllNotes :many
SELECT * FROM note;

-- name: CreateNote :one
INSERT INTO note (
    title, content 
) VALUES (
    ?, ?
)
RETURNING *;

-- name: UpdateNoteById :one
UPDATE note set title = ?, content = ? WHERE id = ? RETURNING *;

-- name: DeleteNoteById :exec
DELETE FROM note WHERE id = ?;
