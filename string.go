package libs

import (
	"regexp"
)

func SoNumeros(doc string) (string, error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(doc, ""), nil
}
