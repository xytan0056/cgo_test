package main

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

/*
#include <add.h>
*/
import "C"

type test struct {
	name   string
	number int
	desc   string
}

type FilterOption func(*test, *test)

type fieldConfig struct {
	filterOption FilterOption
	otherstuff   interface{}
}

func FilterName() FilterOption {
	return func(new *test, old *test) {
		new.name = old.name
	}
}

func FilterNumber() FilterOption {
	return func(new *test, old *test) {
		new.number = old.number
	}
}

func Filter(t *test, opts ...FilterOption) *test {
	n := new(test)
	for _, opt := range opts {
		opt(n, t)
	}
	return n
}


func running() error {
	fmt.Println("running")
	time.Sleep(time.Second * 6)
	return nil
}

// Plugins
type Plugins struct {
	Static map[string][]Directive `yaml:",inline"`
}

type Directive struct {
	Action string
	Values interface{}
}

func (D Directive) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		D.Action: D.Values,
	}, nil
}

func (p *Directive) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		Content map[string]interface{} `yaml:",inline"`
	}
	if err := value.Decode(&tmp); err != nil {
		return err
	}

	if len(tmp.Content) != 1 {
		return fmt.Errorf("err, but got %d", len(tmp.Content))
	}
	for key, config := range tmp.Content {
		p.Action = key
		p.Values = config
		break
	}
	return nil
}

func main() {
	// 	yamlData := []byte(`
	// default:
	//   - add:
	//     - p1#v0.2.0:
	//     - sp2#v2.0.14:
	//         timeout: 300
	// `)

	// 	var data Plugins
	// 	err := yaml.Unmarshal(yamlData, &data)
	// 	if err != nil {
	// 		fmt.Println("Error unmarshalling YAML:", err)
	// 		return
	// 	}

	// 	yamlBytes, err := yaml.Marshal(data)
	// 	if err != nil {
	// 		fmt.Println("Error marshalling YAML:", err)
	// 		return
	// 	}

		// fmt.Println(string(yamlBytes))
		fmt.Println("wow")
	// Call the C function from Go
	result := C.add(3, 7)
	fmt.Printf("Result from C: %d\n", result)

}
