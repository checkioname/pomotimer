package internal

import (
	"fmt"
	"time"
)


type PomoTimer struct {
  Interval time.Duration
  Rest time.Duration
  RemainderTime time.Duration

}


func NewPomoTimer(interval, rest time.Duration) *PomoTimer {
  return &PomoTimer{
    Interval: time.Minute * interval,
    Rest: time.Minute * rest,
    RemainderTime: 0,
  }
}

func (p *PomoTimer) Reset() {
  p.RemainderTime = p.Interval
}

func (p *PomoTimer) StartStudy() {



  p.RemainderTime = p.Interval
  for p.RemainderTime > 0 {
      time.Sleep(time.Second)           // Espera 1 segundo
      p.RemainderTime -= time.Second    // Decrementa 1 segundo do tempo restante
      fmt.Printf("You still got %v of studying\n", p.RemainderTime)
  }

  // for i:= 0; i >= int(p.RemainderTime); p.RemainderTime -= time.Second * 1 {
  //   time.Sleep(time.Second)
  //   fmt.Println("1 second has been passed")    
  // }
  fmt.Println("Good work! Keep pushing!!")
}


func (p *PomoTimer) StartRest () {
  p.RemainderTime = p.Rest
  for i:= 0; i <= int(p.RemainderTime); p.RemainderTime -= time.Second {
    
  }
  fmt.Println("Now it's time to get back to work!! ;)")
}
