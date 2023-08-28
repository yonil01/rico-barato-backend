package api

import (
	"backend-ccff/api/handler/attendance"
	"backend-ccff/api/handler/events"
	"backend-ccff/api/handler/modules"
	"backend-ccff/api/handler/reniec"
	"backend-ccff/api/handler/report"
	Roles "backend-ccff/api/handler/roles"
	"backend-ccff/api/handler/users"
	"github.com/ansrivas/fiberprometheus/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// @title API E11MTI.
// @version 1.0.0
// @description Documentation Api E11 MTI.
// @termsOfService https://www.nexumSign.com/terms/
// @contact.name API Support.
// @contact.email info@e-capture.co
// @license.name Software Owner
// @license.url http://www.ecapture.com.co
// @host localhost:50070
// @BasePath /
func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("API-E11-MTI")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "swagger/doc.json",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()

	reniec.RouterReniec(app, db, TxID)
	users.RouterUsers(app, db, TxID)
	modules.RouterModules(app, db, TxID)
	events.RouterEvents(app, db, TxID)
	report.RouterReport(app, db, TxID)
	attendance.RouterAttendance(app, db, TxID)
	Roles.RouterRoles(app, db, TxID)
	return app
}
