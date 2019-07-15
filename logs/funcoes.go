package logs

import (
	"github.com/fatih/color"
)

var (
	DebugOrigem  bool = false
	DebugErro    bool = false
	DebugSucesso bool = false
)

func Erro(a ...interface{}) error {
	if !DebugErro {
		return nil
	}
	return print(color.FgRed, a...)
}

func Atencao(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgYellow, a...)
}

func Sucesso(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgGreen, a...)
}

func Rosa(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgHiMagenta, a...)
}

func Cyan(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgHiCyan, a...)
}

func Cinza(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgWhite, a...)
}

func Branco(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgWhite, a)
}

func BrancoHi(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgHiWhite, a...)
}

func AtencaoHI(a ...interface{}) error {
	if !DebugSucesso {
		return nil
	}
	return print(color.FgHiYellow, a...)
}
