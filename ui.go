// Demo code for the TextArea primitive.
package main

import (
	"tui/app"
	"tui/flag"
)

func main() {
	tImg := flag.ReadFlag()
	app.Run(tImg)
}
