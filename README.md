# sidecar-backup
A backup tool meant to be used as a sidecar container

# Dev Run
```
go run cmd/sidecar-backup/main.go --config="example/config.yaml"
```

# Build
```
go build -v -o . ./...
```

# Test with Coverage
```
go test ./... -coverprofile=coverage.out
```