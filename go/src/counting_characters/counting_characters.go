/**
 * Create a program that prompts for an input string and dis- plays output that shows the input string and the number of characters the string contains.
 * Be sure the output contains the original string.
• Use a single output statement to construct the output.
• Useabuilt-infunctionoftheprogramminglanguagetodetermine the length of the string.
 */

package main

import (
  "book_utils"
  "fmt"
)

func main() {
  inputString := book_utils.PromptForResponse("Enter a string")
  fmt.Printf("%v is %d characters long\n", inputString, len(inputString))
}
