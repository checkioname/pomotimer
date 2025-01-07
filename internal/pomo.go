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

func (p *PomoTimer) StartStudy() bool {
  p.RemainderTime = p.Interval
  for p.RemainderTime > 0 {
      time.Sleep(time.Second)           // Espera 1 segundo
      p.RemainderTime -= time.Second    // Decrementa 1 segundo do tempo restante
      fmt.Printf("You still got %v of studying\n", p.RemainderTime)
  }

  fmt.Println("Good work, time is done! Keep pushing!!")
  return true
}


func (p *PomoTimer) StartRest () bool {
  p.RemainderTime = p.Rest

  for p.RemainderTime > 0 {
      time.Sleep(time.Second)           // Espera 1 segundo
      p.RemainderTime -= time.Second    // Decrementa 1 segundo do tempo restante
      fmt.Printf("You still got %v of studying\n", p.RemainderTime)
  }
  fmt.Println("Now it's time to get back to work!! ;)")
  return true
}
