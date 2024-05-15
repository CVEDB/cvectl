## cvectl check so-name

Check so name files have not changed in upgrade

### Usage

```
cvectl check so-name
```

### Synopsis

Check so name files have not changed in upgrade

### Options

```
      --apk-index-url string       apk-index-url used to get existing apks.  Defaults to cvedb (default "https://packages.cvedb.dev/os/aarch64/APKINDEX.tar.gz")
  -d, --directory string           directory containing melange configs (default ".")
  -h, --help                       help for so-name
      --package-list-file string   name of the package to compare (default "packages.log")
      --package-name stringArray   override using package-list-file and specify a single package name to compare
      --packages-dir string        directory containing new packages (default "packages")
```

### SEE ALSO

* [cvectl check](cvectl_check.md)	 - Subcommands used for CI checks in Cvedb

