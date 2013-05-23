package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/romanoff/fsmonitor"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	Rules map[string]rule
}

type rule struct {
	Path     string
	Pattern  string
	Ignored_Folders     string
	Command  string
	watcher  *fsmonitor.Watcher
}

func (self *rule) watch(name string) error {
	if self.Path == "" {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		self.Path = path
	}
	w, err := fsmonitor.NewWatcherWithSkipFolders(self.getIgnoredFolders())
	self.watcher = w
	if err != nil {
		return err
	}
	err = w.Watch(self.Path)
	if err != nil {
		return err
	}
	go func(r rule) {
		r.handleEvents()
	}(*self)
	return nil
}

func (self *rule) handleEvents() {
	for {
		select {
		case event := <-self.watcher.Event:
			filename := filepath.Base(event.Name)
			for _, p := range self.getPatterns() {
				match, _ := filepath.Match(p, filename)
				if match {
					if !event.IsCreate() {
						self.Execute()
					}
					break
				}
			}
		case err := <-self.watcher.Error:
			fmt.Println(err)
		}
	}
}

func (self *rule) Execute() {
	commands := strings.Split(self.Command, " ")
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (self *rule) getPatterns() []string {
	return strings.Split(self.Pattern, ",")
}

func (self *rule) getIgnoredFolders() []string {
	return strings.Split(self.Ignored_Folders, ",")
}

func ReadConfig(content []byte) (*Config, error) {
	conf := &Config{}
	_, err := toml.Decode(string(content), &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
