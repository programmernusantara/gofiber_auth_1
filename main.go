package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Route untuk halaman utama
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Selamat datang")
	})

	// Middleware basic authentication
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"wildan": "190205",
			"rani":   "190205",
		},
		// Menyimpan nama pengguna di Lokal.
		ContextUsername: "user",
	}))

	// Route untuk halaman login
	app.Get("/login", func(c *fiber.Ctx) error {
		// Mendapatkan informasi pengguna yang berhasil login dari konteks
		user, _ := c.Locals("user").(string)

		// Mengembalikan pesan selamat dengan nama pengguna
		return c.JSON("Selamat " + user + ", Anda berhasil login")
	})

	// Menjalankan aplikasi pada port 8080
	app.Listen(":8080")
}
