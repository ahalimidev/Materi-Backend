package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func WilayahRouter(db *sql.DB, base *fiber.App) {
	var (
		WilayahRepository repository.WilayahRepository = repository.NewWilayahRepository(db)
		WilayahService    service.WilayahService       = service.NewWilayahService(WilayahRepository)
		WilayahController controller.WilayahController = controller.NewWilayahController(WilayahService)
	)

	root := base.Group("/wilayah")
	root.Get("/provinsi", WilayahController.Provinsi)
	root.Get("/kabupaten/:provinsi_id", WilayahController.Kabupaten)
	root.Get("/kecamatan/:kabupaten_id", WilayahController.Kecamatan)
	root.Get("/desa/:kecamatan_id", WilayahController.Desa)
}
