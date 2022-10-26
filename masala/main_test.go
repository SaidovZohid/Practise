package masala

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	fmt.Println("Main Test")
	os.Exit(m.Run())
}