# golang-resetfull-api-gin

Create new migration file

```
migrate create -ext sql -dir migrations -seq create_users_table
```

Run migration

```
migrate -database ${POSTGRESQL_URL} -path migrations up number_count_version
```

Reverse migration

```
migrate -database ${POSTGRESQL_URL} -path migrations down number_count_version

In case migrate database fail or have error

```

error: Dirty database version 16. Fix and force version.

```

Step 1: force migrate database to last successfull version

```

migrate -database ${POSTGRESQL_URL} -path migrations force last_successfull_version

```

Step 2: Fix migration file
Step 3: Migrate database again

```

migrate -database ${POSTGRESQL_URL} -path migrations up

```

### Run app

Run app

```

go run main.go run-app
```
