package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "log"
)

func main() {
  picker()
}

func printWelcome() {
  welcome := []string{"\nWelcome to Colin's Random Table Picker", "Tables are eliminated when number is drawn,", "And put back in the game when their number is drawn again.", "Last number left wins!", "Good Luck.\n"}
  for _, item := range welcome {
    fmt.Println(item)
    time.Sleep(400 * time.Millisecond)
  }
}

func picker() {
  var tables, comp, wait int
  var compString string
  printWelcome()
  for {
    fmt.Print("How many tables are playing: ")
    if _, err := fmt.Scanln(&tables); err != nil {
      log.Fatal(err)
    }
    fmt.Print("How many ms would you like to wait between each iteration: ")
    if _, err := fmt.Scanln(&wait); err != nil {
      log.Fatal(err)
    }
    if tables > 0 {
      break
    }
  }
  tableArray := make([]string, tables)
  for i, _ := range tableArray {
    tableArray[i] = strconv.Itoa(i + 1)
  }
  var total int = 0
  for {
    r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    comp = r.Intn(tables) + 1;
    compString = strconv.Itoa(comp)
    var numX int = 0
    total++
    fmt.Println("Random number is:", compString)
    if tableArray[comp - 1] == compString {
      tableArray[comp - 1] = "x"
    } else {
      tableArray[comp - 1] = compString
    }
    fmt.Println(tableArray)
    for _, element := range tableArray {
      if element == "x" {
        numX++
      }
    }
    if numX >= tables - 1 {
      for _, element := range tableArray {
        if element != "x" {
          fmt.Println("The lucky winner is", element, "chosen after", total, "rounds.\n")
          break
        }
      }
      break
    }
    time.Sleep(time.Duration(wait) * time.Millisecond)
  }
}
