package jsonplan

import (
	"os"
	"testing"
)

func TestNewParser(t *testing.T) {
	p := NewParser()

	if p == nil {
		t.Fatal("Create parser failed")
	}
}

func Test_parser_Parse(t *testing.T) {
	p := NewParser()

	var raw []byte
	var err error
	if raw, err = os.ReadFile("./test/data/plan.json"); err != nil {
		t.Fatal("Read json file failed, ", err)
	}

	var plan *Plan
	if plan, err = p.Parse(raw); err != nil {
		t.Fatalf("Parse json file raw byte failed, %#v", err)
	}

	if plan.ResourceChanges == nil {
		t.Errorf("Expected non null changes, but got: %v", plan)
	}

	if plan.ResourceDrift == nil {
		t.Errorf("Expected non null drift, but got: %v", plan)
	}
}
