load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "sigs.k8s.io/cluster-api-provider-azure/cmd/clusterctl",
    visibility = ["//visibility:private"],
    deps = [
        "//cmd/versioninfo:go_default_library",
        "//pkg/cloud/azure/actuators/cluster:go_default_library",
        "//vendor/github.com/openshift/cluster-api/cmd/clusterctl/cmd:go_default_library",
        "//vendor/github.com/openshift/cluster-api/pkg/apis/cluster/common:go_default_library",
    ],
)

go_binary(
    name = "clusterctl",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
    x_defs = {
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitBranch": "{GIT_BRANCH}",
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitReleaseTag": "{GIT_RELEASE_TAG}",
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitReleaseCommit": "{GIT_RELEASE_COMMIT}",
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitCommit": "{GIT_COMMIT}",
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitTreeState": "{GIT_TREE_STATE}",
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitMajor": "{GIT_MAJOR}",
        "sigs.k8s.io/cluster-api-provider-azure/cmd/versioninfo.GitMinor": "{GIT_MINOR}",
    },
)
