package internal

import (
	"flag"
	"fmt"
)

type Flag struct {
	Name        string
	Default     interface{}
	Value       interface{}
	Description string
}

type FlagRegister struct {
	Flags   map[string]*Flag
	FlagSet *flag.FlagSet
}

func NewFlagRegistry(name string) *FlagRegister {
  fr := &FlagRegister{
		Flags:   make(map[string]*Flag),
		FlagSet: flag.NewFlagSet(name, flag.ExitOnError),
	}
  fr.Flags["t"] = &Flag{ Name: "t", Default: 45, Value: 45, Description: "Pomo time in minutes"}
  fr.Flags["r"] = &Flag{ Name: "r", Default: 5, Value: 5, Description: "Pomo rest time in minutes"}
  
  fr.FlagSet.Int("t", 45, "Pomo time in minutes")
  fr.FlagSet.Int("r", 5, "Pomo rest time in minutes")
  return fr
}

func (fr *FlagRegister) RegisterFlag(name string, defaultValue interface{}, description string) {
	fr.Flags[name] = &Flag{
		Name:        name,
		Default:     defaultValue,
		Value:       defaultValue,
		Description: description,
	}

	// Registra a flag no FlagSet
	switch defaultValue := defaultValue.(type) {
	case bool:
		fr.FlagSet.Bool(name, defaultValue, description)
	case string:
		fr.FlagSet.String(name, defaultValue, description)
	case int:
		fr.FlagSet.Int(name, defaultValue, description)
	default:
		fmt.Printf("Unsupported flag type for %s\n", name)
	}
}

func (fr *FlagRegister) Parse(args []string) {
  fmt.Println("Parse iniciando")
  err := fr.FlagSet.Parse(args)
  if err != nil {
    fmt.Printf("Houve um erro ao parsear as flags: %v", err)
  }

	// Atualiza os valores das flags no map `Flags`
	for name, f := range fr.Flags {
    fmt.Printf("Nome da flag sendo parseada atualmente: %v", name)
		if flagValue := fr.FlagSet.Lookup(name); flagValue != nil {
			switch f.Default.(type) {
			case bool:
				f.Value = flagValue.Value.(flag.Getter).Get().(bool)
			case string:
				f.Value = flagValue.Value.(flag.Getter).Get().(string)
			case int:
				f.Value = flagValue.Value.(flag.Getter).Get().(int)
			}
		}
	}
}
