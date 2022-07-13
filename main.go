package main

import (
  "fmt"
  "net/http"

  "github.com/go-chi/chi/v5"
  "github.com/volo-h/miniblog/controllers"
  "github.com/volo-h/miniblog/templates"
  "github.com/volo-h/miniblog/views"
  "github.com/volo-h/miniblog/models"
)

func main() {
  r := chi.NewRouter()

  r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
    templates.FS,
    "home.gohtml", "tailwind.gohtml",
  ))))

  r.Get("/contact", controllers.StaticHandler(views.Must(
    views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

  r.Get("/faq", controllers.FAQ(
    views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))


  cfg := models.DefaultPostgresConfig()
  db, err := models.Open(cfg)

  if err != nil {
    panic(err)
  }
  defer db.Close()

  userService := models.UserService{
    DB: db,
  }

  usersC := controllers.Users{
    UserService: &userService,
  }
  usersC.Templates.New = views.Must(views.ParseFS(
    templates.FS, "signup.gohtml", "tailwind.gohtml"))
  usersC.Templates.SignIn = views.Must(views.ParseFS(
    templates.FS, "signin.gohtml", "tailwind.gohtml"))
  r.Get("/signup", usersC.New)
  r.Post("/signup", usersC.Create)
  r.Get("/signin", usersC.SignIn)
  r.Post("/signin", usersC.ProcessSignIn)
  r.Get("/users/me", usersC.CurrentUser)

  r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Page not found", http.StatusNotFound)
  })
  fmt.Println("Starting the server on :3000...")
  http.ListenAndServe(":3000", r)
}
