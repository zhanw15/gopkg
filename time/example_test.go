package time_test

import (
	"fmt"
	"github.com/zhanw15/gopkg/time"
)

func ExampleGetBuildTime() {
	BuildTime := time.GetBuildTime()

	fmt.Println("build time is:", BuildTime)
}