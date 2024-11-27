# ESBuild Tool

## Introduction

The **ESBuild Tool** is a wrapper around the [esbuild](https://esbuild.github.io/) JavaScript bundler and minifier, tailored to work seamlessly with the [Please](https://please.build) build system and the JS Rules plugin. It provides a structured "compile" and "link" workflow, enabling efficient bundling of JavaScript and CSS code in a manner similar to traditional build processes. This tool is integrated under the hood and is used in `js_library`, `jsx_library`, `css_module` and `node_module` build definitions for compiling, and the `js_binary` for linking.

---

## Features

- **Structured Workflow**: Separates the build process into distinct "compile" and "link" phases.
- **Module Resolution**: Custom plugins handle module resolution for Node.js modules and CSS modules within the Please build system.
- **TypeScript Support**: Supports compiling TypeScript files out of the box.
- **JSX and React Support**: Handles JSX syntax with customizable factory and fragment options.
- **Efficient Bundling**: Utilizes esbuild's high performance for fast builds.
---

## How It Works

The ESBuild Tool operates in two primary phases:

1. **Compile Phase**:
   - **Purpose**: Prepares individual modules or libraries by compiling their source code into a format suitable for linking.
   - **Process**:
     - Compiles source files (e.g., JavaScript, TypeScript) into JavaScript.
     - Uses custom plugins to handle module resolution and asset loading.
     - Marks external dependencies to be resolved during the link phase.
   - **Used By**: `js_library`, `jsx_library`, `css_module`, and `node_module` build definitions.

2. **Link Phase**:
   - **Purpose**: Bundles compiled modules and their dependencies into a final output suitable for deployment or execution.
   - **Process**:
     - Resolves and includes all dependencies.
     - Bundles the code into a single or multiple output files.
     - Handles CSS imports and other assets.
   - **Used By**: `js_binary` build definition.

**Why Separate Compile and Link Phases?**

Separating the compile and link phases offers several benefits:

- **Modularity**: Allows individual modules to be compiled independently, promoting reuse, caching build artifacts and parallelism.
- **Efficiency**: Speeds up incremental builds by recompiling only changed modules.
- **Control**: Provides more granular control over the build process, enabling optimizations and customizations.
- **Scalability**: Better supports large projects with many dependencies by breaking down the build into manageable steps.

---

## Usage in Build Definitions

### `node_module` Build Definition

Compiles Node.js modules for use in your project.

#### Function Signature

```python
def node_module(
    name: str,
    scope: str = '',
    module: str = '',
    version: str = '',
    visibility: list = [],
    deps: list = [],
    hashes: list = None,
    binary: bool = False,
    entry_point: str = ''
)
```

#### Description

- **Downloads** the specified module from the NPM registry.
- **Compiles** the module during the compile phase.
- **Exports** the module for use in other build targets.

#### Example Usage

```python
node_module(
    name = "lodash",
    version = "4.17.21",
    visibility = ["PUBLIC"],
)
```

---

### `js_library` Build Definition

Compiles JavaScript source code into a library.

#### Function Signature

```python
def js_library(
    name: str,
    module_name: str = '',
    srcs: list = [],
    deps: list = [],
    visibility: list = [],
    entry_point: str = "index.js"
)
```

#### Description

- **Compiles** JavaScript or TypeScript source files.
- **Defines** an entry point for the library.
- **Specifies** dependencies that will be resolved during the link phase.

#### Example Usage

```python
js_library(
    name = "my_library",
    srcs = ["utils.js", "helpers.js"],
    deps = [":lodash"],
    visibility = ["PUBLIC"],
    entry_point = "main.js",
)
```

---

### `jsx_library` Build Definition

A syntactic sugar for `js_library` tailored for JSX files.

#### Function Signature

```python
def jsx_library(
    name: str,
    module_name: str = '',
    srcs: list = [],
    deps: list = [],
    visibility: list = [],
    entry_point: str = "index.jsx"
)
```

#### Description

- **Compiles** JSX source files.
- **Supports** React and other libraries that use JSX syntax.

#### Example Usage

```python
jsx_library(
    name = "my_component",
    srcs = ["Button.jsx", "Modal.jsx"],
    deps = [":react"],
    visibility = ["PUBLIC"],
    entry_point = "App.jsx",
)
```

---

### `css_module` Build Definition

Compiles CSS modules ensuring they are bundled into the `index.css`.

#### Function Signature

```python
def css_module(
    name: str,
    module_name: str = '',
    srcs: list = [],
    deps: list = [],
    visibility: list = [],
    entry_point: str = "index.css"
)
```

#### Description

- **Compiles** CSS files into CSS modules.
- **Enables** importing CSS in JavaScript files with module support.

#### Example Usage

```python
css_module(
    name = "styles",
    srcs = ["main.css", "theme.css"],
    visibility = ["PUBLIC"],
    entry_point = "main.css",
)
```

---

### `js_binary` Build Definition

Links compiled modules and dependencies into an executable JavaScript bundle.

#### Function Signature

```python
def js_binary(
    name: str,
    entry_point: str = "index.js",
    srcs: list = [],
    deps: list = [],
    visibility: list = [],
    out_dir: str = 'dist'
)
```

#### Description

- **Links** compiled libraries and modules into a final bundle.
- **Resolves** all dependencies specified in the `deps` parameter.
- **Outputs** the bundle to the specified `out_dir`.

#### Example Usage

```python
js_binary(
    name = "app_bundle",
    entry_point = "main.js",
    deps = [
        ":my_library",
        ":my_component",
        ":styles",
        ":lodash",
    ],
    visibility = ["PUBLIC"],
    out_dir = "dist",
)
```

---

## Under the Hood: ESBuild Tool Details

### Command-Line Interface

The ESBuild Tool provides a command-line interface with two primary commands:

1. **Compile Command**:
   - **Usage**: `compile`
   - **Description**: Compiles entry points, marking dependencies as external for later resolution.
   - **Options**:
     - `--entry_point`: Specifies the entry point file.
     - `--package_json`: Specifies the `package.json` file to determine the entry point.
     - `--external`: Marks modules as external.
     - `--binary`: Indicates if the package is a binary module.
     - `--out-dir`: Specifies the output directory.

2. **Link Command**:
   - **Usage**: `link`
   - **Description**: Compiles entry points, resolving and bundling specified modules.
   - **Options**:
     - `--entry_point`: Specifies the entry point file.
     - `--module`: Maps module names to paths for resolution.
     - `--css`: Maps CSS module names to paths.
     - `--out-dir`: Specifies the output directory.

---

### Custom Plugins

The ESBuild Tool uses custom plugins to extend esbuild's functionality:

#### Resolver Plugin

- **Purpose**: Handles module resolution for Node.js modules and CSS modules.
- **How It Works**:
  - Intercepts module resolution requests during the build process.
  - Checks if the requested module matches any provided mappings.
  - Resolves the module to the specified path and assigns it to a custom namespace.

#### Loader Plugin

- **Purpose**: Loads files from the custom "please" namespace.
- **How It Works**:
  - Intercepts file loading requests in the "please" namespace.
  - Reads the file contents and provides them to esbuild.

#### CSS Loader Plugin

- **Purpose**: Loads CSS modules from the "css" namespace.
- **How It Works**:
  - Intercepts file loading requests in the "css" namespace.
  - Reads the CSS file contents and provides them to esbuild with the appropriate loader.

---

### Why Use Custom Plugins?

- **Custom Module Resolution**: Allows precise control over how modules are resolved, enabling integration with the Please build system's dependency management.
- **Namespace Isolation**: Prevents conflicts by isolating modules in custom namespaces.
- **Asset Handling**: Provides specialized handling for different asset types like CSS modules.

---

## Limitations

- **No Hot Module Replacement**: The tool does not support hot module replacement (HMR) during development.
- **Single Entry Point for Binaries**: The `js_binary` build definition expects a single entry point.
- **Limited CSS Processing**: Advanced CSS processing features (e.g., Sass, Less) are not handled out of the box.
- **Plugin Complexity**: Custom plugins add complexity to the build process, which may require understanding the underlying implementation for advanced use cases.

---

If you have any questions or need further assistance, feel free to consult the tool's source code or reach out to the maintainers.
