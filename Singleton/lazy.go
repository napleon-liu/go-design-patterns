package Singleton

var ins2 *singleton

func getInstance() *singleton {
	if ins2 == nil {
		ins2 = &singleton{}
	}
	return ins
}
