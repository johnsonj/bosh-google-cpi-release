package action

import (
	"fmt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	"github.com/frodenas/bosh-google-cpi/google/image_service"
)

const googleInfrastructure = "google"

type CreateStemcell struct {
	stemcellService image.Service
}

func NewCreateStemcell(
	stemcellService image.Service,
) CreateStemcell {
	return CreateStemcell{
		stemcellService: stemcellService,
	}
}

func (cs CreateStemcell) Run(stemcellPath string, cloudProps StemcellCloudProperties) (StemcellCID, error) {
	var err error
	var description, stemcell string

	if cloudProps.Infrastructure != googleInfrastructure {
		return "", bosherr.Errorf("Creating stemcell: Invalid '%s' infrastructure", cloudProps.Infrastructure)
	}

	if cloudProps.Name != "" && cloudProps.Version != "" {
		description = fmt.Sprintf("%s/%s", cloudProps.Name, cloudProps.Version)
	}

	if cloudProps.SourceURL != "" {
		stemcell, err = cs.stemcellService.CreateFromURL(cloudProps.SourceURL, description)
	} else {
		stemcell, err = cs.stemcellService.CreateFromTarball(stemcellPath, description)
	}
	if err != nil {
		return "", bosherr.WrapError(err, "Creating stemcell")
	}

	return StemcellCID(stemcell), nil
}
