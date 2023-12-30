# GoBank

A development project to learn and understand API development in Go without relying on popular packages like gin in order to understand how things work at low level.

## Development

Set environment variable

```
export JWT_SECRET=gobanksecret
```

Run Postgres using Docker

```
docker run --name postgres-gobank -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres
```
