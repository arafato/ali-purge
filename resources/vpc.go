package resources

import (
	"github.com/arafato/ali-purge/cmd"
	"github.com/arafato/ali-purge/logging"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

func init() {
	GetServiceManager().register("vpc", ListVpcs)
}

type Vpc struct {
	client *vpc.Client
	id     string
	name   string
}

func ListVpcs(config *cmd.AlicloudConfig) ([]cmd.Resource, error) {
	resources := []cmd.Resource{}

	client, err := vpc.NewClientWithOptions(config.RegionId, config.Config, config.Creds)
	if err != nil {
		return nil, err
	}

	request := vpc.CreateDescribeVpcsRequest()
	request.Scheme = "https"

	response, err := client.DescribeVpcs(request)
	if err != nil {
		logging.Error("Error calling DescribeVpcs")
		return nil, err
	}

	for _, vpc := range response.Vpcs.Vpc {
		resources = append(resources, &Vpc{client: client, id: vpc.VpcId, name: vpc.VpcName})
	}

	return resources, nil
}

func (vpc *Vpc) Remove() error {
	return nil
}
