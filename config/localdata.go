package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const clientsJSON = "clients.json"

type Client struct {
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

func ParsePoolIDByName(poolName string, storagePath string) (string, error) {
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

		var unmarshalled struct{ Options interface{} }
		err = json.Unmarshal(byteArr, &unmarshalled)
		if err != nil {
			return "", err
		}

		if unmarshalled.Options.(map[string]interface{})["Name"] == poolName {
			return fmt.Sprintf("%v", unmarshalled.Options.(map[string]interface{})["Id"]), nil
		}
	}

	return "", fmt.Errorf("unable to parse pool ID by name %s", poolName)
}

func ParseClientByID(clientID string, storagePath string) (*Client, error) {
	storagePathClientsJSON := storagePath + clientsJSON

	byteArr, err := os.ReadFile(storagePathClientsJSON)
	if err != nil {
		return nil, err
	}

	var unmarshalled struct{ Clients interface{} }
	err = json.Unmarshal(byteArr, &unmarshalled)
	if err != nil {
		return nil, err
	}

	for _, unmarshalledValue := range unmarshalled.Clients.(map[string]interface{}) {
		clientJSON, marshallErr := json.Marshal(unmarshalledValue)
		if marshallErr != nil {
			continue
		}

		client := &Client{}
		err = json.Unmarshal(clientJSON, client)
		if err != nil {
			continue
		}

		if client.ClientId == clientID {
			return client, nil
		}
	}

	return nil, fmt.Errorf("unable to parse client by ID %s", clientID)
}

func ParseConfirmationCode(username string, storagePath string) (string, error) {
	users, err := parseUsersJSON(storagePath)
	if err != nil {
		return "", err
	}

	for _, userParsed := range users {
		if userParsed.Username == username {
			return userParsed.ConfirmationCode, nil
		}
	}

	return "", fmt.Errorf("unable to parse confirmation code by username %s", username)
}

func ParseUserStatus(username string, storagePath string) (string, error) {
	users, err := parseUsersJSON(storagePath)
	if err != nil {
		return "", err
	}

	for _, userParsed := range users {
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

	var unmarshalled struct{ Users interface{} }
	err = json.Unmarshal(byteArr, &unmarshalled)
	if err != nil {
		return nil, err
	}

	var users []*user
	usersUnmarshalledMap := unmarshalled.Users.(map[string]interface{})

	for _, unmarshalledValue := range usersUnmarshalledMap {
		userJSON, marshallErr := json.Marshal(unmarshalledValue)
		if marshallErr != nil {
			continue
		}

		userUnmarshalled := &user{}
		err = json.Unmarshal(userJSON, userUnmarshalled)
		if err != nil {
			continue
		}

		users = append(users, userUnmarshalled)
	}

	if len(users) == 0 && len(usersUnmarshalledMap) > 0 {
		return nil, fmt.Errorf("unable to parse users from %s", storagePath)
	}

	return users, nil
}
