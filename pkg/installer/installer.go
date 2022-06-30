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

package installer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"go.bytebuilders.dev/cli/ui"
	api "go.bytebuilders.dev/installer/apis/installer/v1alpha1"
	"go.bytebuilders.dev/installer/schema"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/unrolled/render"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

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
	rn := render.New()
	m := chi.NewRouter()
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)

	m.Route("/apis", func(m chi.Router) {
		m.Get("/schema.json", func(w http.ResponseWriter, r *http.Request) {
			data, err := schema.FS().ReadFile("ace-options/values.openapiv3_schema.yaml")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data, err = yaml.YAMLToJSON(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		})
		m.Get("/options.json", func(w http.ResponseWriter, r *http.Request) {
			in := NewSampleOptions()
			//if data, err := yaml.Marshal(*in); err != nil {
			//	panic(err)
			//} else {
			//	_ = ioutil.WriteFile(filepath.Join(confDir(), "options.yaml"), data, 0o644)
			//}
			rn.JSON(w, http.StatusOK, in)
			//data, err := schema.FS().ReadFile("ace-options/values.openapiv3_schema.yaml")
			//if err != nil {
			//	http.Error(w, err.Error(), http.StatusInternalServerError)
			//	return
			//}
			//data, err = yaml.YAMLToJSON(data)
			//if err != nil {
			//	http.Error(w, err.Error(), http.StatusInternalServerError)
			//	return
			//}
			//w.Header().Set("Content-Type", "application/json")
			//w.Write(data)
		})
		m.Post("/install", func(w http.ResponseWriter, r *http.Request) {
			var in api.AceOptionsSpec
			err := json.NewDecoder(r.Body).Decode(&in)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			md, err := generateInstaller(&in)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			rn.Text(w, http.StatusOK, md)
		})
	})

	m.Get("/*", http.FileServer(http.FS(ui.FS())).ServeHTTP)

	klog.Infoln()
	url := "http://localhost:4000"
	klog.Infof("Listening on %s", url)

	go func() {
		time.Sleep(100 * time.Millisecond)
		if err := open(url); err != nil {
			log.Println(err)
		}
	}()
	return http.ListenAndServe(":4000", m)
}

// https://gist.github.com/hyg/9c4afcd91fe24316cbf0
// https://stackoverflow.com/a/39324149
func open(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}
