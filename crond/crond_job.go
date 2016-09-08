package main

import (
	"fmt"
	"github.com/robfig/cron"
)

type Job struct {
	id      int
	name    string
	runFunc func(string) string
}

func (j *Job) Run() {
	//t := Newtest()
	fmt.Println(j.runFunc("testfunc"))
}

func Newtest() *Job {
	test := &Job{id: 1, name: "lupuxiao"}
	test.runFunc = func(name string) string {
		title := fmt.Sprintf("id: %d name: %s you func ret is : %s", test.id, test.name, name)
		return title
	}
	return test
}

var (
	mainCron *cron.Cron
)

func AddJob() bool {
	mainCron = cron.New()
	mainCron.Start()
	spec := "*, *, *, *, *, *"
	job := Newtest()
	err := mainCron.AddJob(spec, job)
	if err != nil {
		return false
	}
	return true
}

func main() {
	if AddJob() {
		fmt.Println("ok")
	}
	select {}
}
