# Command Line Options
The command line parameters of sidecar-backup are very simple, requiring only a config specification and debug option.

**CLI Options**
```
Usage of sidecar-backup:
  --config  required  string  config file location 
  -d        optional          debug flag
  -h        optional          help flag
```

**CLI Examples**
```bash
sidecar-backup --config /tmp/config.yaml
sidecar-backup --config /tmp/config.yaml -d
```

These command line options can also be specified using environment variables.

**ENV Options**
```
Usage of sidecar-backup environment configuration:
CONFIG  required  config file location
DEBUG   optional  debug flag
```
**ENV Examples**
```bash
CONFIG=/tmp/config.yaml sidecar-backup
CONFIG=/tmp/config.yaml DEBUG=true sidecar-backup
```