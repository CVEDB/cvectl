package advisory

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	v2 "github.com/cvedb/cvectl/pkg/configs/advisory/v2"
	rwos "github.com/cvedb/cvectl/pkg/configs/rwfs/os"
)

func Test_ImportAdvisoriesYAML(t *testing.T) {
	const testdataDir = "./testdata/export/advisories"

	cases := []struct {
		name            string
		pathToInputData string
	}{
		{
			name:            "test-yaml",
			pathToInputData: "./testdata/export/expected.yaml",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			advisoryFsys := rwos.DirFS(testdataDir)
			advisoryDocs, err := v2.NewIndex(context.Background(), advisoryFsys)
			require.NoError(t, err)

			b, err := os.ReadFile(tt.pathToInputData)
			require.NoError(t, err)

			_, importedDocuments, err := ImporAdvisoriesYAML(b)
			require.NoError(t, err)
			require.Equal(t, advisoryDocs.Select().Len(), importedDocuments.Select().Len())
		})
	}
}
