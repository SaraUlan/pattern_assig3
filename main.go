package main

//singleton pattern
import (
	"fmt"
	"sync"
)

type Singleton struct {
	Info int
}

var once sync.Once

type SingletonHolder struct {
	instance *Singleton
}

var instanceHolder *SingletonHolder

func GetInstance() *Singleton {
	once.Do(func() {
		instanceHolder = &SingletonHolder{
			instance: &Singleton{Info: 42},
		}
	})
	return instanceHolder.instance
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			singleton := GetInstance()
			fmt.Printf("get info", id, singleton.Info)
		}(i)
	}
	wg.Wait()
}
