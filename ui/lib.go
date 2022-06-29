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

package ui

import (
	"embed"
	iofs "io/fs"
)

//go:embed dist/*.html dist/*/*.css dist/*/*.eot dist/*/*.js dist/*/*.svg dist/*/*.ttf dist/*/*.woff dist/*/*.woff2
var fs embed.FS

func FS() iofs.FS {
	fsys, err := iofs.Sub(fs, "dist")
	if err != nil {
		panic(err)
	}
	return fsys
}
