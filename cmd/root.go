package cmd

import (
	"log"
	"os"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRootCommand() *cobra.Command {
	var verbose bool
	var logger *zap.SugaredLogger

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
		loggerConfig := zap.NewProductionConfig()
		loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

		if verbose {
			loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		} else {
			loggerConfig.Level = zap.NewAtomicLevel() //Info Level by default
		}

		_logger, err := loggerConfig.Build()
		if err != nil {
			log.Fatal(err)
		}

		defer _logger.Sync()
		logger = _logger.Sugar()
		logger.Debug("Logger initialized")
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
		// TODO: Initialize Alicloud credentials
		client, err := sts.NewClientWithAccessKey(os.Getenv("ALIBABACLOUD_REGION_ID"), os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET"))
		if err != nil {
			logger.Info("Credentials are not configured via environment variables, aborting.")
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
