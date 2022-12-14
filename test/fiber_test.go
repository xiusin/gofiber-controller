package test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/xiusin/gofiber-controller/wrapper"
)

type TestController struct {
	wrapper.Controller

	dataType string
}

func (*TestController) InitGroupRouter(w *wrapper.GroupRouter) {
	w.Get("/list", "List")
}

func (*TestController) List() error {
	return nil
}

func Test_App(t *testing.T) {

	app := fiber.New()

	routerWrapper := wrapper.New(app)
	routerWrapper.WrapperHandler("/common", new(TestController))

	app.Listen(":3001")
}
