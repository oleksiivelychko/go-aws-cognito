package cmd

import (
	"github.com/oleksiivelychko/go-aws-cognito/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const SuccessfulMessage = "✅ Operation has been successful!"

var (
	cfgFile string
	cfgAWS  *config.AWS
)

var rootCmd = &cobra.Command{
	Short: "Cognito user pools API allows set up user pools and app clients, and authenticate users.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config YAML file")
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		cobra.CheckErr(err)
	}

	cfgAWS = &config.AWS{
		Region:             viper.Get("REGION").(string),
		AwsAccessKeyId:     viper.Get("AWS_ACCESS_KEY_ID").(string),
		AwsSecretAccessKey: viper.Get("AWS_SECRET_ACCESS_KEY").(string),
		Endpoint:           viper.Get("ENDPOINT").(string),
	}
}
