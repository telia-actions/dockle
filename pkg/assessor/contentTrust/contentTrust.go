package contentTrust

import (
	"os"

	"github.com/knqyf263/fanal/extractor"
	"github.com/tomoyamachi/docker-guard/pkg/types"
)

type ContentTrustAssessor struct{}

func (a ContentTrustAssessor) Assess(fileMap extractor.FileMap) ([]types.Assessment, error) {
	if os.Getenv("DOCKER_CONTENT_TRUST") != "1" {
		return []types.Assessment{
			{
				Type:     types.UseContentTrust,
				Filename: "ENVIRONMENT variable",
				Desc:     " export DOCKER_CONTENT_TRUST=1 before docker build",
			},
		}, nil
	}
	return nil, nil
}

func (a ContentTrustAssessor) RequiredFiles() []string {
	return []string{}
}
