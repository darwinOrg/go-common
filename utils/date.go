package utils

import "time"

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
