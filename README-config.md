# Config File
The config file is specified in [YAML][1] and consists of base options as well as rsync and sqlite3 options.

## Example Configuration File
```yaml
enable: true
interval: 0
workers: 1
verbose: false

sql:
  - name: example-sql-dne
    source: example/src/dne.db
    dest: example/src/dne.backup.db
    enable: true

  - name: example-sql-1
    source: /source/test.db
    dest: /source/test.backup.db
    enable: true

rsync:
  - name: example-source-1
    source: /source/
    dest: /dest
    options:
      exclude:
        - "*ignore*"
      delete: true
      deleteduring: true
      deleteexcluded: true
      force: true
    enable: true

  - name: example-dne-1
    source: /srcdne/
    dest: /destdne
    enable: true

  - name: example-disabled-1
    source: /source/
    dest: /dest
    enable: false
```

### Base
```
enable    bool  false  enable sidecar-backup
interval  int   0      backup interval in seconds
workers   int   1      number of workers to run concurrently
verbose   bool  false  set verbosity of tool
```

### Rsync
Rsync backup uses the grsync project. Credit to [zloylos/grsync][2]
```
rsync        []map
    name     string  none   name of the rsync backup job
    source   string  none   source location
    dest     string  none   destination location
    options  []map   none   options, see all options below
    enable   bool    false  enable this rsync backup job
```
### Rsync Options
Rsync Options are directly converted to `rsync` options, please read the [manpage][3] to use them.
```
rsync-options        []map
	verbose          bool    --verbose
	quiet            bool    --quiet
	checksum         bool    --checksum
	archive          bool    --archive
	recursive        bool    --recursive
	relative         bool    --relative
	noimplieddirs    bool    --no-implied-dirs
	update           bool    --update
	inplace          bool    --inplace
	append           bool    --append
	appendverify     bool    --append-verify
	dirs             bool    --dirs
	links            bool    --links
	copylinks        bool    --copy-links
	copyunsafelinks  bool    --copy-unsafe-links
	safelinks        bool    --safe-links
	copydirlinks     bool    --copy-dir-links
	keepdirlinks     bool    --keep-dir-links
	hardlinks        bool    --hard-links
	perms            bool    --perms
	noperms          bool    --no-perms
	executability    bool    --executability
	acls             bool    --acls
	xattrs           bool    --xattrs
	owner            bool    --owner
	noowner          bool    --no-owner
	group            bool    --group
	nogroup          bool    --no-group
	devices          bool    --devices
	specials         bool    --specials
	times            bool    --times
	omitdirtimes     bool    --omit-dir-times
	super            bool    --super
	fakesuper        bool    --fake-super
	sparse           bool    --sparse
	dryrun           bool    --dry-run
	wholefile        bool    --whole-file
	onefilesystem    bool    --one-file-system
	blocksize        int     --block-size
	rsh              string  --rsh
	existing         bool    --existing
	ignoreexisting   bool    --ignore-existing
	removesourcefiles bool   --remove-source-files
	delete           bool    --delete
	deletebefore     bool    --delete-before
	deleteduring     bool    --delete-during
	deletedelay      bool    --delete-delay
	deleteafter      bool    --delete-after
	deleteexcluded   bool    --delete-excluded
	ignoreerrors     bool    --ignore-errors
	force            bool    --force
	maxdelete        int     --max-delete
	maxsize          int     --max-size
	minsize          int     --min-size
	partial          bool    --partial
	partialdir       string  --partial-dir
	delayupdates     bool    --delay-updates
	pruneemptydirs   bool    --prune-empty-dirs
	numericids       bool    --numeric-ids
	timeout          int     --timeout
	contimeout       int     --contimeout
	ignoretimes      bool    --ignore-times
	sizeonly         bool    --size-only
	modifywindow     bool    --modify-window
	tempdir          string  --temp-dir
	fuzzy            bool    --fuzzy
	comparedest      string  --compare-dest
	copydest         string  --copy-dest
	linkdest         string  --link-dest
	compress         bool    --compress
	compresslevel    int     --compress-level
	skipcompress     []string  --skip-compress
	cvsexclude       bool    --cvs-exclude
	stats            bool    --stats
	humanreadable    bool    --human-readable
	progress         bool    --progress
	passwordfile     string  --password-file
	bandwidthlimit   int     --bw-limit
	info             string  --info
	exclude          []string  --exclude
	include          []string  --include
	filter           string  --filter
	chown            string  --chown
	ipv4             bool    --ipv4
	ipv6             bool    --ipv6
	outformat        bool    --outformat
```

### Sql
```
sql:         []map
	name     string  ""     name of the sqlite3 backup job
	source   string  ""     source location
	dest     string  ""     destination location
	options  string  ""     options
	enable   bool    false  enable this sqlite3 backup job
```
* There are currently no options to be specified for sql


[1]: https://yaml.org/spec/1.2.2/
[2]: https://github.com/zloylos/grsync
[3]: https://www.unix.com/man-page/Linux/1/rsync/