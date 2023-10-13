package cognito

import (
	"github.com/oleksiivelychko/go-aws-cognito/config"
	"log"
	"strings"
	"testing"
)

const username = "test@test.test"
const password = "secret"

var (
	poolID       string
	poolClientID string
	accessToken  string
	srv          IService
)

func init() {
	cfg, err := config.ReadYAML("./../config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	srv, err = New(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestCognito_CreatePool(t *testing.T) {
	output, err := srv.CreatePool("My Pool")
	if err != nil {
		t.Fatal(err)
	}

	poolID = *output.UserPool.Id
}

func TestCognito_DescribePool(t *testing.T) {
	_, err := srv.DescribePool(poolID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCognito_CreatePoolClient(t *testing.T) {
	output, err := srv.CreatePoolClient("My service", poolID)
	if err != nil {
		t.Fatal(err)
	}

	poolClientID = *output.UserPoolClient.ClientId
}

func TestCognito_SignUp(t *testing.T) {
	_, err := srv.SignUp(username, password, poolClientID)
	if err != nil {
		t.Error(err)
	}
}

func TestCognito_SameSignUp(t *testing.T) {
	_, err := srv.SignUp(username, password, poolClientID)
	if err == nil {
		t.Fatal("account must exist")
	}

	if !strings.HasPrefix(err.Error(), "UsernameExistsException") {
		t.Fatalf("should have been an UsernameExistsException, got %s", err.Error())
	}
}

func TestCognito_ConfirmSignUp(t *testing.T) {
	pathPoolID := "./../data/db/" + poolID + ".json"

	signupConfirmationCode, err := config.ParseConfirmationCode(username, pathPoolID)
	if signupConfirmationCode == "" {
		t.Fatal(err)
	}

	err = srv.ConfirmSignUp(username, signupConfirmationCode, poolClientID)
	if err != nil {
		t.Fatal(err)
	}

	userStatus, err := config.ParseUserStatus(username, pathPoolID)
	if err != nil {
		t.Fatal(err)
	}

	if userStatus != "CONFIRMED" {
		t.Fatal("unable to confirm user status")
	}
}

func TestCognito_SignIn(t *testing.T) {
	result, err := srv.SignIn(username, password, poolClientID)
	if err != nil {
		t.Fatal(err)
	}

	accessToken = *result.AccessToken
}

func TestCognito_DeleteUser(t *testing.T) {
	err := srv.DeleteUser(accessToken)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCognito_DeletePoolClient(t *testing.T) {
	err := srv.DeletePoolClient(poolID, poolClientID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCognito_DeletePool(t *testing.T) {
	err := srv.DeletePool(poolID)
	if err != nil {
		t.Fatal(err)
	}
}