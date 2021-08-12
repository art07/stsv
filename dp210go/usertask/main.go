package main

import (
	"art/edu/stsv/dp210go/usertask/user"
	"fmt"
)

func main() {
	// Начать диалог с человеком, который еще не заведен в системе как пациент.
	if err := user.StartDialog(); err != nil {
		_ = fmt.Errorf("%s\n", err)
	}
}
