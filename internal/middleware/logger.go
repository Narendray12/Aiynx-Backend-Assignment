package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"aiynx/internal/logger"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		logger.Log.Info("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("duration", time.Since(start)),
		)

		return err
	}
}
