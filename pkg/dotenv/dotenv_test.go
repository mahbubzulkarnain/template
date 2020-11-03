package dotenv_test

import (
	"fmt"
	"testing"

	"template-go/pkg/dotenv"
)

func TestLoad(t *testing.T) {
	goenv := dotenv.Load("GO_ENV")
	fmt.Println(goenv)
}
