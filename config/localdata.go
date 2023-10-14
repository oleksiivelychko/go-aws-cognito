package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const clientsJSON = "clients.json"

type client struct {
	ClientId           string    `json:"ClientId"`
	ClientName         string    `json:"ClientName"`
	CreationDate       time.Time `json:"CreationDate"`
	LastModifiedDate   time.Time `json:"LastModifiedDate"`
	UserPoolId         string    `json:"UserPoolId"`
	TokenValidityUnits struct {
		AccessToken  string `json:"AccessToken"`
		IdToken      string `json:"TokenID"`
		RefreshToken string `json:"RefreshToken"`
	}
}

type user struct {
	Enabled              bool     `json:"Enabled"`
	Password             string   `json:"Password"`
	RefreshTokens        []string `json:"RefreshTokens"`
	UserCreateDate       string   `json:"UserCreateDate"`
	UserLastModifiedDate string   `json:"UserLastModifiedDate"`
	Username             string   `json:"Username"`
	UserStatus           string   `json:"UserStatus"`
	ConfirmationCode     string   `json:"ConfirmationCode"`
	Attributes           []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	}
}

func ParsePoolID(poolName string, storagePath string) (string, error) {
	files, readDirErr := os.ReadDir(storagePath)
	if readDirErr != nil {
		return "", readDirErr
	}

	for _, f := range files {
		if f.Name() == clientsJSON {
			continue
		}

		byteArr, err := os.ReadFile(storagePath + "/" + f.Name())
		if err != nil {
			return "", err
		}

		var optionsUnmarshalled struct{ Options interface{} }
		err = json.Unmarshal(byteArr, &optionsUnmarshalled)
		if err != nil {
			return "", err
		}

		if optionsUnmarshalled.Options.(map[string]interface{})["Name"] == poolName {
			return fmt.Sprintf("%v", optionsUnmarshalled.Options.(map[string]interface{})["Id"]), nil
		}
	}

	return "", fmt.Errorf("unable to parse pool ID by name %s", poolName)
}

func ParseClientID(clientName string, storagePath string) (string, error) {
	storagePathClientsJSON := storagePath + "/" + clientsJSON

	byteArr, err := os.ReadFile(storagePathClientsJSON)
	if err != nil {
		return "", err
	}

	var clientsUnmarshalled struct{ Clients interface{} }
	err = json.Unmarshal(byteArr, &clientsUnmarshalled)
	if err != nil {
		return "", err
	}

	var clientsMarshalled []*client
	clientsUnmarshalledMap := clientsUnmarshalled.Clients.(map[string]interface{})

	for _, clientUnmarshalledValue := range clientsUnmarshalled.Clients.(map[string]interface{}) {
		clientJSON, marshallErr := json.Marshal(clientUnmarshalledValue)
		if marshallErr != nil {
			continue
		}

		clientUnmarshalled := &client{}
		err = json.Unmarshal(clientJSON, clientUnmarshalled)
		if err != nil {
			continue
		}

		clientsMarshalled = append(clientsMarshalled, clientUnmarshalled)
	}

	if len(clientsMarshalled) == 0 && len(clientsUnmarshalledMap) > 0 {
		return "", fmt.Errorf("unable to parse clients from %s", storagePathClientsJSON)
	}

	if err != nil {
		return "", err
	}

	for _, clientMarshalled := range clientsMarshalled {
		if clientMarshalled.ClientName == clientName {
			return clientMarshalled.ClientId, nil
		}
	}

	return "", fmt.Errorf("unable to parse client ID by name %s", clientName)
}

func ParseConfirmationCode(username string, storagePath string) (string, error) {
	usersParsed, err := parseUsersJSON(storagePath)
	if err != nil {
		return "", err
	}

	for _, userParsed := range usersParsed {
		if userParsed.Username == username {
			return userParsed.ConfirmationCode, nil
		}
	}

	return "", fmt.Errorf("unable to parse confirmation code by username %s", username)
}

func ParseUserStatus(username string, storagePath string) (string, error) {
	usersParsed, err := parseUsersJSON(storagePath)
	if err != nil {
		return "", err
	}

	for _, userParsed := range usersParsed {
		if userParsed.Username == username {
			return userParsed.UserStatus, nil
		}
	}

	return "", fmt.Errorf("unable to parse user status by username %s", username)
}

func parseUsersJSON(storagePath string) ([]*user, error) {
	byteArr, err := os.ReadFile(storagePath)
	if err != nil {
		return nil, err
	}

	var usersUnmarshalled struct{ Users interface{} }
	err = json.Unmarshal(byteArr, &usersUnmarshalled)
	if err != nil {
		return nil, err
	}

	var usersParsed []*user
	usersUnmarshalledMap := usersUnmarshalled.Users.(map[string]interface{})

	for _, userUnmarshalledValue := range usersUnmarshalledMap {
		userJSON, marshallErr := json.Marshal(userUnmarshalledValue)
		if marshallErr != nil {
			continue
		}

		userUnmarshalled := &user{}
		err = json.Unmarshal(userJSON, userUnmarshalled)
		if err != nil {
			continue
		}

		usersParsed = append(usersParsed, userUnmarshalled)
	}

	if len(usersParsed) == 0 && len(usersUnmarshalledMap) > 0 {
		return nil, fmt.Errorf("unable to parse users from %s", storagePath)
	}

	return usersParsed, nil
}
