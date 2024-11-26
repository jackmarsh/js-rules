package loader

import (
	"log"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"

	"github.com/evanw/esbuild/pkg/api"
)

// Plugin creates an esbuild plugin for loading Please-namespace files.
func Plugin() api.Plugin {
	return api.Plugin{
		Name: "loader",
		Setup: setupFn,
	}
}

var wd, wdErr = os.Getwd()
func setupFn(build api.PluginBuild) {
	apiOptions := api.OnLoadOptions{
		Namespace: "please",
		Filter: `.*`,
	}
	build.OnLoad(apiOptions, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
		log.Printf("on load: %s", args.Path)
		path := filepath.Join(wd, args.Path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load %v: %v\n", args.Path, err)
			os.Exit(1)
		}

		contents := string(data)
		return api.OnLoadResult{
			Contents: &contents,
		}, nil
	})
}
