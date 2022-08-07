[![codecov](https://codecov.io/gh/preimmortal/sidecar-backup/branch/main/graph/badge.svg?token=63F6CP8ND0)](https://codecov.io/gh/preimmortal/sidecar-backup)

# Sidecar-Backup

Sidecar-Backup is a backup tool meant to be used for syncing directories (remote or local) and creating safe backups of sqlite3 databases. It is meant to run as a [sidecar container][1] in a Kubernetes Pod to handle backing up directories to remote targets and backing sqlite3 databases safely without locking.


# Getting Started


## Docker
To get started with docker, pull the latest image and configure the container
```
docker pull ghcr.io/preimmortal/sidecar-backup:latest
```

Create the [config.yaml][3]
```yaml
enable: true
interval: 0
workers: 1
verbose: false

sql:
  - name: example-sql-1
    source: source/test.db
    dest: source/test.backup.db
    enable: true

rsync:
  - name: example-source-1
    source: source/
    dest: backup
    options:
      exclude:
        - "*ignore*"
    enable: true
```

Run the docker container
```
docker run \
    -v /path/to/source:/source \
    -v /path/to/dest:/dest \
    -v /path/to/config:/config \
    -e "CONFIG=/config/config.yaml" \
    ghcr.io/preimmortal/sidecar-backup
```

## Docker-Compose
```yaml
version: "2.1"
services:
  backup:
    image: ghcr.io/preimmortal/sidecar-backup
    container_name: mybackup
    environment:
      - PUID=1000
      - PGID=1000
      - CONFIG=/config/config.yaml
    volumes:
      - /path/to/source:/source
      - /path/to/dest:/dest
      - /path/to/config:/config
```

Run the Docker-Compose File
```
docker-compose up
```

## Kubernetes Deployment
* TODO


# Detailed Resources

### [Command Line Options][2]
### [Configuration File][3]

[1]: https://kubernetes.io/docs/tasks/access-application-cluster/communicate-containers-same-pod-shared-volume/
[2]: https://github.com/preimmortal/sidecar-backup/blob/main/README-cmdline.md
[3]: https://github.com/preimmortal/sidecar-backup/blob/main/README-config.md