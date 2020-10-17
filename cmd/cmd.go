package main

import gen "github.com/sdutsoftlab/softlab-generator"

func main() {
	arg := "test"
	gen.Init()

	switch arg {
	case "test":
		gen.Compile()
		gen.ListenHTTPServer(9305)
	case "run":
		gen.ListenHTTPServer(9305)
	}
}
