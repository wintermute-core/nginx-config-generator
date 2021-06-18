package main

import (
	"fmt"
	"testing"
)

func TestInputFileLoading(t *testing.T) {

	input := parse("examples/input.yml")

	if input.IpFilter == nil {
		t.Error("Ip filter config not parsed")
	}

	if len(input.IpFilter) == 0 {
		t.Error("Ip filter config loaded empty")
	}

	if input.CatchAll == nil {
		t.Error("CatchAll config not parsed")
	}

	if input.App == nil {
		t.Error("App config not parsed")
	}

	fmt.Printf("%v", input)
}
