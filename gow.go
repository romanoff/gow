package main

import (
	"io/ioutil"
	"log"
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
	for name, rule := range conf.Rules {
		err = rule.watch(name)
		if err != nil {
			log.Fatal(err)
		}
	}
}
