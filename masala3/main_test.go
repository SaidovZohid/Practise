package masala3

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	fmt.Println("Main Test Run")
	os.Exit(m.Run())
}