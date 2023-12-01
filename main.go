package main

import (
	"final_allegro/database"
	"final_allegro/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDB()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./src")

	app.Get("/contoh", func(c *fiber.Ctx) error {
		return c.Render("contoh", nil)
	})

	// Login
	app.Get("/", handlers.LoginPage)
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/logout", handlers.Logout)

	// Dashboard
	app.Get("/dashboard/:selector?", handlers.DashboardPage)
	app.Get("/render/:selector", handlers.Render)
	app.Delete("/delete-project", handlers.DeleteProject)

	// Form
	app.Get("/form", handlers.FormPage)
	app.Post("/form", handlers.FormPost)
	app.Get("/notif", handlers.Notification)

	// Asset Container
	app.Get("/asset-container", handlers.AddProfilePage)
	app.Post("/add-profile", handlers.PostProfile)

	// Risk
	app.Get("/risk", handlers.AddRiskPage)
	app.Post("/risk", handlers.AddRiskPost)
	app.Post("/score/:area/:projectId", handlers.HandleScore)

	// end

	log.Fatal(app.Listen(":3000"))
}

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// type FieldData struct {
// 	FieldName string
// 	Values    []Value
// }

// type Value struct {
// 	Checked bool
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.valueuest) {
// 		if r.Method == http.MethodPost {
// 			// Mendapatkan nilai dari form
// 			field := r.FormValue("field")
// 			valueStr := r.FormValue("value")

// 			// Mengonversi nilai dari string ke integer
// 			value, err := strconv.Atoi(valueStr)
// 			if err != nil || value < 1 || value > 5 {
// 				http.Error(w, "Invalid value", http.StatusBadRequest)
// 				return
// 			}

// 			// Logika backend untuk menentukan nilai checked
// 			fields := []FieldData{
// 				{FieldName: "Reputation and Customer Confidence", Values: []Value{{Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}}},
// 				{FieldName: "Financial", Values: []Value{{Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}}},
// 				{FieldName: "Productivity", Values: []Value{{Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}}},
// 				{FieldName: "Safety and Health", Values: []Value{{Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}}},
// 				{FieldName: "Fines and Legal Penalties", Values: []Value{{Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}, {Checked: false}}},
// 			}

// 			// Menandai nilai sesuai dengan input pengguna
// 			for i := range fields {
// 				if fields[i].FieldName == field {
// 					fields[i].Values[value-1].Checked = true
// 					break
// 				}
// 			}

// 			// Parsing template
// 			tmpl, err := template.New("index").ParseFiles("path/to/your/template.html")
// 			if err != nil {
// 				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 				return
// 			}

// 			// Menjalankan template dengan data
// 			err = tmpl.Execute(w, map[string]interface{}{"Data": fields})
// 			if err != nil {
// 				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 				return
// 			}
// 		} else {
// 			// Menampilkan form untuk input
// 			// ...

// 		}
// 	})

// 	http.ListenAndServe(":8080", nil)
// }
