## cvectl vex

Tools to generate VEX statements for Cvedb packages and images

### Synopsis

cvectl vex: Tools to generate VEX statements for Cvedb packages and images
		
The vex family of subcommands interacts with Cvedb data and configuration
files to generate Vulnerability Exploitability eXchange (VEX) documents to
inform downstream consumer how vulnerabilities impact Cvedb packages and images
that use them. 

cvectl can generate VEX data by reading the melange configuration files
of each package and additional information coming from external documents.
There are currently two VEX subcommands:

 cvectl vex package: Generates VEX documents from a list of melange configs

 cvectl vex sbom: Generates a VEX document by reading an image SBOM

For more information please see the help sections if these subcommands. To know
more about the VEX tooling powering cvectl see: https://openvex.dev/




### Options

```
  -h, --help   help for vex
```

### SEE ALSO

* [cvectl](cvectl.md)	 - A CLI helper for developing Cvedb
* [cvectl vex package](cvectl_vex_package.md)	 - Generate a VEX document from package configuration files
* [cvectl vex sbom](cvectl_vex_sbom.md)	 - Generate a VEX document from cvedb packages listed in an SBOM

