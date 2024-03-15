package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strings"
)

func hasConsecutiveRepeats(number string) bool {
  // Remove hyphens to simplify checking for consecutive repeats
  cleanNumber := strings.Replace(number, "-", "", -1)

  // Check for 4 or more consecutive repeated digits
  for i := 0; i < len(cleanNumber)-3; i++ {
    if cleanNumber[i] == cleanNumber[i+1] && cleanNumber[i+1] == cleanNumber[i+2] && cleanNumber[i+2] == cleanNumber[i+3] {
      return true
    }
  }
  return false
}

func isValidCreditCard(number string) bool {
  // Check if the credit card number starts with 4, 5, or 6 and is exactly 16 digits long,
  // allowing for digits in groups of 4 separated by one hyphen.
  pattern := `^(?:4[0-9]{15}|5[0-9]{15}|6[0-9]{15}|4[0-9]{3}(-[0-9]{4}){3}|5[0-9]{3}(-[0-9]{4}){3}|6[0-9]{3}(-[0-9]{4}){3})$`

  // Compile the regex pattern
  r, err := regexp.Compile(pattern)
  if err != nil {
    fmt.Println("Error compiling regex:", err)
    return false
  }

  // Validate the credit card number structure
  if r.MatchString(number) {
    // Further check for consecutive repeated digits
    if hasConsecutiveRepeats(number) {
      return false
    }
    return true
  }

  return false
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Println("Enter the number of credit card numbers to check:")
  var N int
  if scanner.Scan() {
    fmt.Sscanf(scanner.Text(), "%d", &N)
  }

  fmt.Println("Enter the credit card numbers:")
  for i := 0; i < N; i++ {
    if scanner.Scan() {
      cardNumber := scanner.Text()
      if isValidCreditCard(cardNumber) {
        fmt.Println("Valid")
      } else {
        fmt.Println("Invalid")
      }
    }
  }
}
