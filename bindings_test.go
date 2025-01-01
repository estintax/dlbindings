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
	ret := dlbindings.AddClass("myclass", false, func(args []string, segmentName string) bool {
		switch args[0] {
		case "hello":
			dlbindings.SetReturned("string", "Hello, World!", segmentName)
		}
		return true
	}, nil)
	if !ret {
		t.Fatal("Add class failed")
	}

	dlbindings.RunCode("use, \"myclass\"\nmyclass:hello\n")
	if dlbindings.GetVariableValue("returned").(string) != "Hello, World!" {
		t.Fatal("expected \"Hello, World!\"")
	}

	if !dlbindings.SetClassUsage("myclass", false, false) {
		t.Fatal("exptected true in SetClassUsage")
	}

	dlbindings.CleanUp(true)
}
