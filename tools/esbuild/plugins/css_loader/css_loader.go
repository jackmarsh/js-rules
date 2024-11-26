package cssloader

import (
	"fmt"

	"log"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
	
	"github.com/evanw/esbuild/pkg/api"
)

// Plugin creates an esbuild plugin for loading CSS-namespaced files.
func Plugin() api.Plugin {
	return api.Plugin{
		Name: "css-loader",
		Setup: setupFn,
	}
}

var wd, wdErr = os.Getwd()
func setupFn(build api.PluginBuild) {
	apiOptions := api.OnLoadOptions{
		Namespace: "css",
		Filter: `.*\.css`, //
	}
	build.OnLoad(apiOptions, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
		log.Printf("on load: %s", args.Path)
		path := filepath.Join(wd, args.Path)
		var loader api.Loader
		if strings.HasSuffix(args.Path, ".css") {
			loader = api.LoaderLocalCSS
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load %v: %v\n", args.Path, err)
			os.Exit(1)
		}

		contents := string(data)
		return api.OnLoadResult{
			Contents: &contents,
			Loader: loader,
		}, nil
	})
}
