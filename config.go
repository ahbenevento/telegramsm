package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

//  //  //

type botList map[string]string

type userList map[int64]string

//  //  //

type appConfig struct {
	Bots  botList  `json:"bots"`
	Users userList `json:"users,omitempty"`
}

//  //  //

func loadConfig(filename string) (*appConfig, error) {
	if _, err := os.Stat(filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("%s (%s)", os.ErrNotExist, filename)
		}

		return nil, err
	}

	buf, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	result := appConfig{}
	err = json.Unmarshal(buf, &result)

	if err != nil {
		return nil, err
	} else if result.Users == nil {
		result.Users = make(userList)
	}

	return &result, nil
}

func saveConfig(cfg *appConfig, filename string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	buf, err := json.MarshalIndent(cfg, "", "    ")

	if err != nil {
		return err
	}

	_, err = f.Write(buf)

	return err
}
