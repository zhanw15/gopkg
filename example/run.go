package main

import (
	"fmt"

	"github.com/zhanw15/easycmd"
)

func main() {

	// Call Run method with command you want to run on remote server.
	res, err := easycmd.Run("ls /etc/ | grep hosts | grep -v deny")
	// Handle errors
	if err != nil {
		panic("Can't run local command: " + err.Error())
	} else {
		fmt.Println("result is:", res)
	}

}
