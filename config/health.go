package config

import "github.com/gofiber/fiber/v2"

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health
// @Accept */health*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	// Check DB connection
	res := map[string]interface{}{
		"status": "success",
		"database": map[string]interface{}{
			"status": "connected",
		},
		"data":    "Server is up and running",
		"version": GetVersion(),
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
