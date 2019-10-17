package main

import (
	"testing"

	"github.com/DATA-DOG/godog"
)

func TestHello(t *testing.T) {
	t.Run("test", func(t *testing.T) {
	})
}

func iHaveAGodogsTest() error {
	return nil
}

func iRunIt() error {
	return nil
}

func itShouldPass() error {
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a godogs test$`, iHaveAGodogsTest)
	s.Step(`^I run it$`, iRunIt)
	s.Step(`^it should pass$`, itShouldPass)
}
