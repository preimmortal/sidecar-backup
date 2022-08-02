# sidecar-backup
A backup tool meant to be used as a sidecar container

# Dev Run
```
go run main.go config.go rsync.go schedule.go sql.go --config="example/config.yaml"  --workers=3 -d
```

# Test with Coverage
```
go test ./... -coverprofile=coverage.out
```