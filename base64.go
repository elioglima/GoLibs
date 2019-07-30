package GoLibs

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

func Base64Encode(dados string) (string, error) {
	if len(strings.TrimSpace(dados)) == 0 {
		return "", nil
	}

	StrEncodeB64 := base64.StdEncoding.EncodeToString([]byte(dados))

	if len(strings.TrimSpace(StrEncodeB64)) == 0 {
		return "", errors.New("Nenhum registro encontrado.")
	}

	return StrEncodeB64, nil
}

func Base64Decode(dados string) (string, error) {
	if len(strings.TrimSpace(dados)) == 0 {
		return "", nil
	}

	StrDecodeB64, err := base64.StdEncoding.DecodeString(dados)
	if err != nil {
		return "", err
	}

	resultado := fmt.Sprintf("%s", StrDecodeB64)
	if len(strings.TrimSpace(resultado)) == 0 {
		return "", errors.New("Nenhum registro encontrado.")
	}

	return resultado, nil
}
