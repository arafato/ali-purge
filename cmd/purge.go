package cmd

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

type Purge struct {
	Client *sts.Client
}

func NewPurge(client *sts.Client) *Purge {
	p := Purge{
		Client: client,
	}
	return &p
}

func (p *Purge) Run() error {
	return nil
}
