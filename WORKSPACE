workspace(name = "com_github_tcncloud_protoc_gen_state")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.6/rules_go-0.16.6.tar.gz"],
    sha256 = "ade51a315fa17347e5c31201fdc55aa5ffb913377aa315dceb56ee9725e620ee",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.16.0/bazel-gazelle-0.16.0.tar.gz"],
    sha256 = "7949fc6cc17b5b191103e97481cf8889217263acf52e00b560683413af204fcb",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

# go_proto_dependencies()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.1",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    tag = "v1.4.7",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_iancoleman_strcase",
    commit = "3605ed457bf7",
    importpath = "github.com/iancoleman/strcase",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    tag = "v1.0.1",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    tag = "v1.7.0",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    tag = "v1.4.3",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.2.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.2.2",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "788fd7840127",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    tag = "v1.4.7",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    commit = "dd632973f1e7",
    importpath = "gopkg.in/tomb.v1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.2",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "0709b304e793",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "351d144fa1fc",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sync",
    commit = "42b317875d0f",
    importpath = "golang.org/x/sync",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "4ed8d59d0b35",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.0",
)
