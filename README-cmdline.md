# Command Line Options
The command line parameters of sidecar-backup are very simple, requiring only a config specification and debug option.

**CLI Options**
```
Usage of sidecar-backup:
  --config  required  string  config file location 
  -d        optional          debug flag: sets log level to debug
  -f        optional          force flag: forces immediate backup and exit
  -h        optional          help flag: display help message
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
FORCE   optional  force flag
```
**ENV Examples**
```bash
CONFIG=/tmp/config.yaml sidecar-backup
CONFIG=/tmp/config.yaml DEBUG=true sidecar-backup
```
