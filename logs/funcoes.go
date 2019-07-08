package logs

import (
	"github.com/fatih/color"
)

func Erro(a ...interface{}) error {
	return print(color.FgRed, a...)
}

func Atencao(a ...interface{}) error {
	return print(color.FgYellow, a...)
}

func Sucesso(a ...interface{}) error {
	return print(color.FgGreen, a...)
}

func Rosa(a ...interface{}) error {
	return print(color.FgHiMagenta, a...)
}

func Cyan(a ...interface{}) error {
	return print(color.FgHiCyan, a...)
}

func Cinza(a ...interface{}) error {
	return print(color.FgWhite, a...)
}

func Branco(a ...interface{}) error {
	return print(color.FgWhite, a)
}

func BrancoHi(a ...interface{}) error {
	return print(color.FgHiWhite, a...)
}

func AtencaoHI(a ...interface{}) error {
	return print(color.FgHiYellow, a...)
}
