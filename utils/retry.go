package utils

import "time"

func Retry(maxTimes int, sleepInterval time.Duration, fn func() error) {
	for i := 0; i < maxTimes; i++ {
		if err := fn(); err == nil {
			return
		}

		time.Sleep(sleepInterval)
	}
}
