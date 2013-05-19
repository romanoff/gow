package main

import (
	"io/ioutil"
	"log"
	"sync"
)

func main() {
	content, err := ioutil.ReadFile(".gow")
	if err != nil {
		log.Fatal(err)
	}
	conf, err := ReadConfig(content)
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	for name, rule := range conf.Rules {
		wg.Add(1)
		err = rule.watch(name)
		if err != nil {
			log.Fatal(err)
		}
	}
	wg.Wait()
}
