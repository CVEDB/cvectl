## cvectl image apk

Show APK(s) in a container image

### Usage

```
cvectl image apk <image>
```

### Synopsis

Show APK(s) in a container image

### Examples


  # Show all APKs in an image
  cvectl image apk cgr.dev/chainguard/bash

  # Show all APKs in an image that own a component (based on a Syft analysis)
  cvectl image apk cgr.dev/chainguard/coredns -c 'github.com/aws/aws-sdk-go'

  # Show all APKs in an image that own a component, and show the path to the
  # Melange configuration file for each APK, within the given directory
  cvectl image apk cgr.dev/chainguard/prometheus-operator -c 'github.com/aws/aws-sdk-go' -d ~/code/cvedb-os


### Options

```
  -c, --component string     show only APKs containing the given component
  -d, --distro-dir strings   path to a directory containing Melange build configuration files
  -h, --help                 help for apk
```

### SEE ALSO

* [cvectl image](cvectl_image.md)	 - (Experimental) Commands for working with container images that use Cvedb

