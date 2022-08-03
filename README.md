# sidecar-backup
A backup tool meant to be used as a sidecar container

# Dev Run
```
go run cmd/sidecar-backup/main.go --config="example/config.yaml"  --workers=1
```

# Test with Coverage
```
go test ./... -coverprofile=coverage.out
```