/*
 * Write a version of the program that displays different greetings for different people.
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

  response := getResponse(name[:len(name)-1])

  fmt.Println(response)

}

func getResponse(name string) string {
  switch name {
  case "Derrick":
    return "You are Derrick"
  case "Melissa":
    return "Pookie"
  }

  return fmt.Sprintf("Hello, %v, nice to meet you!", name)
}
