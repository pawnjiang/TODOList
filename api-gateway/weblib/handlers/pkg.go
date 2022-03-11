package handlers

import (
	"errors"
)

//包装错误PanicIfUserError
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		//logging.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		//logging.Info(err)
		panic(err)
	}
}
