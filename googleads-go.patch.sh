#!/usr/bin/env bash

set -eu

IMPORT_PATH=$1
VERSION=$2

if [[ $IMPORT_PATH = '' ]]; then
	echo 'PLEASE SET IMPORT_PATH'
	exit 1
fi

if [[ $VERSION = '' ]]; then
	echo 'PLEASE SET VERSION'
	exit 1
fi

echo "patching for $VERSION ..."

cat <<EOF >> googleapis/google/ads/googleads/$VERSION/BUILD.bazel
##############################################################################
# Go
##############################################################################
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "go_gapic_assembly_pkg",
    "go_gapic_library",
    "go_test",
)

go_gapic_library(
    name = "googleads_go_gapic",
    srcs = [
        ":googleads_proto_with_info",
    ],
    gapic_yaml = "googleads_gapic.yaml",
    grpc_service_config = "googleads_grpc_service_config.json",
    service_yaml = "googleads_$VERSION.yaml",
    importpath = "$IMPORT_PATH/$VERSION;googleads",
    deps = [
        "//google/ads/googleads/$VERSION/common:common_go_proto",
        "//google/ads/googleads/$VERSION/enums:enums_go_proto",
        "//google/ads/googleads/$VERSION/resources:resources_go_proto",
        "//google/ads/googleads/$VERSION/services:services_go_proto",
    ],
)

go_test(
    name = "googleads_go_gapic_test",
    srcs = [":googleads_go_gapic_srcjar_test"],
    embed = [":googleads_go_gapic"],
    importpath = "google.golang.org/google/ads/googleads/$VERSION",
)

# Open Source Packages
go_gapic_assembly_pkg(
    name = "gapi-ads-googleads-$VERSION-go",
    deps = [
        ":googleads_go_gapic",
        "//google/ads/googleads/$VERSION/common:common_go_proto",
        "//google/ads/googleads/$VERSION/enums:enums_go_proto",
        "//google/ads/googleads/$VERSION/errors:errors_go_proto",
        "//google/ads/googleads/$VERSION/resources:resources_go_proto",
        "//google/ads/googleads/$VERSION/services:services_go_proto",
    ],
)
EOF

cat <<EOF >> googleapis/google/ads/googleads/$VERSION/common/BUILD.bazel
##############################################################################
# Go
##############################################################################
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "go_proto_library",
)

go_proto_library(
    name = "common_go_proto",
    importpath = "google.golang.org/genproto/googleapis/ads/googleads/$VERSION/common",
    protos = [":common_proto"],
)
EOF

cat <<EOF >> googleapis/google/ads/googleads/$VERSION/enums/BUILD.bazel
##############################################################################
# Go
##############################################################################
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "go_proto_library",
)

go_proto_library(
    name = "enums_go_proto",
    importpath = "google.golang.org/genproto/googleapis/ads/googleads/$VERSION/enums",
    protos = [":enums_proto"],
)
EOF

cat <<EOF >> googleapis/google/ads/googleads/$VERSION/errors/BUILD.bazel
##############################################################################
# Go
##############################################################################
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "go_proto_library",
)

go_proto_library(
    name = "errors_go_proto",
    importpath = "google.golang.org/genproto/googleapis/ads/googleads/$VERSION/errors",
    protos = [":errors_proto"],
)
EOF

cat <<EOF >>googleapis/google/ads/googleads/$VERSION/resources/BUILD.bazel
##############################################################################
# Go
##############################################################################
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "go_proto_library",
)

go_proto_library(
    name = "resources_go_proto",
    importpath = "google.golang.org/genproto/googleapis/ads/googleads/$VERSION/resources",
    protos = [":resources_proto"],
)
EOF

cat <<EOF >>googleapis/google/ads/googleads/$VERSION/services/BUILD.bazel
##############################################################################
# Go
##############################################################################
load(
    "@com_google_googleapis_imports//:imports.bzl",
    "go_proto_library",
)

go_proto_library(
    name = "services_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "google.golang.org/genproto/googleapis/ads/googleads/$VERSION/services",
    protos = [":services_proto"],
)
EOF

echo "done"
