package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cvedb/cvectl/pkg/advisory"
	"github.com/cvedb/cvectl/pkg/advisory/question"
	"github.com/cvedb/cvectl/pkg/question/graph"
)

func cmdAdvisoryGuideGraph() *cobra.Command {
	return &cobra.Command{
		Use:           "graph",
		Short:         "Generate a DOT graph of the advisory guide interview questions",
		Args:          cobra.NoArgs,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			sampleReq := advisory.Request{
				Package:         "foo",
				VulnerabilityID: "CVE-2024-12345",
			}

			dot, err := graph.Dot(cmd.Context(), question.IsFalsePositive, sampleReq)
			if err != nil {
				return fmt.Errorf("generating DOT: %w", err)
			}

			fmt.Print(dot)

			return nil
		},
	}
}
