# Expenses REST API Golang
This is a simple REST json API written in Golang for learning purposes.

# Install App
We use docker to mount Postgresql & PgAdmin. Check `docker-compose.yml` for further information.
- `cp .env.default .env`
- Edit .env and set variables
- Run `docker-compose up` to start servers
- Run migrations `docker run -v <ABSOLUTE PATH>/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database <CONNECTION STRING> up`
- Run `go run main.go`

# Access pqAdmin
- Go to `http://localhost:8001` with credentials set in docker-compose.yml file.


# Common issues

## Can't login to the database or it's named incorrectly
After changign env variables related to postgres server, you should clear container volumes, to do so:

```
docker-compose down -v
docker-compose up --force-recreate
```

## Author
Mariano Peyregne