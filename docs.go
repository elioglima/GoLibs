package libs

import (
	"errors"
	"regexp"

	brdoc "gopkg.in/Nhanderu/brdoc.v2"
)

func IsCPF(doc string) error {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return err
	}
	processedString := reg.ReplaceAllString(doc, "")
	if ok := brdoc.IsCPF(processedString); !ok {
		return errors.New("CPF inválido")
	}
	return nil
}

func IsCNPJ(doc string) error {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return err
	}

	processedString := reg.ReplaceAllString(doc, "")

	if ok := brdoc.IsCNPJ(processedString); !ok {
		return errors.New("CCNPJ inválido")
	}

	return nil
}
