CREATE TABLE "room" (
  "room_id" integer PRIMARY KEY,
  "max_guests" integer NOT NULL,
  "balcony" boolean NOT NULL,
  "fridge" boolean NOT NULL,
  "indoor_pool" boolean NOT NULL,
  "gaming_console" boolean NOT NULL
);

CREATE TABLE "room_availability" (
  "room_id" integer NOT NULL,
  "date" date NOT NULL,
  "is_available" boolean NOT NULL,
  "night_rate" integer NOT NULL
);

CREATE UNIQUE INDEX ON "room_availability" ("room_id", "date");

ALTER TABLE "room_availability" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("room_id");