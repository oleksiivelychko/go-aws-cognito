package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var confirmSignUpCmd = &cobra.Command{
	Use:   "confirm-sign-up",
	Short: "Confirms registration of a new user.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := cognito.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		err = srv.ConfirmSignUp(
			cmd.Flag("username").Value.String(),
			cmd.Flag("code").Value.String(),
			cmd.Flag("clientID").Value.String(),
		)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(SuccessfulMessage)
		}
	},
}

func init() {
	confirmSignUpCmd.Flags().StringP("username", "u", "", "")
	confirmSignUpCmd.Flags().String("code", "", "confirmation code")
	confirmSignUpCmd.Flags().String("clientID", "", "")

	_ = confirmSignUpCmd.MarkFlagRequired("username")
	_ = confirmSignUpCmd.MarkFlagRequired("code")
	_ = confirmSignUpCmd.MarkFlagRequired("clientID")

	rootCmd.AddCommand(confirmSignUpCmd)
}
