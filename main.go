package main

import (
	"fmt"

	detl "gitlab.com/detl/detl-common"
)

const (
	maxCPUPct uint8 = 95
	maxMemPct uint8 = 95

	// decide on handling environments, but for now temporarily:
	env = "dev"
)

func main() {
	conf := detl.GetConf("transform")
	mapping := detl.GetArbitraryYaml("mapping/" + conf.Settings["mapping"])

	// should this be here or in each goroutine ?
	queue := NewService(conf.Settings["queue"], conf.Confs[env]["queue"])
	// queue.Init()
	// parser = NewService(conf.Settings["readParser"], map[string]string{})

	// //// wrap in goroutine
	// raw := queue.Consume()
	// input := parser.Parse(raw)
	// transformed := maps.Resolve(input, mapping)
	// formatted := formatter.Format(transformed)
	// queue.Publish(formatted)
	// ////

	fmt.Println(conf)

	fmt.Println(mapping)

	fmt.Println("nuthin")
}
