package GoLibs

import (
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
	switch strings.ToLower(strings.TrimSpace(formato)) {
	case "dd/mm/yyyy":
		valor_out = fmt.Sprint("%d/%d/%d", t.Day(), int(t.Month()), t.Year())
	case "dd/mm/yyyy hh:nn:ss":
		valor_out = fmt.Sprintf("%d/%d/%d %d:%d:%d", t.Day(), int(t.Month()), t.Year(), t.Hour(), t.Second(), t.Nanosecond())
	case "yyyy-mm-dd hh:nn:ss":
		valor_out = fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
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
