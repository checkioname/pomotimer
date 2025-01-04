package main

import (
	"fmt"
	"os"

	"github.com/checkioname/pomodoro/internal"
)

func main() {

  register := internal.NewFlagRegistry()

  register.RegisterFlag("t", 45, "Pomo time in minutes")
  

  register.Parse()

  for name, flag := range register.Flags {
        fmt.Printf("Flag: %s, Value: %v (Default: %v), Description: %s\n",
            name, flag.Value, flag.Default, flag.Description)
    }

  argOne := os.Args[1]  
  fmt.Println(argOne)
  

  


  fmt.Println(argOne)

  // timeFlag := flag.Int("t", 1234)
}
