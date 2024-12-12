# Airbnb API Documentation

---
### Dec 10th 11:48 PM: Received Mail

Thank you for applying for the Backend Intern position at Saffronstays. I am Bidipto Bose, Senior Software Development Engineer at Saffronstays, and I'm excited to share the next step in our evaluation process.

To assess your backend development skills, we've prepared a technical assignment that involves creating a REST API using Go. Please find the details below:

---
#### *Assignment Details*

##### Objective

Develop a REST API to provide performance metrics for an Airbnb room.
Requirements

   1. *Framework*: Use the *Go-Gin* framework for building the API.
   2. *Database*: Use any SQL database of your choice (e.g., PostgreSQL,
   MySQL).

**Step 1: Add Room Data**

	Set up the database and create a table to store room data with fields like:
	
	   - Room ID (room_id)
	   - Rate per night (rate_per_night)
	   - Maximum number of guests (max_guests)
	   - Available dates (available_dates)
	   - Any other relevant data
	
	Populate the table with some sample data to work with.

**Step 2: API Specifications**

	1. Endpoint format:
	   GET /<room_id>
	2. Input:
	   - room_id: Path parameter representing the Airbnb room.
	3. Output:
	   - *Occupancy percentage* over the next 5 months (month-on-month) using
	      the availability data from the database.
	      - *Night rates for the next 30 days*, including:
	         - Average rate
	         - Highest rate
	         - Lowest rate
	4. The API must be testable using a browser or Postman.

------------------------------
#### *Submission Details*

1. *Code Repository*: Share the GitHub link to your code.
2. *Documentation*: Include a short note explaining the challenges you faced during the implementation and how you overcame them.

##### Deadline

Please submit your work within *3 days* of receiving this email.
##### Evaluation Criteria

	We will assess:
	
	   - Your understanding of the requirements.
	   - The quality of your code and use of Go-Gin.
	   - Database design and query optimization.
	   - Your problem-solving approach and documentation clarity.

------------------------------

If you have any questions or need further clarification, feel free to reach
out to me.

We are looking forward to reviewing your assignment. Best of luck!

Warm regards,
*Bidipto Bose*
Senior Software Development Engineer
Saffronstays

---
### Dec 11 9:45 AM: Designing Database
#### Fields of Room Data (JSON)

	- Room ID (room_id)
	- Rate per night (rate_per_night) --- each day over next 30 days
	- Maximum number of guests (max_guests)
	- Available dates (available_dates) --- calculative
	- Occupancy Percentage (occupancy_percentage) --- each month ---calculative
	- Average rate (average_rate) --- over next 30 days --- calculative
	- Highest rate (highest_rate) --- over next 30 days --- calculative
	- Lowest rate (lowest_rate) --- over next 30 days --- calculative
	- Balcony (balcony)
	- Mini Fridge (fridge)
	- Indoor Pool (indoor_pool)
	- Gaming Console (gaming_console)

Note: Database needs to keep data of the next 150 days

#### Database Tables
#ChatGPT_helps:P
##### Table `room`
Stores static attributes of each room.

| Column Name      | Data Type  | Description                                 |
| ---------------- | ---------- | ------------------------------------------- |
| `room_id`        | `INT` (PK) | Unique identifier for each room.            |
| `max_guests`     | `INT`      | Maximum occupancy for the room.             |
| `balcony`        | `BOOLEAN`  | Indicates if the room has a balcony.        |
| `fridge`         | `BOOLEAN`  | Indicates if the room has a mini-fridge.    |
| `indoor_pool`    | `BOOLEAN`  | Indicates if the room has access to a pool. |
| `gaming_console` | `BOOLEAN`  | Indicates if the room has a gaming console. |

##### Table `room_availability`
Tracks daily availability and rates for each room.

| Column Name    | Data Type       | Description                          |
| -------------- | --------------- | ------------------------------------ |
| `room_id`      | `INT` (FK)      | Foreign key linking to `rooms`.      |
| `date`         | `DATE`          | Specific date for room availability. |
| `is_available` | `BOOLEAN`       | Indicates if the room is available.  |
| `night_rate`   | `DECIMAL(10,2)` | Custom rate for this specific date.  |

**Indexes**:

- `(room_id, date)` for fast querying by room and date.
#### PostgreSQL Code
#dbdiagram_rocks

```sql
CREATE TABLE "room" (
	"room_id" integer PRIMARY KEY NOT NULL,
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
```

### Dec 11 10:35 AM: Setting up the database in the project
#### Initializing go.mod and git
```bash
go mod init github.com/vivek-344/airbnb-api
git init
git add .
git commit -m "initial commit"
git branch -M master
git remote add origin https://github.com/vivek-344/airbnb-api.git
git push -u origin master
```
#### Starting docker postgres and creating database
```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=vivek -d postgres
docker exec -it postgres createdb --username=root --owner=root room_db
```
#### Setting up golang-migrate
downloading migrate on the system could have been a hassle but I already had it in my system :P
```bash
mkdir -p ./db/migration
migrate create -ext sql -dir db/migration -seq init_schema

// Paste SQL code in the two generated files to create and drop tables

// Run migrations to PostgreSQL
migrate -path db/migration -database "postgresql://root:vivek@localhost:5432/room_db?sslmode=disable" -verbose up
```

### Dec 11 11:15 AM: Setting up SQLC
After a lil daydreaming about the internship, getting back to the work at 11:35 AM
```bash
// Initialize SQLC for the project
sqlc init

// Configured sqlc.yaml
// Wrote sql file for each table

// Generate go code
sqlc generate
```
Struggled a bit to write SQL queries, but google and ChatGPT helped
### Dec 11 1:00 PM: Lunch Time
### Dec 11 3:00 PM: Feeding random data to the database
that was a hassle
### Dec 11 5:00 PM: Snacks and Play Time
### Dec 11 6:30 PM: Fixing Minor Bugs
that was a cake walk with a fresh mind
### Dec 11 6:40 PM: Working on the API route
```bash
// setting up gin
go get -u github.com/gin-gonic/gin
```
### Dec 11 8:00 PM: Done with the basic requirements (Dinner Time)
### Dec 11 10:00 PM: Structuring the codebase
### Dec 11 11:00 PM: Adding comments through ChatGPT :P
### Dec 12 9:30 AM: Adding Tests
Covered more than 50% tests in db package 
### Dec 12 11:30 AM: College Time :')
### Dec 12 5:30 PM: Back at it!!!!
### Dec 12 6:40 PM: Completed tests for db package
coverage: 87.6% of statements
### Dec 12 7:00 PM: Added ci.yml and fixing some bugs