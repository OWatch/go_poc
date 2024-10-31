package designpattern

import "sync"

var instance *Singleton
var once sync.Once

type Singleton struct{}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
