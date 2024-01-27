package repl

import (
  "bufio"
  "io"
  "fmt"
)

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Fprintf(out, PROMPT)
  }
}
