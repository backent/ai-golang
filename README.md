# Helper command
## Add migration file
```bash
docker run -it -v $(pwd)/database/migrations:/migrations migrate/migrate create -ext sql -dir /migrations add_l
anguage_column_to_questions_table
```

## Run migration
```bash
docker run -it -v $(pwd)/database/migrations:/migrations --network global-network migrate/migrate -path=/migrations -database "mysql://root:adminlocal@tcp(global-mysql:3306)/ai" up 1
```