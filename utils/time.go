package utils

import (
	"strings"
	"time"
)

func ParseFromOneDateFormatToAnother(str, fromFormat, toFormat string) (string, error) {
	fromTime, err := time.Parse(fromFormat, str)
	if err != nil {
		return "", err
	}

	return fromTime.Format(toFormat), nil
}

func MustParseFromOneDateFormatToAnother(str, fromFormat, toFormat string) string {
	fromTime, _ := ParseFromOneDateFormatToAnother(str, fromFormat, toFormat)

	return fromTime
}

func FormatToChineseDuration(d time.Duration) string {
	str := d.String()
	str = strings.ReplaceAll(str, "h", "小时")
	str = strings.ReplaceAll(str, "m", "分")
	str = strings.ReplaceAll(str, "s", "秒")

	return str
}

func FormatToChineseYearMonthDay(t time.Time) string {
	return t.Format("2006年1月2日")
}

func FormatToChineseMonthDay(t time.Time) string {
	return t.Format("1月2日")
}

func CurrentDate() time.Time {
	return StartOfDay(time.Now())
}

func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func NextDay(t time.Time) time.Time {
	return NextDays(t, 1)
}

func NextDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func PreDay(t time.Time) time.Time {
	return PreDays(t, 1)
}

func PreDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, -days)
}

func NextMonth(t time.Time) time.Time {
	return NextMonths(t, 1)
}

func NextMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

func PreMonth(t time.Time) time.Time {
	return PreMonths(t, 1)
}

func PreMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, -months, 0)
}

func NextYear(t time.Time) time.Time {
	return NextYears(t, 1)
}

func NextYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

func PreYear(t time.Time) time.Time {
	return PreYears(t, 1)
}

func PreYears(t time.Time, years int) time.Time {
	return t.AddDate(-years, 0, 0)
}
