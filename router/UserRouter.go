package router

import (
	"database/sql"
	"materi/controller"
	"materi/repository"
	"materi/service"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(db *sql.DB, base *fiber.App) {
	var (
		UserRepository repository.UserRepository = repository.NewUserRepository(db)
		UserService    service.UserService       = service.NewUserService(UserRepository)
		UserController controller.UserController = controller.NewUserController(UserService)
	)

	root := base.Group("/user")
	root.Get("/", UserController.All)
	root.Get("/:id_user", UserController.FindByID)
	root.Post("/", UserController.Create)
	root.Put("/:id_user", UserController.Update)
	root.Put("/profil/:id_user", UserController.Upload)
	root.Delete("/:id_user", UserController.Delete)
}
