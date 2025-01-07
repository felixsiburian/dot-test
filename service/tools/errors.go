package tools

import (
	"errors"
	"log"
	"runtime"
)

func Wrap(err error) error {
	_, f, l, _ := runtime.Caller(1)
	log.Printf("on [%s:%d]: %v", f, l, err)

	return errors.New(err.Error())
}
