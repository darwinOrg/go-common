package utils

import "time"

func Retry(maxTimes int, sleepInterval time.Duration, fn func() error) (err error) {
	for i := 0; i < maxTimes; i++ {
		if err = fn(); err == nil {
			return
		}

		time.Sleep(sleepInterval)
	}
	return
}

func RetryWithIndex(maxTimes int, sleepInterval time.Duration, fn func(i int) error) (err error) {
	for i := 0; i < maxTimes; i++ {
		if err = fn(i); err == nil {
			return
		}

		time.Sleep(sleepInterval)
	}
	return
}
