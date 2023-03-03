package cognito

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	awsCognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type IClientCognito interface {
	SignUp(email, password string) (string, error)
	ConfirmSignUp(email, confirmationCode string) error
	SignIn(email, password string) (*SignInResult, error)
}

type ClientCognito struct {
	cognito     *awsCognito.CognitoIdentityProvider
	appClientId string
	authFlow    string
}

type SignInResult struct {
	IdToken      string
	AccessToken  string
	RefreshToken string
	ExpiresIn    uint
	TokenType    string
}

func (result *SignInResult) String() string {
	return fmt.Sprintf("TokenType: %s\nExpiresIn: %d\nIdToken: %s\nAccessToken: %s\nRefreshToken: %s\n",
		result.TokenType, result.ExpiresIn, result.IdToken, result.AccessToken, result.RefreshToken)
}

func NewClientCognito(region, appClientId string) (IClientCognito, error) {
	session, err := awsSession.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	return &ClientCognito{
		cognito:     awsCognito.New(session),
		appClientId: appClientId,
		authFlow:    "USER_PASSWORD_AUTH",
	}, nil
}

func (client *ClientCognito) SignUp(email, password string) (string, error) {
	input := &awsCognito.SignUpInput{
		Username: aws.String(email),
		Password: aws.String(password),
		ClientId: aws.String(client.appClientId),
		UserAttributes: []*awsCognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	output, err := client.cognito.SignUp(input)
	return output.String(), err
}

func (client *ClientCognito) ConfirmSignUp(email, confirmationCode string) error {
	input := &awsCognito.ConfirmSignUpInput{
		Username:         aws.String(email),
		ConfirmationCode: aws.String(confirmationCode),
		ClientId:         aws.String(client.appClientId),
	}

	_, err := client.cognito.ConfirmSignUp(input)
	return err
}

func (client *ClientCognito) SignIn(email, password string) (*SignInResult, error) {
	input := &awsCognito.InitiateAuthInput{
		AuthFlow: aws.String(client.authFlow),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		}),
		ClientId: aws.String(client.appClientId),
	}

	output, err := client.cognito.InitiateAuth(input)
	if err != nil {
		return nil, err
	}

	return &SignInResult{
		IdToken:      *output.AuthenticationResult.IdToken,
		AccessToken:  *output.AuthenticationResult.AccessToken,
		RefreshToken: *output.AuthenticationResult.RefreshToken,
		ExpiresIn:    uint(*output.AuthenticationResult.ExpiresIn),
		TokenType:    *output.AuthenticationResult.TokenType,
	}, nil
}
