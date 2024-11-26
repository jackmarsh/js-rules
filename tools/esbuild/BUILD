go_binary(
    name = "esbuild",
    srcs = ["esbuild.go"],
    deps = [
        "//third_party/go:esbuild_api",
        "//third_party/go:go-flags",
        "//tools/esbuild/plugins/css_loader",
        "//tools/esbuild/plugins/resolver",
        "//tools/esbuild/plugins/loader",
    ],
    visibility = ["PUBLIC"],
)
