package main

import (
	"fmt"

	detl "gitlab.com/detl/detl-common"
)

const (
	maxCPUPct uint8 = 95
	maxMemPct uint8 = 95
)

func main() {
	job := detl.GetSettings("transformer")

	//a := job["settings"]["queue"]

	// queueReader := NewService(job["settings"]["queue"])

	fmt.Println(job)
}
