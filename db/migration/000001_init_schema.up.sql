CREATE TABLE "room" (
  "room_id" integer PRIMARY KEY,
  "max_guests" integer,
  "balcony" boolean,
  "fridge" boolean,
  "indoor_pool" boolean,
  "gaming_console" boolean
);

CREATE TABLE "room_availability" (
  "room_id" integer,
  "date" date,
  "is_available" boolean,
  "night_rate" decimal(10,2)
);

CREATE UNIQUE INDEX ON "room_availability" ("room_id", "date");

ALTER TABLE "room_availability" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");