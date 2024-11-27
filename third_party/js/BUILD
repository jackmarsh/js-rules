subinclude("//build_defs:js")

node_module(
    name="bignumber.js",
    version="9.1.2",
)

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

node_module(
    name = "react-router",
    module = "react-router",
    version = "7.0.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":cookie",
        ":react",
        ":set-cookie-parser",
        ":turbo-stream",
    ],
)

node_module(
    name = "cookie",
    module = "cookie",
    version = "1.0.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "set-cookie-parser",
    module = "set-cookie-parser",
    version = "2.7.1",
    visibility = [
        "PUBLIC",
    ],
    entry_point="lib/set-cookie.js",
)

node_module(
    name = "turbo-stream",
    module = "turbo-stream",
    version = "2.4.0",
    visibility = [
        "PUBLIC",
    ],
    entry_point="dist/turbo-stream.js",
)

node_module(
    name = "zeit_schemas",
    module = "schemas",
    scope = "@zeit",
    version = "2.36.0",
    visibility = [
        "PUBLIC",
    ],
    entry_point = "deployment/config-static.js",
)

node_module(
    name = "accepts",
    module = "accepts",
    version = "1.3.8",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":mime-types.2.1.35",
        ":negotiator",
    ],
)

node_module(
    name = "ajv",
    module = "ajv",
    version = "8.12.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":fast-deep-equal",
        ":json-schema-traverse",
        ":require-from-string",
        ":uri-js",
    ],
)

node_module(
    name = "ansi-align",
    module = "ansi-align",
    version = "3.0.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":string-width.4.2.3",
    ],
)

node_module(
    name = "ansi-regex.5.0.1",
    module = "ansi-regex",
    version = "5.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "ansi-regex",
    module = "ansi-regex",
    version = "6.1.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "ansi-styles.4.3.0",
    module = "ansi-styles",
    version = "4.3.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":color-convert",
    ],
)

node_module(
    name = "ansi-styles",
    module = "ansi-styles",
    version = "6.2.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "arch",
    module = "arch",
    version = "2.2.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "arg",
    module = "arg",
    version = "5.0.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "balanced-match",
    module = "balanced-match",
    version = "1.0.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "boxen",
    module = "boxen",
    version = "7.0.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":ansi-align",
        ":camelcase",
        ":chalk.5.3.0",
        ":cli-boxes",
        ":string-width",
        ":widest-line",
        ":wrap-ansi",
    ],
)

node_module(
    name = "brace-expansion",
    module = "brace-expansion",
    version = "1.1.11",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":balanced-match",
        ":concat-map",
    ],
)

node_module(
    name = "bytes",
    module = "bytes",
    version = "3.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "camelcase",
    module = "camelcase",
    version = "7.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "chalk-template",
    module = "chalk-template",
    version = "0.4.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":chalk.4.1.2",
    ],
)

node_module(
    name = "chalk.4.1.2",
    module = "chalk",
    version = "4.1.2",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":ansi-styles.4.3.0",
        ":supports-color",
    ],
)

node_module(
    name = "chalk",
    module = "chalk",
    version = "5.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "chalk.5.3.0",
    module = "chalk",
    version = "5.3.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "cli-boxes",
    module = "cli-boxes",
    version = "3.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "clipboardy",
    module = "clipboardy",
    version = "3.0.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":arch",
        ":execa",
        ":is-wsl",
    ],
)

node_module(
    name = "color-convert",
    module = "color-convert",
    version = "2.0.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":color-name",
    ],
)

node_module(
    name = "color-name",
    module = "color-name",
    version = "1.1.4",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "compressible",
    module = "compressible",
    version = "2.0.18",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":mime-db",
    ],
)

node_module(
    name = "compression",
    module = "compression",
    version = "1.7.4",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":accepts",
        ":bytes",
        ":compressible",
        ":debug",
        ":on-headers",
        ":safe-buffer",
        ":vary",
    ],
)

node_module(
    name = "concat-map",
    module = "concat-map",
    version = "0.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "content-disposition",
    module = "content-disposition",
    version = "0.5.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "cross-spawn",
    module = "cross-spawn",
    version = "7.0.6",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":path-key",
        ":shebang-command",
        ":which",
    ],
)

node_module(
    name = "debug",
    module = "debug",
    version = "2.6.9",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":ms",
    ],
)

node_module(
    name = "deep-extend",
    module = "deep-extend",
    version = "0.6.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "eastasianwidth",
    module = "eastasianwidth",
    version = "0.2.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "emoji-regex.8.0.0",
    module = "emoji-regex",
    version = "8.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "emoji-regex",
    module = "emoji-regex",
    version = "9.2.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "execa",
    module = "execa",
    version = "5.1.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":cross-spawn",
        ":get-stream",
        ":human-signals",
        ":is-stream",
        ":merge-stream",
        ":npm-run-path",
        ":onetime",
        ":signal-exit",
        ":strip-final-newline",
    ],
)

node_module(
    name = "fast-deep-equal",
    module = "fast-deep-equal",
    version = "3.1.3",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "get-stream",
    module = "get-stream",
    version = "6.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "has-flag",
    module = "has-flag",
    version = "4.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "human-signals",
    module = "human-signals",
    version = "2.1.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "ini",
    module = "ini",
    version = "1.3.8",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "is-docker",
    module = "is-docker",
    version = "2.2.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "is-fullwidth-code-point",
    module = "is-fullwidth-code-point",
    version = "3.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "is-port-reachable",
    module = "is-port-reachable",
    version = "4.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "is-stream",
    module = "is-stream",
    version = "2.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "is-wsl",
    module = "is-wsl",
    version = "2.2.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":is-docker",
    ],
)

node_module(
    name = "isexe",
    module = "isexe",
    version = "2.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "json-schema-traverse",
    module = "json-schema-traverse",
    version = "1.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "merge-stream",
    module = "merge-stream",
    version = "2.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "mime-db.1.33.0",
    module = "mime-db",
    version = "1.33.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "mime-db.1.52.0",
    module = "mime-db",
    version = "1.52.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "mime-db",
    module = "mime-db",
    version = "1.53.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "mime-types",
    module = "mime-types",
    version = "2.1.18",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":mime-db.1.33.0",
    ],
)

node_module(
    name = "mime-types.2.1.35",
    module = "mime-types",
    version = "2.1.35",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":mime-db.1.52.0",
    ],
)

node_module(
    name = "mimic-fn",
    module = "mimic-fn",
    version = "2.1.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "minimatch",
    module = "minimatch",
    version = "3.1.2",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":brace-expansion",
    ],
)

node_module(
    name = "minimist",
    module = "minimist",
    version = "1.2.8",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "ms",
    module = "ms",
    version = "2.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "negotiator",
    module = "negotiator",
    version = "0.6.3",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "npm-run-path",
    module = "npm-run-path",
    version = "4.0.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":path-key",
    ],
)

node_module(
    name = "on-headers",
    module = "on-headers",
    version = "1.0.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "onetime",
    module = "onetime",
    version = "5.1.2",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":mimic-fn",
    ],
)

node_module(
    name = "path-is-inside",
    module = "path-is-inside",
    version = "1.0.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "path-key",
    module = "path-key",
    version = "3.1.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "path-to-regexp",
    module = "path-to-regexp",
    version = "3.3.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "punycode",
    module = "punycode",
    version = "2.3.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "range-parser",
    module = "range-parser",
    version = "1.2.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "rc",
    module = "rc",
    version = "1.2.8",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":deep-extend",
        ":ini",
        ":minimist",
        ":strip-json-comments",
    ],
)

node_module(
    name = "registry-auth-token",
    module = "registry-auth-token",
    version = "3.3.2",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":rc",
        ":safe-buffer.5.2.1",
    ],
)

node_module(
    name = "registry-url",
    module = "registry-url",
    version = "3.1.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":rc",
    ],
)

node_module(
    name = "require-from-string",
    module = "require-from-string",
    version = "2.0.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "safe-buffer",
    module = "safe-buffer",
    version = "5.1.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "safe-buffer.5.2.1",
    module = "safe-buffer",
    version = "5.2.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "serve-handler",
    module = "serve-handler",
    version = "6.1.6",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":bytes",
        ":content-disposition",
        ":mime-types",
        ":minimatch",
        ":path-is-inside",
        ":path-to-regexp",
        ":range-parser",
    ],
)

node_module(
    name = "serve",
    module = "serve",
    version = "14.2.4",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":zeit_schemas",
        ":ajv",
        ":arg",
        ":boxen",
        ":chalk",
        ":chalk-template",
        ":clipboardy",
        ":compression",
        ":is-port-reachable",
        ":serve-handler",
        ":update-check",
    ],
    binary = True,
)

node_module(
    name = "shebang-command",
    module = "shebang-command",
    version = "2.0.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":shebang-regex",
    ],
)

node_module(
    name = "shebang-regex",
    module = "shebang-regex",
    version = "3.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "signal-exit",
    module = "signal-exit",
    version = "3.0.7",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "string-width.4.2.3",
    module = "string-width",
    version = "4.2.3",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":emoji-regex.8.0.0",
        ":is-fullwidth-code-point",
        ":strip-ansi.6.0.1",
    ],
)

node_module(
    name = "string-width",
    module = "string-width",
    version = "5.1.2",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":eastasianwidth",
        ":emoji-regex",
        ":strip-ansi",
    ],
)

node_module(
    name = "strip-ansi.6.0.1",
    module = "strip-ansi",
    version = "6.0.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":ansi-regex.5.0.1",
    ],
)

node_module(
    name = "strip-ansi",
    module = "strip-ansi",
    version = "7.1.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":ansi-regex",
    ],
)

node_module(
    name = "strip-final-newline",
    module = "strip-final-newline",
    version = "2.0.0",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "strip-json-comments",
    module = "strip-json-comments",
    version = "2.0.1",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "supports-color",
    module = "supports-color",
    version = "7.2.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":has-flag",
    ],
)

node_module(
    name = "update-check",
    module = "update-check",
    version = "1.5.4",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":registry-auth-token",
        ":registry-url",
    ],
)

node_module(
    name = "uri-js",
    module = "uri-js",
    version = "4.4.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":punycode",
    ],
)

node_module(
    name = "vary",
    module = "vary",
    version = "1.1.2",
    visibility = [
        "PUBLIC",
    ],
)

node_module(
    name = "which",
    module = "which",
    version = "2.0.2",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":isexe",
    ],
)

node_module(
    name = "widest-line",
    module = "widest-line",
    version = "4.0.1",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":string-width",
    ],
)

node_module(
    name = "wrap-ansi",
    module = "wrap-ansi",
    version = "8.1.0",
    visibility = [
        "PUBLIC",
    ],
    deps = [
        ":ansi-styles",
        ":string-width",
        ":strip-ansi",
    ],
)
