package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/evanw/esbuild/pkg/api"
	"github.com/thought-machine/go-flags"

	"tools/esbuild/plugins/resolver"
	"tools/esbuild/plugins/loader"
	"tools/esbuild/plugins/css_loader"
)

var opts = struct {
	Usage string

	Out         string   `short:"o" long:"out"`
	OutDir      string   `short:"d" long:"out-dir"`
	EntryPoints []string `short:"e" long:"entry_point"`

	Link struct {
		Modules map[string]string `short:"m" long:"module" description:"Module mapping"`
		CSS map[string]string `short:"s" long:"css" description:"CSS Mapping"` 
	} `command:"link" alias:"l" description:"Compile the entry_points, redirecting requires for the provided modules"`
	Compile struct {
		PackageJSON string   `short:"p" long:"package_json"`
		External    []string `long:"external"`
		Binary      bool     `long:"binary" description:"Indicates if the package is a binary module"`
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
		build.OnResolve(api.OnResolveOptions{Filter: `.*`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
			log.Printf("%s on resolve: %s", args.Importer, args.Path)
			if path, ok := opts.Link.Modules[args.Path]; ok {
				log.Printf("module resolved: %s", path)
				return api.OnResolveResult{
					Path:      path,
					Namespace: "please",
				}, nil
			} else {
				log.Printf("module not resolved")
			}
			if path, ok := opts.Link.CSS[args.Path]; ok {
				log.Printf("css resolved: %s", path)
				return api.OnResolveResult{
					Path:      path,
					Namespace: "css",
				}, nil
			} else {
				log.Printf("css not resolved")
			}
				
			return api.OnResolveResult{}, nil
		})
		build.OnLoad(api.OnLoadOptions{Namespace: "please", Filter: `.*`}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
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
	},
}

func findEntryPointFromPkgJSON() string {
	data, err := os.ReadFile(opts.Compile.PackageJSON)
	if err != nil {
		log.Fatalf("failed to read %v: %v", opts.Compile.PackageJSON, err)
	}

	type PkgJSON struct {
		Bin     interface{} `json:"bin"`
		Main    string      `json:"main"`
		Module  string      `json:"module"`
		Types   string      `json:"types"`
	}

	var pkgJSON PkgJSON
	if err := json.Unmarshal(data, &pkgJSON); err != nil {
		log.Fatalf("failed to parse %v: %v", opts.Compile.PackageJSON, err)
	}

	dir := filepath.Dir(opts.Compile.PackageJSON)
	var entryPoint string

	// Handle binary modules if specified
	if opts.Compile.Binary {
		if pkgJSON.Bin != nil {
			switch bin := pkgJSON.Bin.(type) {
			case string:
				entryPoint = bin
			case map[string]interface{}:
				// If "bin" is an object, pick the first entry
				for _, v := range bin {
					if s, ok := v.(string); ok {
						entryPoint = s
						break
					}
				}
			default:
				log.Fatalf("unexpected type for 'bin' field in %v", opts.Compile.PackageJSON)
			}
		} else {
			log.Fatalf("binary module specified but 'bin' field is missing in %v", opts.Compile.PackageJSON)
		}
	} else {
		// Determine non-binary entry point based on priority
		if pkgJSON.Main != "" {
			entryPoint = pkgJSON.Main
		} else if pkgJSON.Module != "" {
			entryPoint = pkgJSON.Module
		} else {
			entryPoint = "index.js" // Default behavior
		}

		// If entryPoint doesn't have a .js extension, handle directory case
		fullPath := filepath.Join(dir, entryPoint)
		if !strings.HasSuffix(entryPoint, ".js") {
			// Check if entryPoint is a directory
			info, err := os.Stat(fullPath)
			if err == nil && info.IsDir() {
				// Use index.js within the directory
				fullPath = filepath.Join(fullPath, "index.js")
			} else {
				// Otherwise, assume it's a file with a .js extension
				fullPath += ".js"
			}
		}

		// Validate that the resolved file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// Fallback: Walk the directory tree to find the first index.js
			log.Printf("entry point '%s' not found; walking file tree to locate index.js", fullPath)
			fullPath = findFirstIndexJS(dir)
		}

		return fullPath
	}

	// Validate binary entry point
	fullPath := filepath.Join(dir, entryPoint)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalf("binary entry point '%s' not found in package directory %s", entryPoint, dir)
	}

	return fullPath
}

// findFirstIndexJS walks the file tree to find the first index.js file
func findFirstIndexJS(dir string) string {
	var firstIndexJS string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == "index.js" {
			firstIndexJS = path
			return filepath.SkipDir // Stop walking once we find the first match
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error walking the file tree: %v", err)
	}
	if firstIndexJS == "" {
		log.Fatalf("no index.js file found in package directory %s", dir)
	}
	return firstIndexJS
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
		Outdir:      opts.OutDir,
		Bundle:      true,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
		Platform:    api.PlatformNode,
		Loader: map[string]api.Loader{
			".js":   api.LoaderJS,
			".jsx":  api.LoaderJSX,
			".ts":   api.LoaderTS,
			".tsx":  api.LoaderTSX,
			".json": api.LoaderJSON,
			// ".css": api.LoaderGlobalCSS,
			// ".module.css": api.LoaderLocalCSS,
			// ".d.ts":  api.LoaderNone, // ??
		},
		JSXFactory: "React.createElement",
		JSXFragment: "React.Fragment",
		Target:      api.ESNext,
		Define: map[string]string{
			"process.env.NODE_ENV": "\"production\"",
		},
		Sourcemap: api.SourceMapLinked,
	}

	log.Printf(p.Command.Name)
	if p.Active.Name == "link" {
		buildOpts.Plugins = []api.Plugin{
			resolver.Plugin(opts.Link.Modules, opts.Link.CSS),
			loader.Plugin(),
			cssloader.Plugin(),
		}
		buildOpts.Format = api.FormatESModule
	} else {
		if len(opts.EntryPoints) == 0 && opts.Compile.PackageJSON != "" {
			buildOpts.EntryPoints = []string{findEntryPointFromPkgJSON()}
		}
		buildOpts.External = opts.Compile.External
		if opts.Compile.Binary {
			buildOpts.Format = api.FormatESModule
		}
	}

	log.Printf("external: %v", buildOpts.External)
	result := api.Build(buildOpts)
	if len(result.Errors) > 0 {
		os.Exit(1)
	}
	
}
