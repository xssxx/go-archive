# How to run server

### Build postgresql database

```sh
cd server
docker-compose up --build
```

### Install dependencies

```sh
go mod download
```

### Run server

```sh
go run main.go
```
