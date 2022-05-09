package main

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/render"
	"go.wandrs.dev/binding"
	httpw "go.wandrs.dev/http"
	"gomodules.xyz/logs"
	"k8s.io/klog/v2"
)

//go:embed templates/*.html
var fs embed.FS

type FormData struct {
	Name    string `form:"name" binding:"Required" json:"name"`
	Email   string `form:"email" binding:"Required;Email" json:"email"`
	CC      string `form:"cc" json:"cc"`
	Product string `form:"product" binding:"Required" json:"product"` // This is now called plan in a parsed LicenseInfo
	Cluster string `form:"cluster" binding:"Required" json:"cluster"`
	Tos     string `form:"tos" binding:"Required" json:"tos"`
	Token   string `form:"token" json:"token"`
}

func main() {
	logs.Init(nil, true)
	defer logs.FlushLogs()

	m := chi.NewRouter()
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger) // middlewares.NewLogger()
	m.Use(middleware.Recoverer)
	m.Use(binding.Injector(render.New(render.Options{
		Directory: "templates",
		FileSystem: &render.EmbedFileSystem{
			FS: fs,
		},
		Extensions: []string{".html", ".tmpl"},
	})))

	// PUBLIC
	m.Route("/html", func(m chi.Router) {
		m.Get("/", binding.HandlerFunc(DisplayFormData))
		m.With(binding.Bind(FormData{})).Post("/", binding.HandlerFunc(CaptureFormData))
	})

	m.Get("/", binding.HandlerFunc(func() string {
		return "Hello world!"
	}))
	klog.Infoln()
	klog.Infoln("Listening on http://localhost:4000/html")
	if err := http.ListenAndServe(":4000", m); err != nil {
		klog.Fatalln(err)
	}
}

func DisplayFormData(ctx httpw.ResponseWriter) {
	ctx.HTML(http.StatusOK, "installer", map[string]string{
		"Product": "console-enterprise",
	})
}

func CaptureFormData(ctx httpw.ResponseWriter, params FormData) {
	ctx.JSON(http.StatusOK, params)
}
