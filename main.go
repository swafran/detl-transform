package main

import (
	"fmt"

	"gitlab.com/detl/detl-common"
	"gitlab.com/detl/transform/factory"
	"gitlab.com/detl/transform/maps"
)

const (
	maxCPUPct uint8 = 95
	maxMemPct uint8 = 95

	// TODO decide on handling environments, but temporarily, for now:
	env = "dev"
)

func main() {
	conf := detl.GetConf("transform")
	mapping := maps.GetYamlMapping("maps/" + conf.Settings["mapping"])

	parser := factory.NewParser(conf.Settings["readParser"], map[string]string{})
	handler := factory.NewHandler(conf.Settings["handler"], mapping, parser)
	queue := factory.NewQueue(conf.Settings["queue"], conf.Confs[env]["queue"], handler)

	defer queue.Close()
	queue.Consume()

	fmt.Println("nuthin")

	fmt.Println("nuthin")

	return
}
