package main

import (
    "fmt"
    "errors"
)

func input() (string, error) {
   var (
       yourName string
       n int
       err error
   )
   fmt.Print("What is your name? ")
   n, err = fmt.Scan(&yourName)
   
   if err != nil {
       return "", err
   } else if n < 1 {
       err = errors.New("Can't read input")
       return "", err
   } else {
       return yourName, nil
   }
}

func concat(name string, err error) (string) {
   if err != nil {
      return err.Error()
   } else {
       return fmt.Sprintf("Hello, %s, nice to meet you!", name)
   }
}

func output(message string) {
   fmt.Println(message)
}

func main() {
    output(concat(input()))
}
