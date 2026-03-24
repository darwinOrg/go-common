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

func RetryWithResult[T any](maxTimes int, sleepInterval time.Duration, fn func() (T, error)) (t T, err error) {
	for i := 0; i < maxTimes; i++ {
		if t, err = fn(); err == nil {
			return t, nil
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
