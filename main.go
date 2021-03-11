package main

import (
	"fmt"

	"gitlab.com/detl/detl-common"
)

const (
	maxCPUPct uint8 = 95
	maxMemPct uint8 = 95

	// TODO decide on handling environments, but temporarily, for now:
	env = "dev"
)

func main() {
	conf := detl.GetConf("transform")
	mapping := detl.GetArbitraryYaml("maps/" + conf.Settings["mapping"])

	factory := factory{}
	parser := factory.NewParser(conf.Settings["readParser"], map[string]string{})
	handler := factory.NewHandler(conf.Settings["handler"], mapping, &parser)
	queue := factory.NewQueue(conf.Settings["queue"], conf.Confs[env]["queue"], &handler)

	defer queue.Close()
	queue.Consume()

	fmt.Println(queue, parser)

	fmt.Println(mapping)

	fmt.Println("nuthin")
}
