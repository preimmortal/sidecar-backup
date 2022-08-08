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

# Create Mocks
mockgen -source=rsync.go  > mocks/mock-rsync.go
mockgen -source schedule.go > mocks/mock-schedule.go

# Test with Coverage
```
go test -v | go-junit-report > report.xml

go test -coverprofile=cover.out
gocov convert cover.out | gocov-xml > coverage.xml
```

# Build Local Docker Container
```
docker build -t preimmortal/sidecar-backup .
```

# Create Tags/Release
* version names always start with 'v' and are of the format 'vX.Y.Z`
git tag -a "<version>" -m "<message>"
git push origin <version>

Create Release on the Github Release page