package utils

import (
//      "log"
        "os"
        "strings"

        "github.com/gofiber/fiber/v2"
)

func BaseCSP() string {
        return strings.Join([]string{
                "default-src 'none'",
                "base-uri 'none'",
                "frame-ancestors 'none'",
                "form-action 'self'",
                "style-src 'self'",
                "img-src 'self'",
                "manifest-src 'self'",
                "media-src 'self'",
                "block-all-mixed-content",
        }, "; ")
}

func CSPWithAnalytics() string {
        umami := os.Getenv("UMAMI_HOST")

        if umami == "" {
                return BaseCSP()
        }

        csp := BaseCSP() + "; " + strings.Join([]string{
                "script-src " + umami,
                "connect-src " + umami,
        }, "; ")

        return csp
}

func SecurityHeaders() fiber.Handler {
        return func(c *fiber.Ctx) error {
                err := c.Next()
                if err != nil {
                        return err
                }
                accept := c.Get("Accept")
                if !strings.Contains(accept, "text/html") {
                        return nil
                }
                c.Set("Content-Security-Policy", CSPWithAnalytics())
                return nil
        }
}