package cmd

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/arafato/ali-purge/logging"
	"github.com/spf13/cobra"
)

const (
	alibabaCloudRegionId        = "ALIBABACLOUD_REGION_ID"
	alibabaCloudAccessKeyId     = "ALIBABACLOUD_ACCESS_KEY_ID"
	alibabaCloudAccessKeySecret = "ALIBABACLOUD_ACCESS_KEY_SECRET"
)

const ALICLOUD_REGION_ID = "ALIBABACLOUD_REGION_ID"
const ALIBABACLOUD_ACCESS_KEY_ID = "ALIBABACLOUD_ACCESS_KEY_ID"
const ALIBABACLOUD_ACCESS_KEY_SECRET = "ALIBABACLOUD_ACCESS_KEY_SECRET"

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
		// TODO: Initialize Alicloud credentials
		// if os.LookupEnv(alibabaCloudRegionId) || os.LookupEnv(alibabaCloudAccessKeyId) || os.LookupEnv(alibabaCloudAccessKeySecret) {
		// 	logging.Info("Credentials are not configured via environment variables, aborting.")
		// 	os.Exit(-1)
		// }
		regionId := os.Getenv(alibabaCloudRegionId)
		accessKey := os.Getenv(alibabaCloudAccessKeyId)
		accessKeySecret := os.Getenv(ALIBABACLOUD_ACCESS_KEY_SECRET)
		if regionId == "" || accessKey == "" || accessKeySecret == "" {
			logging.Info("Credentials are not configured via environment variables, aborting.")
			os.Exit(-1)
		}

		client, err := sts.NewClientWithAccessKey(regionId, accessKey, accessKeySecret)
		if err != nil {
			logging.Info("Error initializing Alicloud client library: " + err.Error())
			os.Exit(-1)
		}
		// TODO: Read configuration file
		// TODO: Start purging
		p := NewPurge(client)
		return p.Run()
	}

	command.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false,
		"Enables debug output.")

	command.AddCommand(NewVersionCommand())
	return command
}
