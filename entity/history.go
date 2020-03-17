package entity

import (
	"fmt"
	"time"
)

// History represents the data of history
type History struct {
	Pwd       string
	Command   string
	Timestamp time.Time
}

// Histories is multiple history
type Histories []*History

// Print prints history record
func (hs Histories) Print() {
	var uhs []*History
	m := make(map[string]bool)
	for _, h := range hs {
		if !m[h.Command] {
			m[h.Command] = true
			uhs = append(uhs, h)
		}
	}
	for _, h := range uhs {
		fmt.Println(h.Command)
	}
}
