package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var createPoolClientCmd = &cobra.Command{
	Use:   "create-pool-client",
	Short: "Creates a user pool client",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := cognito.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		clientID, err := srv.CreatePoolClient(
			cmd.Flag("name").Value.String(),
			cmd.Flag("poolID").Value.String(),
		)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("✅ Pool client ID: %s\n", clientID)
		}
	},
}

func init() {
	createPoolClientCmd.Flags().String("name", "", "user pool client name")
	createPoolClientCmd.Flags().String("poolID", "", "")

	_ = createPoolClientCmd.MarkFlagRequired("name")
	_ = createPoolClientCmd.MarkFlagRequired("poolID")

	rootCmd.AddCommand(createPoolClientCmd)
}
