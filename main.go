package main

import (
	"os"

	detl "github.com/swafran/detl-common"
	"github.com/swafran/detl-transform/maps"
)

const (
	maxCPUPct uint8 = 95
	maxMemPct uint8 = 95
)

func main() {
	env := os.Getenv("ENV")

	conf := detl.GetConf("transform")
	mapping := maps.GetMapping("maps/" + conf.Settings["mapping"])

	parser := NewParser(conf.Settings["readParser"], map[string]string{})
	handler := NewHandler(conf.Settings["handler"], mapping, parser)
	queue := NewQueue(conf.Settings["queue"], conf.Confs[env]["queue"], handler)

	defer queue.Close()
	queue.Consume()
}
