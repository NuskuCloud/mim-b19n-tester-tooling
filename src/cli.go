package main

var CLI struct {
	Ch struct {
		Value string `arg:"" name:"value" help:"[on, off] Write's 1/0 to register 52.'"`
	} `cmd:"" help:"Enables/disables the central heating"`

	Listen struct {
		Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
	} `cmd:"" help:"[NOT IMPLEMENTED] Listens to for messages."`
}
