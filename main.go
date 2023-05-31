package main

import (
	"log"
	"materi/config"
	"materi/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})
	RouterStatic(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}

func RouterStatic(app *fiber.App) {
	currentTime := time.Now()
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length",
		MaxAge:           86400,
	}))

	limiterConfig := limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}

	app.Use(limiter.New(limiterConfig))

	app.Use(logger.New(logger.Config{
		Next:         nil,
		Done:         nil,
		Format:       `{"pid" : "${pid}", "time" : "${time}", "referer" : "${referer}", "protocol" : "${protocol}", "ip" : "${ip}", "ips" : "${ips}", "host" : "${host}", "method" :"${method}", "path" : "${path}", "url" : "${url}", "ua" : "${ua}", "latency" : "${latency}", "status" : "${status}",},` + "\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Asia/Jakarta",
		TimeInterval: 500 * time.Millisecond,
		Output:       getWriter(currentTime.Format("01-02-2006")),
	}))
	db := config.NewDB()
	router.UserRouter(db, app)

}

func getWriter(level string) *lumberjack.Logger {
	filename := "./" + level + ".log"
	hook := lumberjack.Logger{
		Filename:   filename, // ⽇志⽂件路径
		MaxSize:    1024,     // megabytes
		MaxBackups: 3,        // 最多保留3个备份
		MaxAge:     7,        //days
		Compress:   true,     // 是否压缩 disabled by default
	}

	return &hook

}
