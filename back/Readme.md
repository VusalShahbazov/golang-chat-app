### Migration tool


```
migrate -path ./database/migrations  -database "mysql://root:@tcp(localhost:3306)/go_chat_app?multiStatements=true" up
```

```
migrate create -ext sql -dir ./database/migrations <migration_name>
```


```
go build cmd/api/main.go && main.exe
```
