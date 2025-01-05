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
    Interval: interval,
    Rest: rest,
    RemainderTime: 0,
  }
}

func (p *PomoTimer) Reset() {
  p.RemainderTime = p.Interval
}

func (p *PomoTimer) StartStudy() {
  p.RemainderTime = p.Interval
  for i:= 0; i <= int(p.RemainderTime); p.RemainderTime -= time.Second {
    time.Sleep(time.Second)
    fmt.Println("1 second has been passed")    
  }
  fmt.Println("Good work! Keep pushing!!")
}


func (p *PomoTimer) StartRest () {
  p.RemainderTime = p.Rest
  for i:= 0; i <= int(p.RemainderTime); p.RemainderTime -= time.Second {
    
  }
  fmt.Println("Now it's time to get back to work!! ;)")
}
