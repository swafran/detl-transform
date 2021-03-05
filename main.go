package main

import (
	"fmt"

	"gitlab.com/detl/detl-common"
	"gitlab.com/detl/detl-common/queues"
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

	// should this be here or in each goroutine ?
	queue := NewService(conf.Settings["queue"], conf.Confs[env]["queue"]).(queues.RabbitQueue)

	// parser = NewService(conf.Settings["readParser"], map[string]string{})

	// //// wrap in goroutine
	raw := queue.Consume()
	// input := parser.Parse(raw)
	// transformed := maps.Resolve(input, mapping)
	// formatted := formatter.Format(transformed)
	// queue.Publish(formatted)
	// ////

	queue.Close()

	fmt.Println(queue)

	fmt.Println(raw)

	fmt.Println("nuthin")
}
