# JS Rules

This repo provides JavaScript build rules for the [Please](https://please.build) build system.

We use [esbuild](https://esbuild.github.io/), an extremely fast bundler for the web.

## Basic usage
First add the plugin to your project. In `plugins/BUILD`:
```python
plugin_repo(
    name = "js",
    owner = "odonate",
    plugin = "js-rules",
    revision = "<Some git tag, commit, or other reference>",
)
```

Then add the plugin config to `.plzconfig`:
```ini
[Plugin "js"]
Target = //plugins:js_rules
```

You can then compile JavaScript libraries like so:
```python
subinclude("///js//build_defs:js")

js_library(
    name = "components",
    entry_point = "index.js",
    srcs = [
        "ComponentA.rs",
        "ComponentB.rs",
    ],
)
```

You can define third-party node modules using `node_module`:
```python
subinclude("///js//build_defs:js")

node_module(
    name="react",
    version="18.3.1",
    visibility = ["PUBLIC"],
)

node_module(
    name="react-dom",
    version="18.3.1",
    deps = [
        ":react",
        ":scheduler",
    ],
    visibility = ["PUBLIC"],
)

node_module(
    name="scheduler",
    version="0.23.2",
)
```

To compile a binary, you can use `js_binary`:
```python
subinclude("///js//build_defs:js")

js_binary(
    name = "main",
    entry_point = "index.js",
    srcs = [
        "index.css",
    ],
    deps = [
        ":components",
        "//third_party/js:<node_module_name>",
    ],
)
```

JS Rules comes with a dev server which is 10-100x faster than Webpack 5. To run the dev server simply add `dev_server` to your `BUILD` file like so:
```python
dev_server(
    name="dev_server",
    entry_point = ":main",
    static_files = [
        "index.html",
    ],
    port = 8000,
)
```

You can setup a proxy for the dev server with `http_proxy`:
```python
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
		...
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
```
Then run:
```bash
$ plz run //path/to/your:dev_server
[js-rules dev-server] [HPM] Proxy created: /api  ->  http://0.0.0.0:8081
[js-rules dev-server] [HPM] Proxy rewrite rule created: "^/api" ~> ""
[js-rules dev-server] Server is running at the following addresses:
[js-rules dev-server] 	Loopback: http://localhost:8000
[js-rules dev-server] 	On Your Network (IPv4): http://<ip-address>:8000
[js-rules dev-server] Static content being served from 'plz-out/bin/path/to/your/dist' directory
[js-rules dev-server] 404s will fallback to 'plz-out/gen/path/to/your/index.html'
[js-rules dev-server] Compiled successfully in 256.897µs
```

## Tools

### Node Modules

Generates a Please BUILD file for a specified Node.js module and version, including build targets for all its dependencies. This tool simplifies the integration of Node.js modules into the Please build system, ensuring seamless dependency management and build target configuration. More information can be found in the Node Modules tool's [README](tools/node_modules/README.md).

**Use Cases:**

- Automates the inclusion of Node.js modules in projects using Please.
- Reduces manual configuration by automatically resolving and incorporating all dependencies.
- Ideal for monorepo environments where consistent build definitions are crucial.

---

### ESBuild

Provides a Go-based wrapper around the esbuild API to efficiently bundle JavaScript and CSS code. Leveraging esbuild's high performance, this tool accelerates the bundling process of JavaScript applications. More information can be found in the ESBuild tool's [README](tools/esbuild/README.md).

**Use Cases:**

- Enhances build speed for JavaScript and CSS assets in Go-based projects.
- Integrates esbuild directly into Go applications for efficient asset processing.
- Suitable for projects requiring fast and reliable front-end asset compilation.

---

### Dev Server

An HTTP server designed for serving JavaScript bundles, offering functionality similar to the webpack dev server. It provides features like a static file server and proxying capabilities. More information can be found in the HTTP Server tool's [README](tools/dev_server/README.md).

**Use Cases:**

- Serves as a development server for local front-end asset delivery.
- Beneficial for developers seeking a more efficient alternative to webpack dev server.


## Configuration
Plugins are configured through the Plugin section like so:
```ini
[Plugin "js"]
SomeConfig = some-value
```
