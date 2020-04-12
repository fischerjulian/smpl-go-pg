# Smpl-Go-Pg

A minimalistic example of a web app accessing a PostgreSQL [1]database written in Go [2].

The app is intended for demonstration purposes as part of Containerization tutorials.

## Run

Assuming there is a PostgreSQL server running with a user `postgres` and an empty database `postgresl`:

    POSTGRES_HOST=<postgresql-host> POSTGRES_PASSWORD=<postgresl-password> go run main.go

## Links

1. PostgreSQL, https://www.postgresql.org/
2. Go, https://golang.org/
3. Go PQ, https://github.com/lib/pq