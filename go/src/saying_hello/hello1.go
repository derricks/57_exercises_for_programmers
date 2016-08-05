/**
 The “Hello, World” program is the first program you learn
to write in many languages, but it doesn’t involve any input.
So create a program that prompts for your name and prints a greeting using your name.
Example Output
       What is your name? Brian
       Hello, Brian, nice to meet you!

Constraint
  Keep the input, string concatenation, and output sepa- rate.
*/

package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {

  stdinReader := bufio.NewReader(os.Stdin)
  fmt.Print("What is your name? ")
  name, _ := stdinReader.ReadString('\n')

  response := fmt.Sprintf("Hello, %v, nice to meet you!", name[:len(name) - 1])

  fmt.Println(response)
}
