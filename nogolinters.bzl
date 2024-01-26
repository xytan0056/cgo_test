load("@io_bazel_rules_go//go:def.bzl", "nogo")
load("@bazel_skylib//rules:common_settings.bzl", "string_flag")

EXPERIMENTAL_LINTERS = []

DEFAULT_LINTERS = [
    # go vet analyzers
    "@org_golang_x_tools//go/analysis/passes/asmdecl:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/assign:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/atomic:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/bools:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/buildtag:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/composite:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/copylock:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/fieldalignment:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/httpresponse:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/loopclosure:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/lostcancel:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/nilfunc:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/printf:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/shift:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/stdmethods:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/stringintconv:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/structtag:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/tests:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/unmarshal:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/unsafeptr:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/unusedresult:go_default_library",

    "@org_golang_x_tools//go/analysis/passes/atomicalign:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/ctrlflow:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/deepequalerrors:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/findcall:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/inspect:go_default_library",
    "@org_golang_x_tools//go/analysis/passes/shadow:go_default_library",
]

# NilAway is separated since we need to run default linters without it as a reference to continuously monitor the performance overhead of NilAway.
NILAWAY = [

]

def nogo_linters():
    string_flag(
        name = "nogo",
        build_setting_default = "default",
    )

    native.config_setting(
        name = "nogo_disabled",
        flag_values = {
            ":nogo": "disabled",
        },
    )

    native.config_setting(
        name = "nogo_all",
        flag_values = {
            ":nogo": "all",
        },
    )

    native.config_setting(
        name = "nogo_only_experimental",
        flag_values = {
            ":nogo": "experimental",
        },
    )

    native.config_setting(
        name = "nogo_default_no_nilaway",
        flag_values = {
            ":nogo": "no-nilaway",
        },
    )

    nogo(
        name = "my_nogo",
        config = select({
            ":nogo_disabled": None,
            "//conditions:default": ":nogo.json",
        }),
        visibility = ["//visibility:public"],
        deps = select({
            ":nogo_disabled": [],
            ":nogo_only_experimental": EXPERIMENTAL_LINTERS,
            ":nogo_default_no_nilaway": DEFAULT_LINTERS,
            ":nogo_all": EXPERIMENTAL_LINTERS + DEFAULT_LINTERS + NILAWAY,
            "//conditions:default": DEFAULT_LINTERS + NILAWAY,
        }),
    )
