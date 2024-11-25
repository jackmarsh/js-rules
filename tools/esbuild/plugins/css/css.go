package css

import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

func Plugin() api.Plugin {
	return api.Plugin{
		Name: "css-plugin",
		Setup: func(build api.PluginBuild) {
			build.OnLoad(api.OnLoadOptions{Filter: `\.css$`}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				contents := fmt.Sprintf(`import styles from "%s"; export default styles;`, args.Path)
				return api.OnLoadResult{
					Contents:   &contents,
					Loader:     api.LoaderCSS,
					ResolveDir: args.Path,
				}, nil
			})
		},
	}
}
