## cvectl gh release

Performs a GitHub release using git tags to calculate the release version

### Usage

```
cvectl gh release
```

### Synopsis

Performs a GitHub release using git tags to calculate the release version

Examples:

cvectl gh release --bump-major
cvectl gh release --bump-minor
cvectl gh release --bump-patch
cvectl gh release --bump-prerelease-with-prefix rc


### Options

```
      --bump-major                           bumps the major release version
      --bump-minor                           bumps the minor release version
      --bump-patch                           bumps the patch release version
      --bump-prerelease-with-prefix string   bumps the prerelease version using the supplied prefix, if no existing prerelease exists the patch version is also bumped to align with semantic versioning
      --dir string                           directory containing the cloned github repository to release (default ".")
  -h, --help                                 help for release
```

### SEE ALSO

* [cvectl gh](cvectl_gh.md)	 - Commands used to interact with GitHub

