package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var deletePoolClientCmd = &cobra.Command{
	Use:   "delete-pool-client",
	Short: "Deletes the user pool client.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := cognito.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		err = srv.DeletePoolClient(cmd.Flag("poolID").Value.String(), cmd.Flag("clientID").Value.String())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("✅ Operation has been successful!")
		}
	},
}

func init() {
	deletePoolClientCmd.Flags().String("poolID", "", "")
	deletePoolClientCmd.Flags().String("clientID", "", "")

	_ = deletePoolClientCmd.MarkFlagRequired("poolID")
	_ = deletePoolClientCmd.MarkFlagRequired("clientID")

	rootCmd.AddCommand(deletePoolClientCmd)
}
