/**
 * Write a new version of the program without using any variables.
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

 func readInput() string {
   
 }
