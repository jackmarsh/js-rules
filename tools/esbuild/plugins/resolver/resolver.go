package resolver

import (
	"log"

	"github.com/evanw/esbuild/pkg/api"
)

// Plugin creates an esbuild plugin for resolving and namespacing paths.
func Plugin(nodeModules, cssModules map[string]string) api.Plugin {
	return api.Plugin{
		Name: "resolver",
		Setup: func(build api.PluginBuild) {
			apiOptions := api.OnResolveOptions{
				Filter: `.*`,
			}
			build.OnResolve(apiOptions, onResolve(nodeModules, cssModules))
		},
	}
}

func onResolve(nodeModules, cssModules map[string]string) func(args api.OnResolveArgs) (api.OnResolveResult, error) {
	return func(args api.OnResolveArgs) (api.OnResolveResult, error) {
		log.Printf("%s on resolve: %s", args.Importer, args.Path)
		if path, ok := nodeModules[args.Path]; ok {
			log.Printf("module resolved: %s", path)
			return api.OnResolveResult{
				Path:      path,
				Namespace: "please",
			}, nil
		}
		if path, ok := cssModules[args.Path]; ok {
			log.Printf("css resolved: %s", path)
			return api.OnResolveResult{
				Path:      path,
				Namespace: "css",
			}, nil
		}
		return api.OnResolveResult{}, nil
	}
}
			
	
		
