package dlbindings_test

import (
	"testing"

	"github.com/estintax/dlbindings"
)

func TestInitDinolang(t *testing.T) {
	err := dlbindings.InitDinolang("dinolang.so")
	if err != nil {
		t.Fatal(err.Error())
	}
}
