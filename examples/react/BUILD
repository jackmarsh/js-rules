subinclude("//build_defs:js")

js_binary(
    name = "react",
    srcs = [
        "index.jsx",
        "index.css",
    ],
    entry_point = "index.jsx",
    deps = [
        "//examples/react/App:app",
        "//third_party/js:react",
        "//third_party/js:react-dom",
    ],
)

http_proxy(
    name="http_proxy",
    proxy="/api",
    host="0.0.0.0",
    protocol="http",
    port=8081,
    path_rewrite={
        "^/api": "",
    },
    headers={
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
        "Access-Control-Allow-Headers": "X-Requested-With, content-type, Authorization",
        "Cache-Control": "no-store, no-cache, must-revalidate",
    },
)

dev_server(
    name="dev_server",
    entry_point = ":react",
    static_files = [
        "index.html",
    ],
    port = 8000,
    proxy = ":http_proxy",
)
