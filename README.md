![alt tag](https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png)
###### Author : Haocheng Zhang
  Welcome to **BPR Golang Gin Home Server** 

A restful api's with [Gin Framework](https://github.com/gin-gonic/gin/) with a structured project that defaults to **PostgreSQL** database and **JWT** authentication middleware stored in **Redis**

##  with

- [go-gorp](https://github.com/go-gorp/gorp): Go Relational Persistence
- [jwt-go](https://github.com/golang-jwt/jwt): JSON Web Tokens (JWT) as middleware
- [go-redis](https://github.com/go-redis/redis): Redis support for Go
- Go Modules
-  **Custom Validators**
-  **CORS Middleware**
-  **RequestID Middleware**
- Feature **PostgreSQL 12** with JSON/JSONB queries & trigger functions
- Enviroment support
- Unit test


### Installation gudie

```
$ go mod init
```

```
$ go install
```

You will find the **home-database.sql** in `db/home-database.sql`

And you can import the postgres database using this command:

```
$ psql -U postgres -h localhost < ./db/home-database.sql
```



## Running Home Application

Rename .env_rename_me to .env and place your credentials

```
$ mv .env_rename_me .env
```


## Building Home Application

```
$ go build -v
```

## Testing Home Application

```
$ go test -v ./tests/*
```



