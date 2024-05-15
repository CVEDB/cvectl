# cvectl

[![Documentation](https://godoc.org/github.com/cvedb/cvectl?status.svg)](https://pkg.go.dev/mod/github.com/cvedb/cvectl)
[![Go Report Card](https://goreportcard.com/badge/github.com/cvedb/cvectl)](https://goreportcard.com/report/github.com/cvedb/cvectl)

`cvectl` is a command line tool for working with Cvedb

## Installation

You can install  `cvectl` straight from its source code. To do this, clone the git repository and then run `go install`:

```bash
# Clone the repo

git clone git@github.com:cvedb/cvectl.git cvectl && cd $_

# Install the `cvectl` command

go install
```

## Commands

See the [cvectl command reference](https://github.com/cvedb/cvectl/blob/main/docs/cmd/cvectl.md)

## Docs

[Check so_name docs](./docs/check_so_name.md) - CI check for detecting ABI breaking changes in package version updates
[Update docs](./docs/update.md) - for detecting new upstream cvedb package versions and creating a pull request to update Cvedb

## Releases

This repo is configured to automatically create weekly tagged patch releases, mainly so that it can be more easily packaged in Cvedb itself.

Releases happen Monday at 00:00 UTC, and can be manually run as necessary.
