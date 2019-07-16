package GoLibs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func JustDateNow() time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	var d, m string

	tm := time.Now()
	d = strconv.Itoa(tm.Day())
	m = strconv.Itoa(int(tm.Month()))

	if tm.Day() < 10 {
		d = "0" + d
	}

	if int(tm.Month()) < 10 {
		m = "0" + m
	}

	str := fmt.Sprintf("%v", int(time.Now().Year())) + "-" + m + "-" + d
	str += "T00:00:00.000Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}
	}

	return t
}

func TimeToDate(t_in time.Time) time.Time {
	return JustDateParse(t_in)
}

func JustDateParse(tm time.Time) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	var d, m string

	d = strconv.Itoa(tm.Day())
	m = strconv.Itoa(int(tm.Month()))

	if tm.Day() < 10 {
		d = "0" + d
	}

	if int(tm.Month()) < 10 {
		m = "0" + m
	}

	str := fmt.Sprintf("%v", int(time.Now().Year())) + "-" + m + "-" + d
	str += "T00:00:00.000Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}
	}

	return t
}

func GetHora(d time.Time) string {
	h := strconv.Itoa(d.Hour())
	m := strconv.Itoa(d.Minute())
	s := strconv.Itoa(d.Second())

	if d.Hour() < 10 {
		h = "0" + h
	}

	if d.Minute() < 10 {
		m = "0" + m
	}

	if d.Second() < 10 {
		s = "0" + s
	}

	return fmt.Sprintf("%v:%v:%v", h, m, s)
}

func FormatDateTime(formato string, t time.Time) string {
	var valor_out string

	var sdia, smes, shora, smin, sseg string

	dia := t.Day()
	mes := int(t.Month())
	ano := t.Year()
	hora := t.Hour()
	min := t.Minute()
	seg := t.Second()

	sdia = strconv.Itoa(dia)
	smes = strconv.Itoa(mes)
	shora = strconv.Itoa(hora)
	smin = strconv.Itoa(min)
	sseg = strconv.Itoa(seg)

	if dia < 10 {
		sdia = "0" + strconv.Itoa(dia)
	}

	if mes < 10 {
		smes = "0" + strconv.Itoa(mes)
	}

	if hora < 10 {
		shora = "0" + strconv.Itoa(hora)
	}

	if min < 10 {
		smin = "0" + strconv.Itoa(min)
	}

	if seg < 10 {
		sseg = "0" + strconv.Itoa(seg)
	}

	switch strings.ToLower(strings.TrimSpace(formato)) {
	case "dd/mm/yyyy":
		valor_out = fmt.Sprintf("%d/%d/%d", dia, mes, ano)
	case "dd/mm/yyyy hh:nn:ss":
		valor_out = fmt.Sprintf("%d/%d/%d %d:%d:%d", dia, mes, ano, hora, seg, min)
	case "yyyy-mm-dd hh:nn:ss":
		valor_out = fmt.Sprintf("%d-%d-%d %d:%d:%d", ano, mes, dia, hora, seg, min)
	case "yyyy-mm-ddthh:nn:ss.000z":
		valor_out = fmt.Sprintf("%d-%v-%vT%v:%v:%v.000Z", ano, smes, sdia, shora, sseg, smin)
	}
	return valor_out
}

func TimeToDecInt(t time.Time) (d, m, y, h, n, s int) {
	return t.Day(), int(t.Month()), t.Year(), t.Hour(), t.Minute(), t.Second()
}

func TimeToDecStr(t time.Time) (sd, sm, sy, sh, sn, ss string) {
	d, m, y, h, n, s := TimeToDecInt(t)
	sd = strconv.Itoa(d)
	sm = strconv.Itoa(m)
	sy = strconv.Itoa(y)
	sh = strconv.Itoa(h)
	sn = strconv.Itoa(n)
	ss = strconv.Itoa(s)

	if d < 10 {
		sd = "0" + sd
	}

	if m < 10 {
		sm = "0" + sm
	}

	if h < 10 {
		sh = "0" + sh
	}

	if n < 10 {
		sn = "0" + sn
	}

	if s < 10 {
		ss = "0" + ss
	}

	return sd, sm, sy, sh, sn, ss
}

func CheckTime(value_in string) (time.Time, error) {

	val := value_in

	if len(strings.TrimSpace(val)) == 0 {
		return time.Now(), errors.New("Data inválida")
	} else if !strings.Contains(val, "-") {
		return time.Now(), errors.New("Data inválida")
	} else if !strings.Contains(val, " ") {
		return time.Now(), errors.New("Data inválida")
	}

	layout := "2006-01-02T15:04:05.000Z"
	str := strings.Replace(val, " ", "T", -1) + ".000Z"

	DateTime, err := time.Parse(layout, str)

	if err != nil {
		return time.Now(), err
	}

	return DateTime, nil
}
