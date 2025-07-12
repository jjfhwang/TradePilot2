// cmd/tradepilot2/main.go
package main

import (
"flag"
"log"
"os"

"tradepilot2/internal/tradepilot2"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := tradepilot2.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
