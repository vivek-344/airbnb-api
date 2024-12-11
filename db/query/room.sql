-- name: CreateRoom :one
INSERT INTO room (
  room_id,
  max_guests,
  balcony,
  fridge,
  indoor_pool,
  gaming_console
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListAllRoomIDs :many
SELECT room_id 
FROM room;

-- name: GetRoom :one
SELECT * FROM room
WHERE room_id = $1 LIMIT 1;

-- name: ListRooms :many
SELECT * FROM room
ORDER BY room_id
LIMIT $1
OFFSET $2;

-- name: UpdateMaxGuests :exec
UPDATE room
SET max_guests = $2
WHERE room_id = $1
RETURNING *;

-- name: UpdateRoomFridge :exec
UPDATE room
SET fridge = $2
WHERE room_id = $1
RETURNING *;

-- name: UpdateRoomConsole :exec
UPDATE room
SET gaming_console = $2
WHERE room_id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM room WHERE room_id = $1;