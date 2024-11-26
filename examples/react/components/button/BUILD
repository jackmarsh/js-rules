subinclude("//build_defs:js")

js_library(
    name = "button",
    srcs = [
        "Button.jsx",
    ],
    entry_point = "index.jsx",
    deps = [
        "//third_party/js:react",
        "//third_party/js:react-dom",
        ":styles",
    ],
    visibility = ["//examples/react/..."],
)

css_module(
    name = "styles",
    entry_point = "index.js",
    srcs = [
        "Button.module.css",
    ],
    visibility = ["//examples/react/..."],
)
