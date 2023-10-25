package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var describePoolCmd = &cobra.Command{
	Use:   "describe-pool",
	Short: "Returns a configuration information and metadata of the user pool.",
	Run: func(cmd *cobra.Command, args []string) {
		cognito, err := service.New(configAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		output, err := cognito.DescribePool(cmd.Flag("poolID").Value.String())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(output)
		}
	},
}

func init() {
	describePoolCmd.Flags().String("poolID", "", "")

	_ = describePoolCmd.MarkFlagRequired("poolID")

	rootCmd.AddCommand(describePoolCmd)
}
