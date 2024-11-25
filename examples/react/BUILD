subinclude("//build_defs:js")

js_binary(
    name = "react",
    srcs = [
        "index.jsx",
    ],
    entry_point = "index.jsx",
    deps = [
        ":css",
        "//examples/react/App:app",
        "//third_party/js:react",
        "//third_party/js:react-dom",
    ],
)

css_file(
    name = "css",
    srcs = ["index.css"],
)

http_server(
    name="http_server",
    entry_point = ":react",
    static_files = [
        "index.html",
        "index.css",
    ],
)
