package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["checks.go"],
    tags = ["automanaged"],
    deps = [
        "//cmd/kubeadm/app/apis/kubeadm:go_default_library",
        "//pkg/api/validation:go_default_library",
        "//pkg/util/errors:go_default_library",
        "//pkg/util/initsystem:go_default_library",
        "//pkg/util/node:go_default_library",
        "//test/e2e_node/system:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["checks_test.go"],
    library = ":go_default_library",
    tags = ["automanaged"],
)
