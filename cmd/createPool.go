package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var createPoolCmd = &cobra.Command{
	Use:   "create-pool",
	Short: "Creates a user pool.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := cognito.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		poolID, err := srv.CreatePool(cmd.Flag("name").Value.String())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("✅ User pool ID: %s\n", poolID)
		}
	},
}

func init() {
	createPoolCmd.Flags().String("name", "", "")

	_ = createPoolCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createPoolCmd)
}
