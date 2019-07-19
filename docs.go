package GoLibs

import (
	"errors"
	"strings"

	brdoc "gopkg.in/Nhanderu/brdoc.v2"
)

func IsCPF(value string) error {
	SoNumeroStr, err := SoNumeros(value)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(SoNumeroStr)) > 11 {
		SoNumeroStr = SoNumeroStr[len(SoNumeroStr)-11 : len(SoNumeroStr)]

	} else if len(strings.TrimSpace(SoNumeroStr)) < 11 {
		// exemplo de cpf - 216.399.218-77 21639921877
		for index := len(strings.TrimSpace(SoNumeroStr)); index < 11; index++ {
			SoNumeroStr = "0" + SoNumeroStr
		}
	}

	if ok := brdoc.IsCPF(SoNumeroStr); !ok {
		return errors.New("CPF inválido")
	}
	return nil
}

func IsCNPJ(value string) error {
	SoNumeroStr, err := SoNumeros(value)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(SoNumeroStr)) > 14 {
		SoNumeroStr = SoNumeroStr[len(SoNumeroStr)-14 : len(SoNumeroStr)]

	} else if len(strings.TrimSpace(SoNumeroStr)) < 14 {
		// exemplo de cnpj - 00463843/0001-24
		for index := len(strings.TrimSpace(SoNumeroStr)); index < 14; index++ {
			SoNumeroStr = "0" + SoNumeroStr
		}
	}

	if ok := brdoc.IsCNPJ(SoNumeroStr); !ok {
		return errors.New("CNPJ inválido")
	}

	return nil
}

func ImprimeCPF(value string) (string, error) {

	SoNumeroStr, err := SoNumeros(value)
	if err != nil {
		return "", err
	}

	if len(strings.TrimSpace(SoNumeroStr)) > 11 {
		SoNumeroStr = SoNumeroStr[len(SoNumeroStr)-11 : len(SoNumeroStr)]

	} else if len(strings.TrimSpace(SoNumeroStr)) < 11 {
		// exemplo de cpf - 216.399.218-77 21639921877
		for index := len(strings.TrimSpace(SoNumeroStr)); index < 11; index++ {
			SoNumeroStr = "0" + SoNumeroStr
		}
	}

	if err := IsCPF(SoNumeroStr); err != nil {
		return "", err
	}

	CPFFormatado := SoNumeroStr[0:3]
	CPFFormatado += "." + SoNumeroStr[3:6]
	CPFFormatado += "." + SoNumeroStr[6:9]
	CPFFormatado += "-" + SoNumeroStr[9:11]
	return CPFFormatado, nil
}

func ImprimeCNPJ(value string) (string, error) {
	SoNumeroStr, err := SoNumeros(value)
	if err != nil {
		return "", err
	}

	if len(strings.TrimSpace(SoNumeroStr)) > 14 {
		SoNumeroStr = SoNumeroStr[len(SoNumeroStr)-14 : len(SoNumeroStr)]

	} else if len(strings.TrimSpace(SoNumeroStr)) < 14 {
		// exemplo de cnpj - 00463843/0001-24
		for index := len(strings.TrimSpace(SoNumeroStr)); index < 14; index++ {
			SoNumeroStr = "0" + SoNumeroStr
		}
	}

	if err := IsCNPJ(SoNumeroStr); err != nil {
		err := errors.New("Formatação:" + err.Error())
		return err.Error(), err
	}

	CPFFormatado := SoNumeroStr[0:8]
	CPFFormatado += "/" + SoNumeroStr[8:12]
	CPFFormatado += "-" + SoNumeroStr[12:14]
	return CPFFormatado, nil
}
