package GoLibs

import (
	"fmt"
	"strings"
)

func FormataMoeda(valor_in float64, especie_in string) string {
	valor := fmt.Sprintf("%.2f", valor_in)
	if len(strings.TrimSpace(especie_in)) > 0 {
		valor = especie_in + " " + valor
	}
	return valor
}
