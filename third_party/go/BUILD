package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "gotool",
    strip_srcs = True,
    version = "1.21.3",
)

go_mod_download(
    name = "esbuild_download",
    module = "github.com/evanw/esbuild",
    version = "v0.24.0",
)

go_module(
    name = "esbuild",
    binary = True,
    download = ":esbuild_download",
    install = ["cmd/esbuild"],
    module = "github.com/evanw/esbuild",
    deps = [":x_sys"],
)

go_module(
    name = "esbuild_api",
    download = ":esbuild_download",
    install = ["pkg/api", "internal/..."],
    module = "github.com/evanw/esbuild",
    deps = [":x_sys"],
)

go_module(
    name = "x_sys",
    install = ["..."],
    module = "golang.org/x/sys",
    version = "765f4ea38db36397e827c4153018aa272eed7835",
)

go_module(
    name = "go-flags",
    module = "github.com/thought-machine/go-flags",
    version = "v1.5.0",
    deps = [":x_sys"],
)

go_module(
    name = "x_sync",
    install = ["..."],
    module = "golang.org/x/sync",
    version = "v0.6.0",
)

go_module(
    name = "semver",
    module = "github.com/Masterminds/semver/v3",
    version = "v3.3.1",
)
