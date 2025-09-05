package middleware

import (
	"slices"
	"strings"

	"github.com/connor-davis/threereco-nextgen/internal/constants"
	"github.com/connor-davis/threereco-nextgen/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authorized(requiredPermissions []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser, ok := c.Locals("user").(*models.User)

		if !ok || currentUser == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   constants.UnauthorizedError,
				"details": constants.UnauthorizedErrorDetails,
			})
		}

		combinedPermissions := []string{}

		for _, role := range currentUser.Roles {
			combinedPermissions = append(combinedPermissions, role.Permissions...)
		}

		if slices.Contains(combinedPermissions, "*") {
			return c.Next()
		}

		for _, permission := range combinedPermissions {
			noWildCardPermission := strings.TrimSuffix(permission, ".*")

			for _, requiredPermission := range requiredPermissions {
				if strings.HasPrefix(requiredPermission, noWildCardPermission) {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   constants.ForbiddenError,
			"details": constants.ForbiddenErrorDetails,
		})
	}
}
