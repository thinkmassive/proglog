
## Testing

1. Start server
```
go run cmd/server/main.go
```

2. Create records
```
curl -X POST http://localhost:8080 -d '{"record":{"value":"TGV0J3MgR28gIzEK"}}'
curl -X POST http://localhost:8080 -d '{"record":{"value":"TGV0J3MgR28gIzIK"}}'
curl -X POST http://localhost:8080 -d '{"record":{"value":"TGV0J3MgR28gIzMK"}}'
```

3. Read records

```
curl -X GET http://localhost:8080 -d '{"offset":0}'
curl -X GET http://localhost:8080 -d '{"offset":1}'
curl -X GET http://localhost:8080 -d '{"offset":2}'
curl -X GET http://localhost:8080 -d '{"offset":9999}'
```
