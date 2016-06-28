package main

import (
	"log"
	"os"
	"os/exec"
)

// Entry is a cron entry
type Entry struct {
	When string `json:"when"`
	Cmd  string `json:"cmd"`
}

// Run makes Entry implement the go-cron job interface
func (en Entry) Run() {

	cmd := exec.Command(en.Cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("[tcron] command: sh -c %s", en.Cmd)
	err := cmd.Run()
	if err != nil {
		log.Printf("[tcron] error %s", err)
	}
}
