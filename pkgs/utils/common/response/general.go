package responses

import "github.com/gofiber/fiber/v2"

type General struct {
	Status  int         `json:"-"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *General) JSON(c *fiber.Ctx) error {
	return c.Status(g.Status).JSON(g)
}
