## cvectl text

Print a sorted list of downstream dependent packages

### Usage

```
cvectl text
```

### Synopsis

Print a sorted list of downstream dependent packages

### Options

```
  -a, --arch string                 architecture to build for (default "x86_64")
  -d, --dir string                  directory to search for melange configs (default ".")
  -h, --help                        help for text
  -k, --keyring-append strings      path to extra keys to include in the build environment keyring (default [https://packages.cvedb.dev/os/cvedb-signing.rsa.pub])
      --pipeline-dir string         directory used to extend defined built-in pipelines
  -r, --repository-append strings   path to extra repositories to include in the build environment (default [https://packages.cvedb.dev/os])
  -t, --type string                 What type of text to emit; values can be one of: [target makefile name version name-version] (default "target")
```

### SEE ALSO

* [cvectl](cvectl.md)	 - A CLI helper for developing Cvedb

