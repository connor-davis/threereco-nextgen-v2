package middleware

import (
	"time"

	"github.com/connor-davis/threereco-nextgen/internal/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func (m *Middleware) Authenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentSession, err := m.Sessions.Get(c)

		if err != nil {
			log.Errorf("🔥 Error retrieving session: %s", err.Error())

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   constants.UnauthorizedError,
				"details": constants.UnauthorizedErrorDetails,
			})
		}

		currentUserId, ok := currentSession.Get("user_id").(string)

		if !ok || currentUserId == "" {
			log.Warn("🚫 Unauthorized access attempt: No user ID in session")

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   constants.UnauthorizedError,
				"details": constants.UnauthorizedErrorDetails,
			})
		}

		currentUserIdUUID, err := uuid.Parse(currentUserId)

		if err != nil {
			log.Errorf("🔥 Error parsing user ID: %s", err.Error())

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   constants.UnauthorizedError,
				"details": constants.UnauthorizedErrorDetails,
			})
		}

		log.Infof("🔐 Authorized User with ID: %s", currentUserIdUUID)

		currentUser, err := m.Services.Users().Find(currentUserIdUUID)

		if err != nil {
			log.Errorf("🔥 Error retrieving user: %s", err.Error())

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   constants.UnauthorizedError,
				"details": constants.UnauthorizedErrorDetails,
			})
		}

		if currentUser == nil {
			log.Warn("🚫 Unauthorized access attempt: User not found")

			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   constants.UnauthorizedError,
				"details": constants.UnauthorizedErrorDetails,
			})
		}

		c.Locals("user_id", currentUser.Id.String())
		c.Locals("user", currentUser)

		currentSession.Set("user_id", currentUser.Id.String())
		currentSession.SetExpiry(1 * time.Hour)

		if err := currentSession.Save(); err != nil {
			log.Errorf("🔥 Error saving session: %s", err.Error())

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   constants.InternalServerError,
				"details": constants.InternalServerErrorDetails,
			})
		}

		return c.Next()
	}
}
