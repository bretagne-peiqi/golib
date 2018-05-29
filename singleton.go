package manager

import (
	"fmt"
	"sync"
)

var m *Manager
var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
