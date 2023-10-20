package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
	"os"
)

var signInCmd = &cobra.Command{
	Use:   "sign-in",
	Short: "Initiates sign-in for user in the Cognito user directory.",
	Run: func(cmd *cobra.Command, args []string) {
		cognito, err := service.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		accessToken, err := cognito.SignIn(
			cmd.Flag("username").Value.String(),
			cmd.Flag("password").Value.String(),
			cmd.Flag("clientID").Value.String(),
		)

		if err != nil {
			cobra.CheckErr(err)
		}

		var accessTokenTxt = "access-token.txt"

		file, err := os.Create(accessTokenTxt)
		if err != nil {
			cobra.CheckErr(err)
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		_, err = file.Write([]byte(accessToken))
		if err != nil {
			cobra.CheckErr(err)
		}

		err = file.Sync()
		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("✅ Access token has been written to %s\n", accessTokenTxt)
	},
}

func init() {
	signInCmd.Flags().StringP("username", "u", "", "")
	signInCmd.Flags().StringP("password", "p", "", "")
	signInCmd.Flags().String("clientID", "", "")

	_ = signInCmd.MarkFlagRequired("username")
	_ = signInCmd.MarkFlagRequired("password")
	_ = signInCmd.MarkFlagRequired("clientID")

	rootCmd.AddCommand(signInCmd)
}
