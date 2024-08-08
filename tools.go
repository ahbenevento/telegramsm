package main

import (
	"fmt"
	"strconv"
	"strings"
)

//  //  //

func getBotTokenByName(bots botList, botName string) (string, error) {
	token, ok := bots[botName]

	if ok {
		return token, nil
	}

	return "", fmt.Errorf("bot no encontrado: \"%s\"", botName)
}

func getUserID(users userList, username string) (int64, error) {
	funcIsDigit := func(c rune) bool {
		return c >= '0' && c <= '9'
	}

	if strings.ContainsFunc(username, funcIsDigit) {
		return strconv.ParseInt(username, 10, 64)
	}

	for id, name := range users {
		if name == username {
			return id, nil
		}
	}

	return 0, fmt.Errorf("usuario no encontrado: \"%s\"", username)
}
