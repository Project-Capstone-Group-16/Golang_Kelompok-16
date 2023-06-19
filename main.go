package main

import (
	"Capstone/config"
	"Capstone/middleware"
	"Capstone/routes"
	"Capstone/usecase"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/robfig/cron/v3"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("failed to load .env")
	// }

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	db := config.InitDB()
	e := echo.New()

	routes.Routes(e, db)
	middleware.Logmiddleware(e)

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	go func() {
		JobUpdateStatusDone()
	}()

	select {}
}

func JobUpdateStatusDone() {
	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	// set task yang akan dijalankan scheduler
	// gunakan crontab string untuk mengatur jadwal
	scheduler.AddFunc("1 7 * * 1-7", usecase.UpdateStatusDone)

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
} // new
