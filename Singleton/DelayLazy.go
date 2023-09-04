package Singleton

import "sync"

var ins4 *singleton
var once sync.Once

func getInsOr() *singleton {
	once.Do(func() {
		ins4 = &singleton{}
	})

	return ins4
}
