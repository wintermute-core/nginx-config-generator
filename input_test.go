package main

import (
	"testing"
)

func TestInputFileLoading(t *testing.T) {

	input := parse("examples/input.yml")

	if input.IpFilter == nil {
		t.Fatal("Ip filter config not parsed")
	}

	if len(input.IpFilter) == 0 {
		t.Fatal("Ip filter config loaded empty")
	}

	if input.CatchAll == nil {
		t.Fatal("CatchAll config not parsed")
	}

	if input.App == nil {
		t.Fatal("App config not parsed")
	}

}
