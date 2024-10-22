# RSS Aggregator

#### This is a rss scrapper service built with Golang

## Getting Started

To get started with this project, you can clone the repository and build the
project:

```bash
git clone https://github.com/haroonalbar/rss-aggregater
cd rss-aggregater
go build
```

Set up .env from .env.example
Set your db url in .env

```sh
mv .env.example .env
```

Migrate DB using goose

```sh
cd sql/schema
goose postgres <connection-url> up
```

Once the project is built, you can run the service:

```bash
./rss-aggregater
```

The service will start and begin scraping RSS feeds as configured.

## Technologies

- [Golang](https://go.dev/): The core programming language used for building the
  service.
- [PostgreSQL](https://www.postgresql.org/): Database for storing feed and post
  data.
- [Chi Router](https://github.com/go-chi/chi): Lightweight, idiomatic and
  composable router for building Go HTTP services.
- [UUID](https://github.com/google/uuid): For generating unique identifiers for
  feeds and posts.
- [CORS](https://github.com/go-chi/cors): Middleware for handling Cross-Origin
  Resource Sharing.
