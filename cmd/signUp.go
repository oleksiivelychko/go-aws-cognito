package cmd

import (
	"fmt"
	"github.com/oleksiivelychko/go-aws-cognito/config"
	"github.com/oleksiivelychko/go-aws-cognito/service"
	"github.com/spf13/cobra"
)

var signUpCmd = &cobra.Command{
	Use:   "sign-up",
	Short: "Registers a new user in the user pool.",
	Run: func(cmd *cobra.Command, args []string) {
		cognito, err := service.New(cfgAWS)
		if err != nil {
			cobra.CheckErr(err)
		}

		username := cmd.Flag("username").Value.String()
		clientID := cmd.Flag("clientID").Value.String()

		err = cognito.SignUp(username, cmd.Flag("password").Value.String(), clientID)
		if err != nil {
			cobra.CheckErr(err)
		}

		if cfgAWS.Endpoint == config.LocalEndpoint {
			code := parseLocalConfirmationCode(clientID, username)
			fmt.Printf("%s Confirmation code is %s\n", SuccessfulMessage, code)
		} else {
			fmt.Println(SuccessfulMessage)
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

func parseLocalConfirmationCode(clientID, username string) string {
	client, err := config.ParseClientByID(clientID, config.LocalData)
	if err != nil {
		cobra.CheckErr(err)
	}

	code, _ := config.ParseConfirmationCode(username, config.LocalData+client.UserPoolId+".json")
	return code
}
