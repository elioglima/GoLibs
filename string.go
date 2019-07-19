package GoLibs

import (
	"errors"
	"regexp"
	"unicode"
)

func SoNumeros(doc string) (string, error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(doc, ""), nil
}

func Asp(str string) string {
	return "'" + str + "'"
}

func StrJustNumber(s string) error {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return errors.New("A string " + s + " contem números.")
		}
	}
	return nil
}

func StrJustCharacter(s string) error {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return errors.New("A string " + s + " contém caracteres.")
		}
	}
	return nil
}
