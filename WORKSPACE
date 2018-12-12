workspace(name = "cilium")

#
# We grep for the following line to generate SOURCE_VERSION file for non-git
# distribution builds. This line must start with the string ENVOY_SHA followed by
# an equals sign and a git SHA in double quotes.
#
# No other line in this file may have ENVOY_SHA followed by an equals sign!
#
ENVOY_SHA = "69c57e632d68d9205d776a91dae8b5eedf23024d"
ENVOY_SHA256 = "f4c238abc526fd889d300a2bf754b57e125031f5bd1c85c338b343e5983d4b46"

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "envoy",
    url = "https://github.com/jrajahalme/envoy/archive/" + ENVOY_SHA + ".zip",
    sha256 = ENVOY_SHA256,
    strip_prefix = "envoy-" + ENVOY_SHA,
)

#
# Bazel does not do transitive dependencies, so we must basically
# include all of Envoy's WORKSPACE file below, with the following
# changes:
# - Skip the 'workspace(name = "envoy")' line as we already defined
#   the workspace above.
# - loads of "//..." need to be renamed as "@envoy//..."
#
load("@envoy//bazel:repositories.bzl", "envoy_dependencies", "GO_VERSION")
load("@envoy//bazel:cc_configure.bzl", "cc_configure")

envoy_dependencies()
cc_configure()

load("@envoy_api//bazel:repositories.bzl", "api_dependencies")
api_dependencies()

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(go_version = GO_VERSION)


# Dependencies for Istio filters.
# Cf. https://github.com/istio/proxy.

ISTIO_PROXY_SHA = "67a0375be569f9158b361e8f5c2a76a0c1b0a02e"
ISTIO_PROXY_SHA256 = "c625d3f9624aa5e3d794181733661da8bf6f4185d0f07dd032a3508b4782ca39"

http_archive(
    name = "istio_proxy",
    url = "https://github.com/istio/proxy/archive/" + ISTIO_PROXY_SHA + ".zip",
    sha256 = ISTIO_PROXY_SHA256,
    strip_prefix = "proxy-" + ISTIO_PROXY_SHA,
)

load("@istio_proxy//:repositories.bzl", "mixerapi_dependencies")
mixerapi_dependencies()

bind(
    name = "boringssl_crypto",
    actual = "//external:ssl",
)
