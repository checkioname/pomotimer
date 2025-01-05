package main

import (
	"fmt"

	"github.com/checkioname/pomodoro/internal"
)

func main() {

  register := internal.NewFlagRegistry()

  register.RegisterFlag("t", 45, "Pomo time in minutes")
  register.RegisterFlag("r", 5, "Rest time in minutes")

  register.Parse()

	for name, flag := range register.Flags {
		fmt.Printf("Flag: %s, Value: %v (Default: %v), Description: %s\n",
			name, flag.Value, flag.Default, flag.Description)
  }
  

  // pomo := internal.NewPomoTimer(register.Flags["t"].Value, register.Flags["r"].Value)
  // for {
  //   pomo.StartStudy()
  //
  // }
  //

  // timeFlag := flag.Int("t", 1234)
}
