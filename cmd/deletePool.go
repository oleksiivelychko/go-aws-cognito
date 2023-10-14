package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var cognitoDeletePoolCmd = &cobra.Command{
	Use:   "delete-pool",
	Short: "Deletes the user pool.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := cognito.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		err = srv.DeletePool(cmd.Flag("poolID").Value.String())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(SuccessfulMessage)
		}
	},
}

func init() {
	cognitoDeletePoolCmd.Flags().String("poolID", "", "")

	_ = cognitoDeletePoolCmd.MarkFlagRequired("poolID")

	rootCmd.AddCommand(cognitoDeletePoolCmd)
}
