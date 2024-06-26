package cli

import (
	"log"
	"os"
	"time"

	"github.com/google/go-github/v58/github"
	"github.com/spf13/cobra"
	"github.com/cvedb/cvectl/pkg/http"
	"github.com/cvedb/cvectl/pkg/update"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
)

func Package() *cobra.Command {
	o := update.PackageOptions{
		Logger: log.New(log.Writer(), "cvectl update: ", log.LstdFlags|log.Lmsgprefix),
	}

	cmd := &cobra.Command{
		Use:     "package",
		Short:   "Proposes a single melange package update via a pull request",
		Long:    `"Proposes a single melange package update via a pull request".`,
		Example: `cvectl update package cheese --version v1.2.3 --target-repo https://github.com/cvedb/os`,
		Args:    cobra.RangeArgs(1, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !o.DryRun {
				if _, err := (ghTokenSource{}).Token(); err != nil {
					return err
				}
			}

			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
			)

			ratelimit := &http.RLHTTPClient{
				Client: oauth2.NewClient(cmd.Context(), ts),

				// 1 request every (n) second(s) to avoid DOS'ing server. https://docs.github.com/en/rest/guides/best-practices-for-integrators?apiVersion=2022-11-28#dealing-with-secondary-rate-limits
				Ratelimiter: rate.NewLimiter(rate.Every(3*time.Second), 1),
			}

			o.GithubClient = github.NewClient(ratelimit.Client)

			o.PackageName = args[0]
			return o.UpdatePackageCmd(cmd.Context())
		},
	}

	cmd.Flags().BoolVar(&o.DryRun, "dry-run", false, "prints proposed package updates rather than creating a pull request")
	cmd.Flags().BoolVar(&o.Advisories, "sec-fixes", true, "checks commit messages since last release, for `fixes: CVE###` and generates melange security advisories")
	cmd.Flags().StringVar(&o.PullRequestBaseBranch, "pull-request-base-branch", "main", "base branch to create a pull request against")
	cmd.Flags().StringVar(&o.TargetRepo, "target-repo", "https://github.com/cvedb/os", "target git repository containing melange configuration to update")
	cmd.Flags().StringVar(&o.Version, "version", "", "version to bump melange package to")
	cmd.Flags().StringVar(&o.Epoch, "epoch", "0", "the epoch used to identify fix, defaults to 0 as this command is expected to run in a release pipeline that's creating a new version so epoch will be 0")
	cmd.Flags().BoolVar(&o.UseGitSign, "use-gitsign", false, "enable gitsign to sign the git commits")

	return cmd
}
