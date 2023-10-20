package service

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
	cognito      IService
)

func init() {
	cfg, err := config.ReadYAML("./../config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	cognito, err = New(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestCreatePool(t *testing.T) {
	var err error
	poolName := "test pool"

	poolID, err = cognito.CreatePool(poolName)
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

func TestDescribePool(t *testing.T) {
	_, err := cognito.DescribePool(poolID)
	if err != nil {
		t.Error(err)
	}
}

func TestCreatePoolClient(t *testing.T) {
	var err error
	clientName := "test client"

	poolClientID, err = cognito.CreatePoolClient(clientName, poolID)
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

func TestSignUp(t *testing.T) {
	err := cognito.SignUp(username, password, poolClientID)
	if err != nil {
		t.Error(err)
	}
}

func TestSameSignUp(t *testing.T) {
	err := cognito.SignUp(username, password, poolClientID)
	if err == nil {
		t.Fatal("account must exist")
	}

	if !strings.HasPrefix(err.Error(), "UsernameExistsException") {
		t.Errorf("expected UsernameExistsException, got %s", err.Error())
	}
}

func TestConfirmSignUp(t *testing.T) {
	pathPoolID := localData + poolID + ".json"

	signupConfirmationCode, err := config.ParseConfirmationCode(username, pathPoolID)
	if signupConfirmationCode == "" {
		t.Fatal(err)
	}

	err = cognito.ConfirmSignUp(username, signupConfirmationCode, poolClientID)
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

func TestSignIn(t *testing.T) {
	var err error
	accessToken, err = cognito.SignIn(username, password, poolClientID)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := cognito.DeleteUser(accessToken)
	if err != nil {
		t.Error(err)
	}
}

func TestDeletePoolClient(t *testing.T) {
	err := cognito.DeletePoolClient(poolID, poolClientID)
	if err != nil {
		t.Error(err)
	}
}

func TestDeletePool(t *testing.T) {
	err := cognito.DeletePool(poolID)
	if err != nil {
		t.Error(err)
	}
}
