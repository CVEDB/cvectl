## cvectl check diff

Create a diff comparing proposed apk changes following a melange build, to the latest available in an APKINDEX

### Usage

```
cvectl check diff
```

### Synopsis

Create a diff comparing proposed apk changes following a melange build, to the latest available in an APKINDEX

### Options

```
      --apk-index-url string       apk-index-url used to get existing apks.  Defaults to cvedb (default "https://packages.cvedb.dev/os/%s/APKINDEX.tar.gz")
      --dir string                 directory the command is executed from and will contain the resulting diff.log file (default ".")
  -h, --help                       help for diff
      --package-list-file string   name of the package to compare (default "packages.log")
      --packages-dir string        directory containing new packages (default "packages")
```

### SEE ALSO

* [cvectl check](cvectl_check.md)	 - Subcommands used for CI checks in Cvedb

