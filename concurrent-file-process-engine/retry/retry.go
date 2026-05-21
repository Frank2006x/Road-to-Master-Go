package retry

import "time"

func Do(attempts int, fn func() error) error {

	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		time.Sleep(time.Duration(i+1) * time.Second)
	}
	return err
}