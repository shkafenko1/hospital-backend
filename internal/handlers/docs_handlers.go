package handlers

import (
	"github.com/gofiber/fiber/v3"
)

const redocHTML = `<!DOCTYPE html>
<html>
  <head>
    <title>Hospital API – Redoc</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet"/>
    <style>body { margin: 0; padding: 0; }</style>
  </head>
  <body>
    <redoc spec-url="/openapi.yaml"></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
  </body>
</html>`

// Redoc serves a Redoc UI page that loads /openapi.yaml.
func Redoc(c fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(redocHTML)
}

// OpenAPISpec serves the raw OpenAPI YAML document.
func OpenAPISpec(c fiber.Ctx) error {
	c.Set("Content-Type", "application/yaml; charset=utf-8")
	return c.SendFile("docs/openapi.yaml")
}
