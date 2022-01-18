package jsonplan

import (
	"fmt"
	"log"
	"os"
)

func ExampleParser_Parse() {
	p := NewParser()

	var raw []byte
	var err error
	if raw, err = os.ReadFile("./test/data/plan.json"); err != nil {
		log.Fatalf("Read json file failed: %v", err)
	}

	var plan *Plan
	if plan, err = p.Parse(raw); err != nil {
		log.Fatalf("Parse json file raw byte failed, %#v", err)
	}

	fmt.Println(plan.FormatVersion)
	fmt.Println(len(plan.ResourceChanges))
	fmt.Println(len(plan.ResourceDrift))

	// Output:
	// 0.2
	// 8
	// 1
}
