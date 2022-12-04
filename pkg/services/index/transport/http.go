package transport

import (
	"blog/pkg/services/users"
	userStore "blog/pkg/services/users/store"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//index handler struct
type handler struct {
	UserService users.Service
}

//Activate function for use services
func Activate(router *fiber.App, db *gorm.DB) {
	userService := users.New(userStore.New(db))
	newHandler(router, userService)
}

//new handler implement
func newHandler(router *fiber.App, us users.Service) {
	h := handler{
		UserService: us,
	}

	router.Get("/", h.index)

}

//Index page method
func (h *handler) index(c *fiber.Ctx) error {
	isLogin, _ := h.UserService.IsLogin(c)
	return c.Render("index/views/templates/index", fiber.Map{"isLogin": isLogin})
}
