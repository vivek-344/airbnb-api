-- name: CreateRoomAvailability :one
INSERT INTO room_availability (
  room_id,
  date,
  is_available,
  night_rate
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateRoomAvailability :exec
UPDATE room_availability
SET is_available = $3,
    night_rate = $4
WHERE room_id = $1 AND date = $2;

-- name: GetRoomAvailabilityByDate :one
SELECT * FROM room_availability
WHERE room_id = $1 AND date = $2 
LIMIT 1;

-- name: ListRoomAvailability :many
SELECT date, is_available, night_rate FROM room_availability
WHERE room_id = $1
LIMIT 30;

-- name: ListAvailableDates :many
SELECT date FROM room_availability
WHERE room_id = $1 AND is_available = TRUE;

-- name: GetAverageRate :one
SELECT AVG(night_rate) 
FROM room_availability
WHERE room_id = $1
  AND date >= CURRENT_DATE 
  AND date < CURRENT_DATE + INTERVAL '30 days';

-- name: GetMaximumRate :one
SELECT MAX(night_rate) 
FROM room_availability
WHERE room_id = $1
  AND date >= CURRENT_DATE 
  AND date < CURRENT_DATE + INTERVAL '30 days';

-- name: GetMinimumRate :one
SELECT MIN(night_rate) 
FROM room_availability
WHERE room_id = $1
  AND date >= CURRENT_DATE 
  AND date < CURRENT_DATE + INTERVAL '30 days';

-- name: GetAvailabilityPercentage :many
SELECT
  EXTRACT(YEAR FROM date) AS year,
  EXTRACT(MONTH FROM date) AS month,
  COUNT(CASE WHEN is_available THEN 1 END) * 100.0 / COUNT(*) AS availability_percentage
FROM room_availability
WHERE room_id = $1
GROUP BY year, month
ORDER BY year, month;

-- name: GetDateCount :one
SELECT COUNT(DISTINCT date) FROM room_availability;

-- name: DeleteRoomAvailabilityData :exec
DELETE FROM room_availability WHERE date < CURRENT_DATE;