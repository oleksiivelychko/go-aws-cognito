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
	var err error
	poolName := "My pool"

	poolID, err = srv.CreatePool(poolName)
	if err != nil {
		t.Fatal(err)
	}

	parsedPoolID, err := config.ParsePoolID(poolName, "./../data/db")
	if err != nil {
		t.Fatal(err)
	}

	if poolID != parsedPoolID {
		t.Fatalf("poolID %s doesn't match with the parsed one %s", poolID, parsedPoolID)
	}
}

func TestCognito_DescribePool(t *testing.T) {
	_, err := srv.DescribePool(poolID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCognito_CreatePoolClient(t *testing.T) {
	clientName := "My service"

	_, err := srv.CreatePoolClient(clientName, poolID)
	if err != nil {
		t.Fatal(err)
	}

	poolClientID, err = config.ParseClientID(clientName, "./../data/db")
	if err != nil {
		t.Fatalf("unable to detect client ID: %s", err)
	}

	if poolClientID == "" {
		t.Fatal("got empty client ID")
	}
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
