## cvectl dot

Generate graphviz .dot output

### Usage

```
cvectl dot
```

### Synopsis


Generate .dot output and pipe it to dot to generate an SVG

  cvectl dot zlib | dot -Tsvg > graph.svg

Generate .dot output and pipe it to dot to generate a PNG

  cvectl dot zlib | dot -Tpng > graph.png

Open browser to explore crane

  cvectl dot --web crane

Open browser to explore crane's deps recursively, only showing a minimum subgraph

  cvectl dot --web -R -S crane


### Options

```
  -d, --dir string                  directory to search for melange configs (default ".")
  -h, --help                        help for dot
  -k, --keyring-append strings      path to extra keys to include in the build environment keyring (default [https://packages.cvedb.dev/os/cvedb-signing.rsa.pub])
      --pipeline-dir string         directory used to extend defined built-in pipelines
  -R, --recursive                   recurse through package dependencies
  -r, --repository-append strings   path to extra repositories to include in the build environment (default [https://packages.cvedb.dev/os])
  -D, --show-dependents             show packages that depend on these packages, instead of these packages' dependencies
  -S, --spanning-tree               does something like a spanning tree to avoid a huge number of edges
      --web                         do a website
```

### SEE ALSO

* [cvectl](cvectl.md)	 - A CLI helper for developing Cvedb

