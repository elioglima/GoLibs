package logs

import (
	libs "GoLibs"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

func print(CorSel color.Attribute, a ...interface{}) error {

	td, tm, ty, th, tn, ts := libs.TimeToDecStr(time.Now())
	b := fmt.Sprintf("%v/%v/%v %v:%v:%v", td, tm, ty, th, tn, ts)

	color.Set(color.FgMagenta)
	fmt.Printf("%s", b)

	color.Set(CorSel)

	valor := fmt.Sprintf("%v", a)
	valor = strings.Replace(valor, "[", "", -1)
	valor = strings.Replace(valor, "]", "", -1)

	if len(strings.TrimSpace(valor)) > 0 {
		fmt.Printf(" %v ", valor)
	}

	if DebugOrigem {
		_, file, line, _ := runtime.Caller(2)

		fpcs := make([]uintptr, 1)
		// Skip 2 levels to get the caller
		nfpcs := runtime.Callers(3, fpcs)
		if nfpcs == 0 {
			fmt.Println("MSG: NO CALLER")
		}

		caller := runtime.FuncForPC(fpcs[0] - 1)
		if caller == nil {
			fmt.Println("MSG CALLER WAS NIL")
		}

		ANF := strings.Split(caller.Name(), ".")
		color.Set(color.FgWhite)
		fmt.Printf("\nFunc: %v(), Ln: %v, File: %v\n", ANF[len(ANF)-1], line, file)
	} else {
		fmt.Printf("\n")

	}

	return nil
}
