package main

import (
	// "fmt"
	// "strings"

	"github.com/nsf/termbox-go"
)

type mapValue map[string]termbox.Attribute

// func (m mapValue) Set(val termbox.Attribute) error {
// 	v := strings.SplitN(val, "=", 2)
// 	if len(v) != 2 {
// 		return fmt.Errorf("Flag should be in the format of <name>=<value>")
// 	}
//
// 	m[v[0]] = v[1]
//
// 	return nil
// }

// func (m mapValue) String() string {
// 	s := make([]string, 0)
// 	for k, v := range m {
// 		s = append(s, fmt.Sprintf("%s=%s", k, v))
// 	}
//
// 	return strings.Join(s, ",")
// }

func (m mapValue) Type() string {
	return "map"
}
