package dlbindings_test

import (
	"log"

	"github.com/estintax/dlbindings"
)

func ExampleInitDinolang() {
	err := dlbindings.InitDinolang("dinolang.so")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	dlbindings.AddClass("myclass", false, func(args []string, segmentName string) bool {
		switch args[0] {
		case "hello":
			dlbindings.SetReturned("string", "Hello, World!", segmentName)
		}
		return true
	}, nil)

	dlbindings.PiniginShell()
}
