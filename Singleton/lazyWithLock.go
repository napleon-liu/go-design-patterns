package Singleton

import "sync"

var ins3 *singleton
var m sync.Mutex

func GetInstance() *singleton {
	m.Lock()
	if ins3 == nil {
		ins3 = &singleton{}
	}
	m.Unlock()
	return ins3
}
