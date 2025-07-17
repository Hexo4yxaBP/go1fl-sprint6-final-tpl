package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseConvert(sourceString string) string {

	trimmedString := strings.Trim(sourceString, ".- \n")

	if len(trimmedString) == 0 {
		return morse.ToText(sourceString)
	}

	return morse.ToMorse(sourceString)

}
