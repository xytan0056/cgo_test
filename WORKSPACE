load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-ZzSnGZk7G6Tr6YBuhThkOVqNOWitJ/nddZwZaz6zq+g=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.45.1/rules_go-v0.45.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_pkg_diff",
    importpath = "github.com/pkg/diff",
    sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
    version = "v0.0.0-20210226163009-20ebb0f2a09e",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:cWPaGQEPrBb5/AsnsZesgZZ9yb1OQ+GOISoDNXVBh4M=",
    version = "v1.11.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
    version = "v1.8.4",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
    version = "v1.0.0-20201130134442-10cb98267c6c",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:KENHtAZL2y3NLMYZeHY9DW8HW8V+kQyJsY/V9JlKvCs=",
    version = "v0.9.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:v4INt8xihDGvnrfjMDVXGxw9wrfxYyCjk0KbXjhR55s=",
    version = "v0.0.0-20220722155257-8c9f86f7a55f",
)

# go_repository(
#     name = "org_golang_x_tools",
#     sum = "h1:VveCTK38A2rkS8ZqFY25HIDFscX5X9OoEhJd3quQmXU=",
#     version = "v0.1.12",
#     build_extra_args = [
#         "-go_naming_convention_external=go_default_library",
#         # exclude dirs with intentionally-invalid go code
#         "-exclude=**/testdata",
#     ],
#     build_file_generation = "on",
#     build_file_proto_mode = "disable",
#     importpath = "golang.org/x/tools",
#     replace = "golang.org/x/tools",
# )

http_archive(
    name = "libpam",
    sha256 = "eff47a4ecd833fbf18de9686632a70ee8d0794b79aecb217ebd0ce11db4cd0db",
    strip_prefix = "Linux-PAM-1.3.1/libpam",
    urls = [
        "http://deb.debian.org/debian/pool/main/p/pam/pam_1.3.1.orig.tar.xz",
    ],
    build_file = "//third_party/C/libpam:BUILD.system",
)

go_repository(
    name = "com_github_bazelbuild_rules_python_gazelle",
    build_extra_args = ["-go_naming_convention_external=go_default_library"],
    build_file_generation = "off",
    build_file_proto_mode = "disable",
    importpath = "github.com/bazelbuild/rules_python/gazelle",
    sum = "h1:ZhYrFOC3ffjTNS8bZ8v1zRIeKv0k1SzlMBaW5TbPE/M=",
    version = "v0.0.0-20231221070459-2cbdc1b57bb7",
)

go_repository(
    name = "com_github_emirpasic_gods",
    build_extra_args = ["-go_naming_convention_external=go_default_library"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "github.com/emirpasic/gods",
    sum = "h1:FXtiHYKDGKCW2KzwZKx0iC0PQmdlorYgdFG9jPXJ1Bc=",
    version = "v1.18.1",
)

go_repository(
    name = "com_github_bmatcuk_doublestar",
    build_extra_args = ["-go_naming_convention_external=go_default_library"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "github.com/bmatcuk/doublestar",
    sum = "h1:gPypJ5xD31uhX6Tf54sDPUOBXTqKH4c9aPY66CyQrS0=",
    version = "v1.3.4",
)
go_repository(
    name = "in_gopkg_yaml_v2",
    build_extra_args = ["-go_naming_convention_external=go_default_library"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)
go_rules_dependencies()

go_register_toolchains(version = "1.21.3", nogo = "@//:my_nogo")

gazelle_dependencies()
