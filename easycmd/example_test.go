package easycmd_test

import (
	"fmt"
	"github.com/zhanw15/gopkg/easycmd"
)

func ExampleRun() {
	// Call Run method with command you want to run on local.
	res, err := easycmd.Run("ls /etc/ | grep hosts | grep -v deny")
	// Handle errors
	if err != nil {
		fmt.Println("Can't run local command: " + err.Error())
		return
	} else {
		fmt.Println("result is:", res)
	}
}