package utils

import "fmt"

func ErrWithCode(err error, code int) error {
	return fmt.Errorf("#%02d#%s", code, err.Error())
}
