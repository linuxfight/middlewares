package middlewares

import "github.com/gofiber/fiber/v3"

func NewRecovery() fiber.Handler {
	return func(c fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				_ = c.Status(fiber.StatusInternalServerError).JSON(
					&fiber.Error{
						Message: "internal server error",
						Code:    fiber.StatusInternalServerError},
				)
			}
		}()

		return c.Next()
	}
}
