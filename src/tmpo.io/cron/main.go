package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/robfig/cron"
)

var (
	cronFile = flag.String("cron", "cron.json", "Crontab json file")
)

func main() {
	flag.Parse()
	entries := loadJSONCron()
	cr := cron.New()
	for _, en := range entries {
		cr.AddJob(en.When, en)
	}

	cr.Start()

	select {}

}

func loadJSONCron() []*Entry {
	log.Printf("Loading json file %s\n", *cronFile)
	file, err := ioutil.ReadFile(*cronFile)
	if err != nil {
		log.Fatalf("Invalid config %s: %s", *cronFile, err)
	}
	var opts []*Entry
	err = json.Unmarshal(file, &opts)
	if err != nil {
		log.Fatalf("Invalid config %s: %s", *cronFile, err)
	}
	return opts
}
