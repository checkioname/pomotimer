package main

import (
	"fmt"
	"os"
	"time"

	"github.com/checkioname/pomodoro/internal"
)

func main() {

  register := internal.NewFlagRegistry("PomoTimer")

  args := os.Args[1:]
  
  fmt.Println("Parse iniciando")
  register.Parse(args)
  fmt.Println(args)

	for name, flag := range register.Flags {
		fmt.Printf("Flag: %s, Value: %v (Default: %v), Description: %s\n",
			name, flag.Value, flag.Default, flag.Description)
  }
  
  rest := time.Duration(register.Flags["r"].Value.(int))

  fmt.Printf("Value of flag r: %v\n", rest)
  studyTime := time.Duration(register.Flags["t"].Value.(int))
  pomo := internal.NewPomoTimer(studyTime, rest)

  for {
    done := pomo.StartStudy()
    if done {
      fmt.Println("Now it's rest time!")
    }
    done = pomo.StartRest()
    if done {
      break
    }
  }

}
  // timeFlag := flag.Int("t", 1234)
