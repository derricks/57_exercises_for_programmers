package book_utils

import (
  "bufio"
  "fmt"
  "os"
)

/** Given a string, use that as a prompt to the user and return their response
*/
func PromptForResponse(prompt string) string {
  normalizedName := ""

  for ;len(normalizedName) == 0; {
    stdinReader := bufio.NewReader(os.Stdin)
    fmt.Printf("%v ", prompt)
    name, _ := stdinReader.ReadString('\n')
    normalizedName = name[:len(name) - 1]
  }
  return normalizedName
}
