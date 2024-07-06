package utils

import "time"

func Retry(fn func() error, maxCount int, interval time.Duration) {
	for i := 0; i < maxCount; i++ {
		err := fn()
		if err == nil {
			break
		}

		time.Sleep(interval)
	}
}
