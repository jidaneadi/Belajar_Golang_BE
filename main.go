package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jidaneadi/projectkp-backend/controllers/authcontroller"
	"github.com/jidaneadi/projectkp-backend/controllers/joincontroller"
	"github.com/jidaneadi/projectkp-backend/controllers/joinsuratcontroller"
	"github.com/jidaneadi/projectkp-backend/controllers/masyarakatcontroller"
	"github.com/jidaneadi/projectkp-backend/controllers/suratcontroller"
	"github.com/jidaneadi/projectkp-backend/controllers/usercontroller"
	"github.com/jidaneadi/projectkp-backend/middleware"
	"github.com/jidaneadi/projectkp-backend/models"
)

func main() {
	models.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	user := api.Group("/user")
	masyarakat := api.Group("/masyarakat")
	join := api.Group("/join")
	surat := api.Group("/surat")
	auth := api.Group("/auth")
	penyuratan := api.Group("/data")

	user.Get("/", middleware.Auth, usercontroller.Index)
	user.Get("/:id", usercontroller.ShowId)
	user.Post("/", usercontroller.Create)
	user.Put("/:id", usercontroller.Update)
	user.Delete("/:id", usercontroller.DeleteId)

	masyarakat.Get("/", masyarakatcontroller.Index)
	masyarakat.Get("/:nik", masyarakatcontroller.ShowId)
	masyarakat.Post("/", masyarakatcontroller.Create)
	masyarakat.Put("/:nik", masyarakatcontroller.Update)
	masyarakat.Delete("/:nik", masyarakatcontroller.Delete)

	join.Get("/", joincontroller.IndexJoin)
	join.Get("/:idm", joincontroller.ShowIdJoin)

	penyuratan.Get("/", joinsuratcontroller.JoinSurat)
	penyuratan.Get("/:nik", joinsuratcontroller.JoinSuratId)

	surat.Get("/", suratcontroller.Index)
	surat.Get("/:id", suratcontroller.IndexId)
	surat.Post("/", suratcontroller.Create)
	surat.Put("/:id", suratcontroller.Update)
	surat.Delete("/:id", suratcontroller.Delete)

	auth.Post("/", authcontroller.Register)
	auth.Post("/login", authcontroller.Login)
	auth.Post("/refresh-token", authcontroller.RefreshToken)

	app.Listen(":3005")
}
