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

JS Rules comes with a dev server which is 10-100x faster than Webpack 5. To run the dev server simply add `http_server` to your `BUILD` file like so:
```
http_server(
    name="http_server",
    entry_point = ":main",
    static_files = [
        "index.html",
    ],
    port = 8000,
)
```

You can setup a proxy for the dev server with `http_proxy`:
```
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

http_server(
    name="http_server",
    entry_point = ":react",
    static_files = [
        "index.html",
    ],
    port = 8000,
    proxy = ":http_proxy",
)
```
Then run:
```plz run //path/to/your:http_server```

## Configuration
Plugins are configured through the Plugin section like so:
```ini
[Plugin "rust"]
SomeConfig = some-value
```
