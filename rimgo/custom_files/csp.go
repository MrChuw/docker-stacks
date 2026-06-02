package utils

import (
//      "log"
        "os"
        "strings"

        "github.com/gofiber/fiber/v2"
)


func BaseCSPSlice() []string {
	return []string{
		"default-src 'none'",
		"base-uri 'none'",
		"frame-ancestors 'none'",
		"form-action 'self'",
		"style-src 'self'",
		"img-src 'self'",
		"media-src 'self'",
		"manifest-src 'self'",
	}
}

func CSPWithAnalytics() string {
	directives := BaseCSPSlice()
	umami := os.Getenv("UMAMI_HOST")

	if umami != "" {
		directives = append(directives, "script-src 'self' "+umami)
		directives = append(directives, "connect-src 'self' "+umami)
	} else {
		directives = append(directives, "script-src 'self'")
	}

	return strings.Join(directives, "; ")
}

func SecurityHeaders() fiber.Handler {
        return func(c *fiber.Ctx) error {
                err := c.Next()
                if err != nil {
                        return err
                }

                c.Set("Content-Security-Policy", CSPWithAnalytics())
                return nil
        }
}