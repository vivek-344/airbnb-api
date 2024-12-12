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

-- name: UpdateRoomAvailability :one
UPDATE room_availability
SET is_available = $3,
    night_rate = $4
WHERE room_id = $1 AND date = $2
RETURNING *;

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
WHERE room_id = $1 AND is_available = TRUE
LIMIT 30;

-- name: GetAverageRate :one
SELECT AVG(night_rate) 
FROM room_availability
WHERE room_id = $1
  AND date >= CURRENT_DATE 
  AND date < CURRENT_DATE + INTERVAL '30 days';

-- name: GetMaximumRate :one
SELECT night_rate
FROM room_availability
WHERE room_id = $1
  AND date >= CURRENT_DATE 
  AND date < CURRENT_DATE + INTERVAL '30 days'
ORDER BY night_rate DESC
LIMIT 1;

-- name: GetMaxDate :one
SELECT date
FROM room_availability
ORDER BY night_rate DESC
LIMIT 1;

-- name: GetMinimumRate :one
SELECT night_rate
FROM room_availability
WHERE room_id = $1
  AND date >= CURRENT_DATE 
  AND date < CURRENT_DATE + INTERVAL '30 days'
ORDER BY night_rate ASC
LIMIT 1;

-- name: GetAvailabilityPercentage :many
SELECT
 EXTRACT(YEAR FROM date) AS year,
 EXTRACT(MONTH FROM date) AS month,
 CAST(COUNT(CASE WHEN is_available THEN 1 END) * 100.0 / COUNT(*) AS DECIMAL(10,2)) AS availability_percentage
FROM room_availability
WHERE room_id = $1
GROUP BY year, month
ORDER BY year, month;

-- name: GetDateCount :one
SELECT COUNT(DISTINCT date) FROM room_availability;

-- name: DeleteOldRoomAvailabilityData :exec
DELETE FROM room_availability WHERE date < CURRENT_DATE;

-- name: DeleteAllAvailabilityForRoom :exec
DELETE FROM room_availability WHERE room_id = $1;