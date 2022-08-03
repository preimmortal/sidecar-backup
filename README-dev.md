# sidecar-backup
A backup tool meant to be used as a sidecar container

# Dev Run
```
go run cmd/sidecar-backup/main.go --config="example/config.yaml"

CONFIG=example/config.yaml go run cmd/sidecar-backup/main.go
```

# Build
```
go build -v -o . ./...
```

# Test with Coverage
```
go test ./... -coverprofile=coverage.out
```

# Build Local Docker Container
```
docker build -t preimmortal/sidecar-backup .
```