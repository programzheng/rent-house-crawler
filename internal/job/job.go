package job

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/programzheng/rent-house-crawler/config"
	"github.com/programzheng/rent-house-crawler/internal"
)

func RunJobs() {
	s := gocron.NewScheduler(time.Now().Local().Location())
	s.Cron(config.Cfg.GetString("jobs.update-591-homes")).Do(UpsertRent591Homes) // every daily
	s.StartAsync()
}

func UpsertRent591Homes() {
	ctx := context.Background()
	internal.SaveHomeData(ctx)
	internal.UpsertAllHomeByDetailData(ctx)
}
