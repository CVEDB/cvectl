## cvectl advisory update

Update an existing advisory with a new event

### Usage

```
cvectl advisory update
```

### Synopsis

Update an existing advisory with a new event.

Use this command to update an existing advisory by adding a new "event" to the
advisory, i.e. when the given package/vulnerability combination already exists
in the advisories repo. If the package/vulnerability combination doesn't yet
exist, use the "create" command instead.

This command will prompt for all required fields, and will attempt to fill in
as many optional fields as possible. You can abort the advisory update at any
point in the prompt by pressing Ctrl+C.

You can specify required values on the command line using the flags relevant to
the advisory event you are adding. If not all required values are provided on
the command line, the command will prompt for the missing values.

If the --no-prompt flag is specified, then the command will fail if any
required fields are missing.

### Options

```
  -a, --advisories-repo-dir string   directory containing the advisories repository
      --arch strings                 package architectures to find published versions for (default [x86_64,aarch64])
  -d, --distro-repo-dir string       directory containing the distro repository
      --fixed-version string         package version where fix was applied (used only for 'fixed' event type)
      --fp-note string               prose explanation of the false positive (used only for false positives)
      --fp-type string               type of false positive [vulnerability-record-analysis-contested, component-vulnerability-mismatch, vulnerable-code-version-not-used, vulnerable-code-not-included-in-package, vulnerable-code-not-in-execution-path, vulnerable-code-cannot-be-controlled-by-adversary, inline-mitigations-exist]
  -h, --help                         help for update
      --no-distro-detection          do not attempt to auto-detect the distro
      --no-prompt                    do not prompt the user for input
  -p, --package string               package name
  -r, --package-repo-url string      URL of the APK package repository
      --timestamp string             timestamp of the event (RFC3339 format) (default "now")
      --tp-note string               prose explanation of the true positive (used only for true positives)
  -t, --type string                  type of event [detection, true-positive-determination, fixed, false-positive-determination, analysis-not-planned, fix-not-planned, pending-upstream-fix]
  -V, --vuln string                  vulnerability ID for advisory
```

### SEE ALSO

* [cvectl advisory](cvectl_advisory.md)	 - Commands for consuming and maintaining security advisory data
