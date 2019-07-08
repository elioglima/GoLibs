package libs

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	Logger *Logs
)

type Logs struct {
	*color.Color
}

// overload
func NewLogs() *Logs {
	c := &Logs{color.New(color.FgWhite)}
	return c
}

func (c *Logs) Erro(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgRed, a...)
}

func (c *Logs) Atencao(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgYellow, a...)
}

func (c *Logs) Sucesso(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgGreen, a...)
}

func (c *Logs) Rosa(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgHiMagenta, a...)
}

func (c *Logs) Cyan(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgHiCyan, a...)
}

func (c *Logs) Cinza(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgWhite, a...)
}

func (c *Logs) Branco(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgWhite, a)
}

func (c *Logs) BrancoHi(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgHiWhite, a...)
}

func (c *Logs) AtencaoHI(a ...interface{}) (n int, err error) {
	return c.Formats(color.FgHiYellow, a...)
}

func (c *Logs) Formats(cor color.Attribute, a ...interface{}) (n int, err error) {

	td, tm, ty, th, tn, ts := TimeToDecStr(time.Now())
	b := fmt.Sprintf("%v/%v/%v %v:%v:%v", td, tm, ty, th, tn, ts)

	c = &Logs{color.New(color.FgMagenta)}
	c.Set()
	fmt.Fprintln(color.Output, "")
	fmt.Fprintf(color.Output, b)

	c = &Logs{color.New(cor)}
	c.Set()
	valor := fmt.Sprintf("%v", a)
	valor = strings.Replace(valor, "[", "", -1)
	valor = strings.Replace(valor, "]", "", -1)

	if len(strings.TrimSpace(valor)) > 0 {
		fmt.Fprintf(color.Output, " %v ", valor)
	}

	// fmt.Fprintln(color.Output, "")

	// _, file, line, _ := runtime.Caller(2)

	// fpcs := make([]uintptr, 1)
	// // Skip 2 levels to get the caller
	// nfpcs := runtime.Callers(3, fpcs)
	// if nfpcs == 0 {
	// 	fmt.Println("MSG: NO CALLER")
	// }

	// caller := runtime.FuncForPC(fpcs[0] - 1)
	// if caller == nil {
	// 	fmt.Println("MSG CALLER WAS NIL")
	// }

	// ANF := strings.Split(caller.Name(), ".")

	// c = &Logs{color.New(color.FgWhite)}
	// c.Set()
	// fmt.Fprintf(color.Output, "Func: %v(), Ln: %v, File: %v", ANF[len(ANF)-1], line, file)
	fmt.Fprintln(color.Output, "")

	return 0, nil //fmt.Fprintln(color.Output, va...)
}
