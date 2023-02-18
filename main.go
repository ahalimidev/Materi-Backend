package main

import (
	"log"
	"materi/config"
	"materi/router"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/helmet/v2"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	db := config.NewDB()
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
	})
	RouterStatic(app)
	router.WilayahRouter(db, app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}

func RouterStatic(app *fiber.App) {
	currentTime := time.Now()

	app.Use(helmet.New())

	csrfConfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
		CookieName:     "my_csrf_",            // name of the session cookie
		CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
		Expiration:     3 * time.Hour,         // expiration is the duration before CSRF token will expire
		KeyGenerator:   utils.UUID,            // creates a new CSRF token
	}

	app.Use(csrf.New(csrfConfig))

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
