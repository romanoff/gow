package main

import (
	"testing"
)

var testData string = `
[rules]
  [rules.sass]
  pattern="*.sass,*.css"
  command="make regenerate_css"
  ignored_folders=".git"

  [rules.js]
  path="/path/to/folder/with/javascripts"
  pattern="*.js"
  command="make regenerate_js"
`

func TestConfig(t *testing.T) {
	conf, err := ReadConfig([]byte(testData))
	if err != nil {
		t.Errorf("Expected not to have error parsing toml config, but got %v", err)
	}
	if (conf.Rules["sass"] == rule{}) {
		t.Errorf("Expected sass rules to be present in config")
	}
	if (conf.Rules["js"] == rule{}) {
		t.Errorf("Expected js rules to be present in config")
	}
}
