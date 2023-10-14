package cognito

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/oleksiivelychko/go-aws-cognito/config"
)

type IService interface {
	SignIn(username, password, poolClientID string) (*cognitoidentityprovider.AuthenticationResultType, error)
	SignUp(username, password, poolClientID string) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUp(username, confirmationCode, poolClientID string) error
	DeleteUser(accessToken string) error
	CreatePool(name string) (string, error)
	CreatePoolClient(name, poolID string) (*cognitoidentityprovider.CreateUserPoolClientOutput, error)
	DescribePool(poolID string) (string, error)
	DeletePool(poolID string) error
	DeletePoolClient(poolID, clientID string) error
}

type service struct {
	client *cognitoidentityprovider.CognitoIdentityProvider
}

func New(config *config.AWS) (IService, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(config.Region),
		Endpoint:    aws.String(config.Endpoint),
		Credentials: credentials.NewStaticCredentials(config.AwsAccessKeyId, config.AwsSecretAccessKey, ""),
	}

	awsSession, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	return &service{client: cognitoidentityprovider.New(awsSession)}, nil
}

func (service *service) SignUp(username, password, poolClientID string) (*cognitoidentityprovider.SignUpOutput, error) {
	return service.client.SignUp(&cognitoidentityprovider.SignUpInput{
		Username: aws.String(username),
		Password: aws.String(password),
		ClientId: aws.String(poolClientID),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("username"),
				Value: aws.String(username),
			},
		},
	})
}

func (service *service) ConfirmSignUp(username, confirmationCode, poolClientID string) error {
	_, err := service.client.ConfirmSignUp(&cognitoidentityprovider.ConfirmSignUpInput{
		Username:         aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
		ClientId:         aws.String(poolClientID),
	})

	return err
}

func (service *service) SignIn(username, password, poolClientID string) (*cognitoidentityprovider.AuthenticationResultType, error) {
	output, err := service.client.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		}),
		ClientId: aws.String(poolClientID),
	})

	if err != nil {
		return nil, err
	}

	return output.AuthenticationResult, nil
}

func (service *service) DeleteUser(accessToken string) error {
	_, err := service.client.DeleteUser(&cognitoidentityprovider.DeleteUserInput{
		AccessToken: aws.String(accessToken),
	})

	return err
}

func (service *service) CreatePool(name string) (string, error) {
	output, err := service.client.CreateUserPool(&cognitoidentityprovider.CreateUserPoolInput{
		PoolName: aws.String(name),
		Policies: &cognitoidentityprovider.UserPoolPolicyType{
			PasswordPolicy: &cognitoidentityprovider.PasswordPolicyType{
				MinimumLength:                 aws.Int64(6),
				RequireLowercase:              aws.Bool(false),
				RequireNumbers:                aws.Bool(false),
				RequireSymbols:                aws.Bool(false),
				RequireUppercase:              aws.Bool(false),
				TemporaryPasswordValidityDays: aws.Int64(7),
			},
		},
	})

	if err != nil {
		return "", err
	}

	return *output.UserPool.Id, nil
}

func (service *service) CreatePoolClient(name, poolID string) (*cognitoidentityprovider.CreateUserPoolClientOutput, error) {
	return service.client.CreateUserPoolClient(&cognitoidentityprovider.CreateUserPoolClientInput{
		ClientName: aws.String(name),
		UserPoolId: aws.String(poolID),
		TokenValidityUnits: &cognitoidentityprovider.TokenValidityUnitsType{
			AccessToken:  aws.String(cognitoidentityprovider.TimeUnitsTypeHours),
			IdToken:      aws.String(cognitoidentityprovider.TimeUnitsTypeMinutes),
			RefreshToken: aws.String(cognitoidentityprovider.TimeUnitsTypeDays),
		},
		AccessTokenValidity:  aws.Int64(24),
		IdTokenValidity:      aws.Int64(60),
		RefreshTokenValidity: aws.Int64(7),
	})
}

func (service *service) DeletePool(poolID string) error {
	_, err := service.client.DeleteUserPool(&cognitoidentityprovider.DeleteUserPoolInput{
		UserPoolId: aws.String(poolID),
	})

	return err
}

func (service *service) DeletePoolClient(poolID, clientID string) error {
	_, err := service.client.DeleteUserPoolClient(&cognitoidentityprovider.DeleteUserPoolClientInput{
		ClientId:   aws.String(clientID),
		UserPoolId: aws.String(poolID),
	})

	return err
}

func (service *service) DescribePool(poolID string) (string, error) {
	output, err := service.client.DescribeUserPool(&cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: aws.String(poolID),
	})

	return output.String(), err
}
