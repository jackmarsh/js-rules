package main

import (
	"encoding/json"
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/thought-machine/go-flags"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var opts = struct {
	Usage string

	Out         string   `short:"o" long:"out"`
	EntryPoints []string `short:"e" long:"entry_point"`

	Link struct {
		Modules map[string]string `short:"m" long:"module" description:"Module mapping"`
	} `command:"link" alias:"c" description:"Compile the entry_points, redirecting requires for the provided modules"`
	Compile struct {
		PackageJSON string   `short:"p" long:"package_json"`
		External    []string `long:"external"`
	} `command:"compile" alias:"c" description:"Compile the entry_points, redirecting requires for the provided modules"`
}{
	Usage: `
esbuild provides a wrapper around esbuild, using plugins to perform a more traditional "compile" and "link" workflow 
around bundling. 
`,
}

var wd, wdErr = os.Getwd()
var plugin = api.Plugin{
	Name: "please",
	Setup: func(build api.PluginBuild) {
		build.OnResolve(api.OnResolveOptions{Filter: `.*`},
			func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				if path, ok := opts.Link.Modules[args.Path]; ok {
					return api.OnResolveResult{
						Path:      path,
						Namespace: "please",
					}, nil
				}
				return api.OnResolveResult{}, nil
			})
		build.OnLoad(api.OnLoadOptions{Namespace: "please", Filter: `.*`}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
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
	},
}

func findEntryPointFromPkgJSON() string {
	data, err := os.ReadFile(opts.Compile.PackageJSON)
	if err != nil {
		log.Fatalf("failed to read %v: %v", opts.Compile.PackageJSON, err)
	}
	pkgJSON := struct {
		Main  string `json:"main"`
		Types string `json:"types"`
		Module string `json:"module"`
	}{}

	if err := json.Unmarshal(data, &pkgJSON); err != nil {
		log.Fatalf("failed to parse %v: %v", opts.Compile.PackageJSON, err)
	}

	// Check for entry points in order of priority: `types`, `main`, `module`, fallback to `index.js`
	var entryPoint string
	dir := filepath.Dir(opts.Compile.PackageJSON)
	if pkgJSON.Types != "" {
		entryPoint = pkgJSON.Types
	} else if pkgJSON.Main != "" {
		entryPoint = pkgJSON.Main
	} else if pkgJSON.Module != "" {
		entryPoint = pkgJSON.Module
	} else {
		entryPoint = "index.js"
	}
	fullPath := filepath.Join(dir, entryPoint)
	if _, err := os.Lstat(fullPath); err != nil {
		log.Printf("Warning: %s not found, falling back to index.js", entryPoint)
		return filepath.Join(dir, "index.js")
	}
	return fullPath
}

func main() {
	p := flags.NewParser(&opts, flags.Default)

	_, err := p.Parse()
	if err != nil {
		os.Exit(1)
	}

	if wdErr != nil {
		panic(wdErr)
	}

	buildOpts := api.BuildOptions{
		EntryPoints: opts.EntryPoints,
		Outfile:     opts.Out,
		Bundle:      true,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
		Platform:    api.PlatformNode,
	}

	log.Printf(p.Command.Name)
	if p.Active.Name == "link" {
		buildOpts.Plugins = []api.Plugin{plugin}
		buildOpts.Format = api.FormatESModule
	} else {
		if len(opts.EntryPoints) == 0 && opts.Compile.PackageJSON != "" {
			buildOpts.EntryPoints = []string{findEntryPointFromPkgJSON()}
		}
		buildOpts.External = opts.Compile.External
		buildOpts.Format = api.FormatCommonJS
	}

	log.Printf("external: %v", buildOpts.External)
	result := api.Build(buildOpts)
	if len(result.Errors) > 0 {
		os.Exit(1)
	}

}
