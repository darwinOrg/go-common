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
