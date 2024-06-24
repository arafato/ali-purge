package cmd

import (
	"os"
	"time"

	"github.com/arafato/ali-purge/logging"
	"github.com/spf13/cobra"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
)

const (
	alibabaCloudRegionId        = "ALIBABACLOUD_REGION_ID"
	alibabaCloudAccessKeyId     = "ALIBABACLOUD_ACCESS_KEY_ID"
	alibabaCloudAccessKeySecret = "ALIBABACLOUD_ACCESS_KEY_SECRET"
)

type AlicloudConfig struct {
	Config   *sdk.Config
	Creds    *credentials.AccessKeyCredential
	RegionId string
}

func NewRootCommand() *cobra.Command {
	var verbose bool

	command := &cobra.Command{
		Use:   "ali-purge",
		Short: "ali-purge removes every resource from your Alibaba Cloud account.",
		Long: `
 @@@@@@   @@@       @@@  @@@@@@@   @@@  @@@  @@@@@@@    @@@@@@@@  @@@@@@@@  
 @@@@@@@@  @@@       @@@  @@@@@@@@  @@@  @@@  @@@@@@@@  @@@@@@@@@  @@@@@@@@  
 @@!  @@@  @@!       @@!  @@!  @@@  @@!  @@@  @@!  @@@  !@@        @@!       
 !@!  @!@  !@!       !@!  !@!  @!@  !@!  @!@  !@!  @!@  !@!        !@!       
 @!@!@!@!  @!!       !!@  @!@@!@!   @!@  !@!  @!@!!@!   !@! @!@!@  @!!!:!    
 !!!@!!!!  !!!       !!!  !!@!!!    !@!  !!!  !!@!@!    !!! !!@!!  !!!!!:    
 !!:  !!!  !!:       !!:  !!:       !!:  !!!  !!: :!!   :!!   !!:  !!:       
 :!:  !:!   :!:      :!:  :!:       :!:  !:!  :!:  !:!  :!:   !::  :!:       
 ::   :::   :: ::::   ::   ::       ::::: ::  ::   :::   ::: ::::   :: ::::  
  :   : :  : :: : :  :     :         : :  :    :   : :   :: :: :   : :: ::   
																			 
 A tool which removes every resource from an Alibaba Cloud account.  
 Use it with caution, since it cannot distinguish between production and non-production!`,
	}

	command.PreRun = func(cmd *cobra.Command, args []string) {
		if verbose {
			logging.SetVerbose(true)
		}
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
		regionId := os.Getenv(alibabaCloudRegionId)
		accessKeyId := os.Getenv(alibabaCloudAccessKeyId)
		accessKeySecret := os.Getenv(alibabaCloudAccessKeySecret)
		if regionId == "" || accessKeyId == "" || accessKeySecret == "" {
			logging.Info("Credentials and region are not configured via environment variables, aborting.")
			os.Exit(-1)
		}

		config := sdk.NewConfig().
			WithTimeout(5 * time.Second).
			WithDebug(true)

		credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
		p := NewPurge(AlicloudConfig{Config: config, Creds: credential, RegionId: regionId})
		return p.Run()
	}

	command.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false,
		"Enables debug output.")

	command.AddCommand(NewVersionCommand())
	return command
}
