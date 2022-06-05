/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"net/http"

	"go.bytebuilders.dev/cli"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/unrolled/render"
	"go.wandrs.dev/binding"
	httpw "go.wandrs.dev/http"
	"k8s.io/klog/v2"
)

type FormData struct {
	Name    string `form:"name" binding:"Required" json:"name"`
	Email   string `form:"email" binding:"Required;Email" json:"email"`
	CC      string `form:"cc" json:"cc"`
	Product string `form:"product" binding:"Required" json:"product"` // This is now called plan in a parsed LicenseInfo
	Cluster string `form:"cluster" binding:"Required" json:"cluster"`
	Tos     string `form:"tos" binding:"Required" json:"tos"`
	Token   string `form:"token" json:"token"`
}

func NewCmdInstaller() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "install",
		Short:             `Install ACE`,
		Long:              `Install ACE`,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInstaller()
		},
	}

	return cmd
}

func runInstaller() error {
	m := chi.NewRouter()
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger) // middlewares.NewLogger()
	m.Use(middleware.Recoverer)
	m.Use(binding.Injector(render.New(render.Options{
		Directory: "templates",
		FileSystem: &render.EmbedFileSystem{
			FS: cli.FS(),
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
	return http.ListenAndServe(":4000", m)
}

func DisplayFormData(ctx httpw.ResponseWriter) {
	ctx.HTML(http.StatusOK, "installer", map[string]string{
		"Product": "console-enterprise",
	})
}

func CaptureFormData(ctx httpw.ResponseWriter, params FormData) {
	ctx.JSON(http.StatusOK, params)
}
