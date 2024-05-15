## cvectl advisory discover

Automatically create advisories by matching distro packages to vulnerabilities in NVD

### Usage

```
cvectl advisory discover
```

### Synopsis

Automatically create advisories by matching distro packages to vulnerabilities in NVD

### Options

```
  -a, --advisories-repo-dir string   directory containing the advisories repository
  -d, --distro-repo-dir string       directory containing the distro repository
  -h, --help                         help for discover
      --no-distro-detection          do not attempt to auto-detect the distro
      --nvd-api-key string           NVD API key (Can also be set via the environment variable 'CVECTL_NVD_API_KEY'. Using an API key significantly increases the rate limit for API requests. If you need an NVD API key, go to https://nvd.nist.gov/developers/request-an-api-key.)
  -p, --package string               package name
  -r, --package-repo-url string      URL of the APK package repository
```

### SEE ALSO

* [cvectl advisory](cvectl_advisory.md)	 - Commands for consuming and maintaining security advisory data

