package cognito

import (
	"github.com/oleksiivelychko/go-aws-cognito/config"
	"log"
	"strings"
	"testing"
)

const (
	username  = "test@test.test"
	password  = "secret"
	localData = "./." + config.LocalData
)

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
	poolName := "test pool"

	poolID, err = srv.CreatePool(poolName)
	if err != nil {
		t.Fatal(err)
	}

	parsedPoolID, err := config.ParsePoolIDByName(poolName, localData)
	if err != nil {
		t.Errorf("unable to parse pool ID: %s", err)
	}

	if poolID != parsedPoolID {
		t.Errorf("poolID %s doesn't match with the parsed one %s", poolID, parsedPoolID)
	}
}

func TestCognito_DescribePool(t *testing.T) {
	_, err := srv.DescribePool(poolID)
	if err != nil {
		t.Error(err)
	}
}

func TestCognito_CreatePoolClient(t *testing.T) {
	var err error
	clientName := "test client"

	poolClientID, err = srv.CreatePoolClient(clientName, poolID)
	if err != nil {
		t.Fatal(err)
	}

	client, err := config.ParseClientByID(poolClientID, localData)
	if err != nil {
		t.Errorf("unable to parse pool client ID: %s", err)
	}

	if poolClientID != client.ClientId {
		t.Errorf("poolClientID %s doesn't match with the parsed one %s", poolClientID, client.ClientId)
	}
}

func TestCognito_SignUp(t *testing.T) {
	err := srv.SignUp(username, password, poolClientID)
	if err != nil {
		t.Error(err)
	}
}

func TestCognito_SameSignUp(t *testing.T) {
	err := srv.SignUp(username, password, poolClientID)
	if err == nil {
		t.Fatal("account must exist")
	}

	if !strings.HasPrefix(err.Error(), "UsernameExistsException") {
		t.Errorf("expected UsernameExistsException, got %s", err.Error())
	}
}

func TestCognito_ConfirmSignUp(t *testing.T) {
	pathPoolID := localData + poolID + ".json"

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
		t.Error("unable to confirm user status")
	}
}

func TestCognito_SignIn(t *testing.T) {
	var err error
	accessToken, err = srv.SignIn(username, password, poolClientID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCognito_DeleteUser(t *testing.T) {
	err := srv.DeleteUser(accessToken)
	if err != nil {
		t.Error(err)
	}
}

func TestCognito_DeletePoolClient(t *testing.T) {
	err := srv.DeletePoolClient(poolID, poolClientID)
	if err != nil {
		t.Error(err)
	}
}

func TestCognito_DeletePool(t *testing.T) {
	err := srv.DeletePool(poolID)
	if err != nil {
		t.Error(err)
	}
}
