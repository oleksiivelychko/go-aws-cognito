package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var signUpCmd = &cobra.Command{
	Use:   "sign-up",
	Short: "Registers a new user in the user pool.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := cognito.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		err = srv.SignUp(
			cmd.Flag("username").Value.String(),
			cmd.Flag("password").Value.String(),
			cmd.Flag("clientID").Value.String(),
		)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("✅ Operation has been successful!")
		}
	},
}

func init() {
	signUpCmd.Flags().StringP("username", "u", "", "")
	signUpCmd.Flags().StringP("password", "p", "", "")
	signUpCmd.Flags().String("clientID", "", "")

	_ = signUpCmd.MarkFlagRequired("username")
	_ = signUpCmd.MarkFlagRequired("password")
	_ = signUpCmd.MarkFlagRequired("clientID")

	rootCmd.AddCommand(signUpCmd)
}
