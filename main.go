package main

import (
	"github.com/programzheng/rent-house-crawler/cmd"
	"github.com/programzheng/rent-house-crawler/internal/job"
)

func main() {
	job.RunJobs()
	cmd.Execute()
}
