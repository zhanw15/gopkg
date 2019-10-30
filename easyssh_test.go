package easyssh

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	fmt.Println( Run("ls"))
}
