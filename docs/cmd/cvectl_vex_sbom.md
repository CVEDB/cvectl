## cvectl vex sbom

Generate a VEX document from cvedb packages listed in an SBOM

### Usage

```
cvectl vex sbom [flags] sbom.spdx.json
```

### Synopsis

cvectl vex sbom: Generate a VEX document from cvedb packages listed in an SBOM
		
The vex sbom subcommand generates VEX documents describing how vulnerabilities
impact Cvedb packages listed in an SBOM. This subcommand reads SPDX SBOMs and
will recognize and capture all packages identified as Cvedb OS components 
by its purl. For example, if an SBOM contains a package with the following
purl:

	pkg:apk/cvedb/curl@7.87.0-r0
	
cvectl will read the melange configuration file that created the package and
create a VEX document containing impact assessments in its advisories and
secfixes.

cvectl will read the melange config files from an existing cvedb/os clone
or, if not specified, it will clone the repo for you.


### Examples

cvectl vex sbom --author=joe@doe.com sbom.spdx.json

### Options

```
      --author string   author of the VEX document
  -h, --help            help for sbom
      --repo string     path to a local clone of the cvedb/os repo
      --role string     role of the author of the VEX document
```

### SEE ALSO

* [cvectl vex](cvectl_vex.md)	 - Tools to generate VEX statements for Cvedb packages and images

