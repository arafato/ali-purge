package resources

import (
	"github.com/arafato/ali-purge/cmd"
	"github.com/arafato/ali-purge/logging"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

func init() {
	GetServiceManager().register("vpc-vswitch", ListVSwitches)
}

type VSwitch struct {
	client *vpc.Client
	id     string
	name   string
}

func ListVSwitches(config *cmd.AlicloudConfig) ([]cmd.Resource, error) {
	resources := []cmd.Resource{}

	client, err := vpc.NewClientWithOptions(config.RegionId, config.Config, config.Creds)
	if err != nil {
		return nil, err
	}

	request := vpc.CreateDescribeVSwitchesRequest()
	request.Scheme = "https"

	response, err := client.DescribeVSwitches(request)
	if err != nil {
		logging.Error("Error calling DescribeVSwitches")
		return nil, err
	}

	for _, vswitch := range response.VSwitches.VSwitch {
		resources = append(resources, &VSwitch{client: client, id: vswitch.VSwitchId, name: vswitch.VSwitchName})
	}

	return resources, nil
}

func (vpc *VSwitch) Remove() error {
	return nil
}
