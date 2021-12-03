-- name: CreateTodo :one
INSERT INTO todos (
    task,
    is_done
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: FindTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: UpdateTodo :one
UPDATE todos SET 
    task = $2,
    is_done = $3
WHERE id = $1
RETURNING id AS updated_id;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING id AS deleted_id;