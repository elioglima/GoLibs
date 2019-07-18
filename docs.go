package GoLibs

import (
	"errors"

	brdoc "gopkg.in/Nhanderu/brdoc.v2"
)

func IsCPF(value string) error {
	processedString, err := SoNumeros(value)
	if err != nil {
		return err
	}

	if ok := brdoc.IsCPF(processedString); !ok {
		return errors.New("CPF inválido")
	}
	return nil
}

func IsCNPJ(value string) error {
	processedString, err := SoNumeros(value)
	if err != nil {
		return err
	}

	if ok := brdoc.IsCNPJ(processedString); !ok {
		return errors.New("CCNPJ inválido")
	}

	return nil
}

func ImprimeCPF(value string) (string, error) {

	processedString, err := SoNumeros(value)
	if err != nil {
		return "", err
	}

	if err := IsCPF(processedString); err != nil {
		return "", err
	}

	CPFFormatado := processedString[0:3]
	CPFFormatado += "." + processedString[3:6]
	CPFFormatado += "." + processedString[6:9]
	CPFFormatado += "-" + processedString[9:11]
	return CPFFormatado, nil
}

func ImprimeCNPJ(value string) (string, error) {
	processedString, err := SoNumeros(value)
	if err != nil {
		return "", err
	}

	if err := IsCPF(processedString); err != nil {
		return "", err
	}

	CPFFormatado := processedString[0:3]
	CPFFormatado += "." + processedString[3:6]
	CPFFormatado += "." + processedString[6:9]
	CPFFormatado += "-" + processedString[9:11]
	return CPFFormatado, nil
}
