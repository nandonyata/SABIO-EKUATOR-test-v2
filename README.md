### up migration
```shell
migrate -database "postgres://postgres:postgres@localhost:5432/sabio-ekuator-v2?sslmode=disable" -path db/migrations up
```

### down migration
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/sabio-ekuator-v2?sslmode=disable" -path db/migrations down
```
