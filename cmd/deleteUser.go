package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Allows user to delete himself or herself.",
	Run: func(cmd *cobra.Command, args []string) {
		cognito, err := service.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		err = cognito.DeleteUser(cmd.Flag("token").Value.String())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(SuccessfulMessage)
		}
	},
}

func init() {
	deleteUserCmd.Flags().String("token", "", "access token")

	_ = deleteUserCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(deleteUserCmd)
}
