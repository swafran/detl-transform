package main

import (
	"os"

	detl "github.com/swafran/detl-common"
	"github.com/swafran/detl-transform/factory"
	"github.com/swafran/detl-transform/maps"
)

const (
	maxCPUPct uint8 = 95
	maxMemPct uint8 = 95
)

var env string = os.Getenv("ENV")

func main() {
	conf := detl.GetConf("transform")
	mapping := maps.GetMapping("maps/" + conf.Settings["mapping"])

	parser := factory.NewParser(conf.Settings["readParser"], map[string]string{})
	handler := factory.NewHandler(conf.Settings["handler"], mapping, parser)
	queue := factory.NewQueue(conf.Settings["queue"], conf.Confs[env]["queue"], handler)

	defer queue.Close()
	queue.Consume()
}
