package testing

import (
	"fmt"
	"github.com/Cristien/go-patterns/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Init test")
	setup()
	code := m.Run()
	os.Exit(code)
}
func setup() {
	app.BuildAppContext()
}