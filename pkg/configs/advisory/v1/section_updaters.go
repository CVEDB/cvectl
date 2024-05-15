package v1

import (
	"github.com/cvedb/cvectl/pkg/configs"
)

func NewAdvisoriesSectionUpdater(
	updater configs.SectionUpdater[Advisories, Document],
) configs.EntryUpdater[Document] {
	yamlASTMutater := configs.NewTargetedYAMLASTMutater[Advisories, Document](
		"advisories",
		updater,
		func(cfg Document, data Advisories) Document {
			cfg.Advisories = data
			return cfg
		},
	)

	return configs.NewYAMLUpdateFunc[Document](yamlASTMutater)
}
