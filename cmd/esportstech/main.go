// cmd/esportstech/main.go
package main

import (
"flag"
"log"
"os"

"esportstech/internal/esportstech"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := esportstech.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
