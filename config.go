package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

//  //  //

type botList map[string]string

//  //  //

type appConfig struct {
	Bots botList `json:"bots"`
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
	}

	return &result, nil
}
