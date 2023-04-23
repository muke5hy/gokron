# Cron Job Management

A simple web-based cron job management system built using Go, PostgreSQL, and the html/template package. This project allows you to schedule, manage, and view logs of cron jobs through a web UI.

## Features

- Schedule cron jobs with a web UI
- View and manage existing cron jobs
- View logs of executed cron jobs
- API for managing cron jobs and logs
- PostgreSQL for storing cron jobs and logs

## Prerequisites

- [Go](https://golang.org/doc/install) (1.16 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

## Getting Started

1. Clone the repository:

```sh
git clone https://github.com/yourusername/cron-job-product.git
cd cron-job-product
```

2. Update the PostgreSQL configuration in the db/db.go file with your database credentials.

3. Install the Go dependencies:

```sh
    go mod download
```

4. Build the project:

```sh
go build -o cron-job-manager
```

5. Run the project:
    
```sh   
./cron-job-manager
```

6. Visit http://localhost:8080 in your browser to access the web UI.

## API
This project also includes a RESTful API for managing cron jobs and logs:

- GET /api/cron-job: List all cron jobs
- POST /api/cron-job: Create a new cron job
- GET /api/cron-job/:id: Get details of a specific cron job
- PUT /api/cron-job/:id: Update a specific cron job
- DELETE /api/cron-job/:id: Delete a specific cron job
- GET /api/log/:cronJobID: Get logs for a specific cron job

## License
This project is released under the MIT License.
