# Smpl-Go-Pg

A minimalistic example of a web app accessing a PostgreSQL [1]database written in Go [2].

The app is intended for demonstration purposes as part of Containerization tutorials.

# Go

## Run

Assuming there is a PostgreSQL server running with a user `postgres` and an empty database `postgresl`:

    POSTGRES_HOST=<postgresql-host> POSTGRES_USERNAME=<postgresql-username> POSTGRES_PASSWORD=<postgresl-password> go run main.go

# Docker

## Build

    docker image build -t smpl-go-pg:0.1.0 .

## Tag

    docker image tag smpl-go-pg:0.1.0 fischerjulian/smpl-go-pg:0.1.0

## Publish to Registry

    docker image push fischerjulian/smpl-go-pg:1.0.0                                   

## Pull Image from Registry

    docker image pull fischerjulian/smpl-go-pg:0.1.0

## Run
In order to run the image you will also have to set the env vars `POSTGRES_HOST` and `POSTGRES_PASSWORD` which is not contained in the examples below as the images will be used in the context of Kubernetes, only.

Run local image with version tag `0.1.0`:

    docker container run -p 8080:8080 smpl-go-pg:0.1.0

Run remote image with version tag `0.1.0`:

    docker container run -p  8080:8080 fischerjulian/smpl-go-pg:0.1.0


## Links

1. PostgreSQL, https://www.postgresql.org/
2. Go, https://golang.org/
3. Go PQ, https://github.com/lib/pq