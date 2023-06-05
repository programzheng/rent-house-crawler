package main

import (
	"context"

	"github.com/programzheng/rent-house-crawler/internal"
)

func main() {
	ctx := context.Background()
	// internal.SaveHomeData(ctx)
	internal.UpsertAllHomeByDetailData(ctx)
}
