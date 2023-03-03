package timer

import (
	"code-market-admin/internal/app/blockchain"
	"code-market-admin/internal/app/service"
	"github.com/robfig/cron/v3"
)

func Timer() {
	pendingWithdrawCron()
	statsCron()
}

func pendingWithdrawCron() {
	t := cron.New()
	t.AddFunc("*/15 * * * *", func() {
		_ = blockchain.UpdateAllPendingWithdraw()
	})
	t.Start()
}

func statsCron() {
	go service.UpdateStats()
	t := cron.New()
	t.AddFunc("*/15 * * * *", func() {
		service.UpdateStats()
	})
	t.Start()
}
