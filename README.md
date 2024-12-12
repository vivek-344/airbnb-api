# Airbnb Performance Metrics API  

This project implements a REST API for retrieving performance metrics of Airbnb rooms using **GoLang** and the **Gin framework**. It provides insights such as room occupancy percentages and nightly rates based on availability data stored in a PostgreSQL database.

---

## Features  

- **Room Occupancy**: Calculates the occupancy percentage for the next 5 months (month-on-month).  
- **Nightly Rates**: Provides the average, highest, and lowest rates for the next 30 days.  
- **Scalable Design**: Built with efficient SQL queries and a modular code structure.  
- **Test Coverage**: Achieved over 85% test coverage for the `db` package.  
- **Continuous Integration**: Configured GitHub Actions for automated testing.  

---

## Tech Stack  

- **Backend**: GoLang with the Gin framework  
- **Database**: PostgreSQL  
- **Migrations**: golang-migrate  
- **SQL Queries**: sqlc  
- **Testing**: Go's built-in testing framework  

---

## Prerequisites  

Ensure the following are installed on your system:  
- [Go](https://go.dev/doc/install) (v1.20 or later)  
- [Docker](https://docs.docker.com/get-docker/)  
- [golang-migrate](https://github.com/golang-migrate/migrate)  
- [PostgreSQL](https://www.postgresql.org/download/)  

---

## Setup  

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/vivek-344/airbnb-api.git
   cd airbnb-api
   ```  

2. **Run PostgreSQL and create Database**  
   Use Docker to start a PostgreSQL instance and create database:  
   ```bash
   make postgres
   make createdb
   ```  

3. **Run Database Migrations**  
   ```bash
   make migrateup
   ```  

4. **Run Tests**  
   ```bash
   make test
   ```  

5. **Start the Server**  
   ```bash
   make server
   ```  

---

## API Endpoints  

### 1. Get Room Metrics  
**Endpoint**: `GET /<room_id>`  

**Input**:  
- `room_id`: Path parameter representing the Airbnb room  

**Example Output**:  
The API returns a detailed JSON response containing the following metrics:  
- **Room ID**: Identifier for the Airbnb room.  
- **Occupancy Percentage**: Month-by-month occupancy for the next 5 months, showing the availability percentage.  
- **Nightly Rates**: A breakdown of the nightly rates for the next 30 days, including the highest, lowest, and average rates.  
- **Availability**: Dates when the room is available.  
- **Amenities**: Information about specific room amenities like a balcony, fridge, indoor pool, and gaming console.  

For a full example, see the [example.json](https://github.com/vivek-344/airbnb-api/example.json) file included in the repository.

---

## Testing  

1. **Run Tests**  
   ```bash
   go test ./... -cover
   ```  
2. **Test Coverage**  
   Current coverage: **87.6%** for the `db` package.

---

## Project Structure  

```
airbnb-api/
├── .github/workflows/ci.yml      # GitHub Actions CI workflow
├── api/                          # API routes and server configuration
├── db/                           # Database migrations and SQLC code
│   ├── migration/                # SQL migration files
│   ├── query/                    # Raw SQL queries
│   ├── sqlc/                     # SQLC generated code
├── util/                         # Utility files (config, random generators, etc.)
├── app.env                       # Environment configuration
├── main.go                       # Entry point of the application
├── Makefile                      # Automation commands
├── README.md                     # Documentation
├── journey.md                    # Development journey documentation
└── sqlc.yaml                     # SQLC configuration
```  

---

## Challenges and Learnings  

1. **Dynamic SQL Queries**: Designing efficient queries for availability and rate calculations.  
2. **Debugging**: Resolved minor bugs in the availability logic by leveraging Go's debugging tools.  
3. **Testing**: Writing comprehensive tests to ensure the accuracy of database operations.  

---

## Future Improvements  

- Add authentication and role-based access.  
- Enhance database schema for scalability.  
- Integrate Swagger for API documentation.  

---

## Author  

**Vivekraj Singh Sisodiya**  
📧 [vivekuit344@gmail.com](mailto:vivekuit344@gmail.com)
[GitHub](https://github.com/vivek-344) | [LinkedIn](https://www.linkedin.com/in/vivek344/)  

---

This README provides all necessary details for anyone setting up, testing, or evaluating your project. Let me know if you'd like to tweak anything!