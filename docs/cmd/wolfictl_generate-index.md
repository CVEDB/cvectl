## cvectl generate-index



### Usage

```
cvectl generate-index
```

### Synopsis

This command generates an APKINDEX from the contents of a remote bucket.

Specify the bucket with --bucket. The default is "cvedb", the main prod Cvedb bucket.
Other acceptable values include "stage1", "stage2" and "stage3" for the bootstrap buckets.
Otherwise, specify any GCS bucket location with the gs:// prefix.

If --signing-key is passed, the APKINDEX will be signed with that key.

If --publish is passed, the APKINDEX will be published back to the bucket.
Otherwise it's written to APKINDEX.tar.gz.


### Options

```
      --arch string          arch of package to get (default "x86_64")
      --bucket string        bucket to get packages from (default "cvedb")
  -h, --help                 help for generate-index
      --publish              if true, publish APKINDEX.tar.gz back to the repo (must be signed)
      --signing-key string   if set, key to use to sign the index
```

### SEE ALSO

* [cvectl](cvectl.md)	 - A CLI helper for developing Cvedb

