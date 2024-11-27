# Dev Server Tool

## Introduction

The **Dev Server Tool** is a development server designed to serve JavaScript bundles efficiently, offering a fast and smooth development experience. Similar to the webpack dev server, it provides essential features like static file serving and proxying API requests but with improved performance and simplicity. Integrated into the [JS Rules](https://please.build) Please plugin, it seamlessly fits into projects using the Please build system without the need for additional installations or configurations.

---

## Features

- **High Performance**: Delivers JavaScript bundles quickly, significantly enhancing development speed.
- **Static File Serving**: Easily serves static assets such as HTML, CSS, and images.
- **Proxy Support**: Configures proxies to redirect API calls to backend servers.
- **Seamless Integration**: Works smoothly with the Please build system and JS Rules plugin.
- **Easy Configuration**: Utilizes simple build definitions to set up and run the server.

---

## How It Works

The Dev Server Tool operates by:

1. **Serving Static Files**: Hosts files from specified directories, allowing you to serve your bundled JavaScript and other assets.
2. **Proxying Requests**: Intercepts API calls and proxies them to a backend server, which is useful for development environments where the frontend and backend are separate.
3. **Fallback Mechanism**: If a requested resource is not found, it falls back to serving a default `index.html`, supporting single-page applications (SPAs).

---

## Usage

The preferred way to use the Dev Server Tool is through the `dev_server` build definition provided by the JS Rules plugin. This method integrates the server directly into your Please build files without the need for separate installation or command-line execution.

### Setting Up the Dev Server

Add the `dev_server` build definition to your `BUILD` file:

```python
dev_server(
    name = "dev_server",
    entry_point = ":main",
    static_files = [
        "index.html",
    ],
    port = 8000,
)
```

- **Parameters**:
  - `name`: Name of the build target.
  - `entry_point`: The main JavaScript binary (`js_binary`) build target to serve.
  - `static_files`: List of static files (e.g., `['index.html']`) to be served.
  - `port`: Port number for the HTTP server (default: `8080`).

### Starting the Dev Server

Use the Please command-line interface to run the dev server:

```bash
$ plz run //path/to/your:dev_server
[js-rules dev-server] Server is running at the following addresses:
[js-rules dev-server] 	Loopback: http://localhost:8000
[js-rules dev-server] 	On Your Network (IPv4): http://<ip-address>:8000
[js-rules dev-server] Static content being served from 'plz-out/bin/path/to/your/dist' directory
[js-rules dev-server] 404s will fallback to 'plz-out/gen/path/to/your/index.html'
[js-rules dev-server] Compiled successfully in 256.897µs
```

This command builds and runs the development server, serving your application on the specified port.

---

## Setting Up a Proxy

If your application requires proxying API requests to a backend server during development, you can use the `http_proxy` build definition.

### Define the Proxy Configuration

Add the `http_proxy` build definition to your `BUILD` file:

```python
http_proxy(
    name = "http_proxy",
    proxy = "/api",
    host = "localhost",
    protocol = "http",
    port = 8081,
    headers = {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
        "Access-Control-Allow-Headers": "Content-Type, Authorization",
        # ... other headers as needed
    },
    path_rewrite = {
        "^/api": "",
    },
)
```

- **Parameters**:
  - `name`: Name of the build target.
  - `proxy`: Proxy path prefix (e.g., `"/api"`).
  - `host`: Host address for the proxy target (e.g., `"localhost"`).
  - `protocol`: Protocol to use (`"http"` or `"https"`, default: `"http"`).
  - `port`: Port number of the proxy target (default: `8081`).
  - `headers`: Dictionary of headers to include in the proxy requests.
  - `path_rewrite`: Dictionary of regex patterns and replacements for rewriting request paths.

### Update the Dev Server to Use the Proxy

Modify the `dev_server` build definition to include the proxy:

```python
dev_server(
    name = "dev_server",
    entry_point = ":react",
    static_files = [
        "index.html",
    ],
    port = 8000,
    proxy = ":http_proxy",
)
```

### Starting the Dev Server with Proxy

Run the server using the same command:

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

Now, API requests matching the proxy configuration will be forwarded to the specified backend server.

---

## Build Definitions Documentation

### `dev_server` Build Definition

Defines a development server target that serves JavaScript bundles and static files.

#### Signature

```python
def dev_server(
    name: str,
    entry_point: str,
    static_files: list = [],
    port: int = 8080,
    proxy: str = None
)
```

#### Parameters

- `name`: **(str)** Name of the build target.
- `entry_point`: **(str)** The main JavaScript file or build target to serve.
- `static_files`: **(list)** List of static files to be served.
- `port`: **(int)** Port number for the HTTP server.
- `proxy`: **(str)** Optional build target for proxy configuration (e.g., `":http_proxy"`).

#### Description

- Creates a `filegroup` for static files.
- Constructs the command to run the development server with appropriate arguments.
- Uses `sh_cmd` to define a build target that runs the server.
- Integrates with the Please build system and JS Rules plugin.

---

### `http_proxy` Build Definition

Defines a proxy configuration target that generates a JSON file for proxy settings.

#### Signature

```python
def http_proxy(
    name: str,
    proxy: str,
    host: str,
    protocol: str = "http",
    port: int = 8081,
    headers: dict = {},
    path_rewrite: dict = {}
)
```

#### Parameters

- `name`: **(str)** Name of the build target.
- `proxy`: **(str)** Proxy path prefix to match (e.g., `"/api"`).
- `host`: **(str)** Host address for the backend server.
- `protocol`: **(str)** Protocol to use (`"http"` or `"https"`, default: `"http"`).
- `port`: **(int)** Port number on which the backend server is running.
- `headers`: **(dict)** Headers to include in the proxied requests and responses.
- `path_rewrite`: **(dict)** Rules for rewriting request paths before proxying.

#### Description

- Uses `genrule` to generate a `proxy.json` file with the provided configuration.
- The generated `proxy.json` is used by the dev server to set up proxying.

---

## Proxy Configuration: Detailed Description

When using the `http_proxy` build definition, the proxy configuration is generated automatically. Below is a detailed explanation of the fields within the `ProxyConfig`:

### `proxy` (string)

- **Purpose**: Defines the path prefix that the server should proxy to the backend.
- **Usage**: Any request path that matches this prefix will be forwarded to the backend server.
- **Example**: If set to `"/api"`, all requests starting with `/api` will be proxied.

### `host` (string)

- **Purpose**: Specifies the hostname or IP address of the backend server to which the requests should be proxied.
- **Usage**: Set to the address where your backend server is running.
- **Example**: `"localhost"` or `"192.168.1.100"`.

### `protocol` (string)

- **Purpose**: Indicates the protocol to use when connecting to the backend server.
- **Allowed Values**: `"http"` or `"https"`.
- **Example**: `"http"` for unencrypted connections.

### `port` (integer)

- **Purpose**: Defines the port number on which the backend server is listening.
- **Usage**: Must match the port your backend server is using.
- **Example**: `8081`.

### `headers` (dictionary)

- **Purpose**: Specifies additional HTTP headers to add to both proxied requests and responses.
- **Usage**: Useful for setting CORS headers or custom authentication tokens.
- **Example**:

  ```python
  headers = {
      "Access-Control-Allow-Origin": "*",
      "Authorization": "Bearer <token>"
  }
  ```

### `path_rewrite` (dictionary)

- **Purpose**: Defines rules for modifying the request path before proxying.
- **Usage**: Each key is a regex pattern to match against the request path, and the value is the replacement string.
- **Example**:

  ```python
  path_rewrite = {
      "^/api": ""
  }
  ```

  This removes the `/api` prefix from the request path.

---

## Advantages Over Traditional Development Servers

- **Performance**: Significantly faster startup and response times compared to traditional tools like webpack dev server.
- **Integration**: Seamlessly integrates with the Please build system and JS Rules plugin.
- **Simplicity**: Minimal configuration required to get started.
- **Flexibility**: Easy to set up proxies and serve static files without additional tools.

---

## Limitations

- **Single Proxy Configuration**: Currently supports a single proxy configuration. Multiple proxies are not supported out of the box.
- **No SSL Termination**: SSL termination is not handled by the tool. For HTTPS support, consider using a reverse proxy like Nginx.
- **No Hot Module Replacement**: Unlike some development servers, this tool does not support hot module replacement (HMR).

---

## Contributing

Contributions to the Dev Server Tool are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on the project's repository.

---

## License

This tool is licensed under the MIT License.

---

If you have any questions or need further assistance, feel free to consult the tool's source code or reach out to the maintainers.
