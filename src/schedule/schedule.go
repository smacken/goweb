package schedule

import (
	"github.com/robfig/cron"
)

func init() {
	schedule := cron.New()
	schedule.AddFunc("@daily", daily)
	schedule.AddFunc("@weekly", weekly)
	schedule.AddFunc("@monthly", monthly)
	schedule.Start()
}

func daily() {
	// backup
	//diagnostics
	// logging
}

func weekly() {

}

func monthly() {

}
