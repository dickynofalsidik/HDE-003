package main

import (
	"fmt"
	"math"
  "strings"
  "strconv"
  "bufio"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("input:")
  text, _ := reader.ReadString('\n')

  testArray := strings.Fields(text)  
  var kuadrat = 0.0; 
  for _, v := range testArray {
    feetFloat, _ := strconv.ParseFloat(strings.TrimSpace(v), 64)
    if math.Signbit(feetFloat) == false { //positive value
      kuadrat = kuadrat + (feetFloat * feetFloat)
    }
  }
  fmt.Println("output:", kuadrat)
}
