package tempfunctions

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Convert() (float64, string, error) {
	userInputString, err := Input()
	if err != nil {
		return 0.0, "", err
	}
	userInputString = strings.TrimSpace(strings.ToLower(userInputString))
	matched, err := regexp.MatchString(`^\d*\.?\d*\s[cfk](?i)$`, userInputString)
	if err != nil {
		return 0.0, "", err
	}
	if !matched {
		return 0.0, "", errors.New("Invalid Format!!")
	}

	userInputSplit := strings.Split(userInputString, " ")
	userNumber := userInputSplit[0]
	unit := userInputSplit[1]

	temperature, err := strconv.ParseFloat(userNumber, 64)
	if err != nil {
		return 0.0, "", err
	}
	return temperature, unit, err
}
