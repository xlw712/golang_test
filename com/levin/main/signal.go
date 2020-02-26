package main

import "sync"

type Signal struct {
	value int
}

var instance *Signal
var lock sync.Mutex
var once sync.Once
/**
 *双重校验
 */
func instance1() *Signal {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = new(Signal)
		}
		lock.Unlock()
	}
	return instance
}
//内部实现也是基于双重校验实现的
func instance2() *Signal{
	once.Do(func() {
		instance=new (Signal)
	})
	return instance
}

func main() {

}
