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
	Flags map[string]*Flag
}

func NewFlagRegistry() *FlagRegister {
	return &FlagRegister{Flags: make(map[string]*Flag)}
}

func (fr *FlagRegister) RegisterFlag(name string, defaultValue interface{}, description string) {
	fr.Flags[name] = &Flag{
		Name:        name,
		Default:     defaultValue,
		Value:       defaultValue,
		Description: description,
	}
}

func (fr *FlagRegister) Parse() {
	for name, flag := range fr.Flags {
		switch defaultvalue := flag.Default.(type) {
		case bool:
			flag.Value = *flagBool(name, defaultvalue, flag.Description)
		case string:
			flag.Value = *flagString(name, defaultvalue, flag.Description)
		case int:
			flag.Value = *flagInt(name, defaultvalue, flag.Description)
		default:
			fmt.Printf("Unsupported flag type %s/n", name)
		}
	}
	flag.Parse()
}

func flagBool(name string, defaultValue bool, description string) *bool {
	val := flag.Bool(name, defaultValue, description)
	return val
}

func flagString(name string, defaultValue string, description string) *string {
	val := flag.String(name, defaultValue, description)
	return val
}

func flagInt(name string, defaultValue int, description string) *int {
	val := flag.Int(name, defaultValue, description)
	return val
}

// const (
//
//   timeFlag = flag.Int("t", 1234, "Set how much time for study");
//
// )
