package main

import (
	"github.com/owncloud-ci/drone-docker-buildx/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
//
//nolint:maintidx
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "dry-run",
			EnvVars:     []string{"PLUGIN_DRY_RUN"},
			Usage:       "disable docker push",
			Destination: &settings.Dryrun,
		},
		&cli.StringFlag{
			Name:        "daemon.mirror",
			EnvVars:     []string{"PLUGIN_MIRROR", "DOCKER_PLUGIN_MIRROR"},
			Usage:       "registry mirror to pull images",
			Destination: &settings.Daemon.Mirror,
		},
		&cli.StringFlag{
			Name:        "daemon.storage-driver",
			EnvVars:     []string{"PLUGIN_STORAGE_DRIVER"},
			Usage:       "docker daemon storage driver",
			Destination: &settings.Daemon.StorageDriver,
		},
		&cli.StringFlag{
			Name:        "daemon.storage-path",
			EnvVars:     []string{"PLUGIN_STORAGE_PATH"},
			Usage:       "docker daemon storage path",
			Value:       "/var/lib/docker",
			Destination: &settings.Daemon.StoragePath,
		},
		&cli.StringFlag{
			Name:        "daemon.bip",
			EnvVars:     []string{"PLUGIN_BIP"},
			Usage:       "allow the docker daemon to bride ip address",
			Destination: &settings.Daemon.Bip,
		},
		&cli.StringFlag{
			Name:        "daemon.mtu",
			EnvVars:     []string{"PLUGIN_MTU"},
			Usage:       "docker daemon custom mtu setting",
			Destination: &settings.Daemon.MTU,
		},
		&cli.StringSliceFlag{
			Name:        "daemon.dns",
			EnvVars:     []string{"PLUGIN_CUSTOM_DNS"},
			Usage:       "custom docker daemon dns server",
			Destination: &settings.Daemon.DNS,
		},
		&cli.StringSliceFlag{
			Name:        "daemon.dns-search",
			EnvVars:     []string{"PLUGIN_CUSTOM_DNS_SEARCH"},
			Usage:       "custom docker daemon dns search domain",
			Destination: &settings.Daemon.DNSSearch,
		},
		&cli.BoolFlag{
			Name:        "daemon.insecure",
			EnvVars:     []string{"PLUGIN_INSECURE"},
			Usage:       "allow the docker daemon to use insecure registries",
			Value:       false,
			Destination: &settings.Daemon.Insecure,
		},
		&cli.BoolFlag{
			Name:        "daemon.ipv6",
			EnvVars:     []string{"PLUGIN_IPV6"},
			Usage:       "enable docker daemon IPv6 support",
			Value:       false,
			Destination: &settings.Daemon.IPv6,
		},
		&cli.BoolFlag{
			Name:        "daemon.experimental",
			EnvVars:     []string{"PLUGIN_EXPERIMENTAL"},
			Usage:       "enable docker daemon experimental mode",
			Value:       false,
			Destination: &settings.Daemon.Experimental,
		},
		&cli.BoolFlag{
			Name:        "daemon.debug",
			EnvVars:     []string{"PLUGIN_DEBUG"},
			Usage:       "enable verbose debug mode for the docker daemon",
			Value:       false,
			Destination: &settings.Daemon.Debug,
		},
		&cli.BoolFlag{
			Name:        "daemon.off",
			EnvVars:     []string{"PLUGIN_DAEMON_OFF"},
			Usage:       "disable the startup of the docker daemon",
			Value:       false,
			Destination: &settings.Daemon.Disabled,
		},
		&cli.StringFlag{
			Name:        "daemon.buildkit-config",
			EnvVars:     []string{"PLUGIN_BUILDKIT_CONFIG"},
			Usage:       "content of the docker buildkit toml config",
			Destination: &settings.Daemon.BuildkitConfig,
		},
		&cli.StringFlag{
			Name:        "dockerfile",
			EnvVars:     []string{"PLUGIN_DOCKERFILE"},
			Usage:       "dockerfile to use for the image build",
			Value:       "Dockerfile",
			Destination: &settings.Build.Dockerfile,
		},
		&cli.StringFlag{
			Name:        "context",
			EnvVars:     []string{"PLUGIN_CONTEXT"},
			Usage:       "path of the build context",
			Value:       ".",
			Destination: &settings.Build.Context,
		},
		&cli.StringSliceFlag{
			Name:        "named-context",
			EnvVars:     []string{"PLUGIN_NAMED_CONTEXT"},
			Usage:       "additional named build context",
			Destination: &settings.Build.NamedContext,
		},
		&cli.StringSliceFlag{
			Name:        "tags",
			EnvVars:     []string{"PLUGIN_TAG", "PLUGIN_TAGS"},
			Usage:       "repository tags to use for the image",
			FilePath:    ".tags",
			Destination: &settings.Build.Tags,
		},
		&cli.BoolFlag{
			Name:        "tags.auto",
			EnvVars:     []string{"PLUGIN_DEFAULT_TAGS", "PLUGIN_AUTO_TAG"},
			Usage:       "generate tag names automatically based on git branch and git tag",
			Value:       false,
			Destination: &settings.Build.TagsAuto,
		},
		&cli.StringFlag{
			Name:        "tags.suffix",
			EnvVars:     []string{"PLUGIN_DEFAULT_SUFFIX", "PLUGIN_AUTO_TAG_SUFFIX"},
			Usage:       "generate tag names with the given suffix",
			Destination: &settings.Build.TagsSuffix,
		},
		&cli.StringSliceFlag{
			Name:        "extra.tags",
			EnvVars:     []string{"PLUGIN_EXTRA_TAGS"},
			Usage:       "additional tags to use for the image including registry",
			FilePath:    ".extratags",
			Destination: &settings.Build.ExtraTags,
		},
		&cli.StringSliceFlag{
			Name:        "args",
			EnvVars:     []string{"PLUGIN_BUILD_ARGS"},
			Usage:       "custom build arguments for the build",
			Destination: &settings.Build.Args,
		},
		&cli.StringSliceFlag{
			Name:        "args-from-env",
			EnvVars:     []string{"PLUGIN_BUILD_ARGS_FROM_ENV"},
			Usage:       "forward environment variables as custom arguments to the build",
			Destination: &settings.Build.ArgsEnv,
		},
		&cli.BoolFlag{
			Name:        "quiet",
			EnvVars:     []string{"PLUGIN_QUIET"},
			Usage:       "enable suppression of the build output",
			Value:       false,
			Destination: &settings.Build.Quiet,
		},
		&cli.StringFlag{
			Name:        "output",
			EnvVars:     []string{"PLUGIN_OUTPUT"},
			Usage:       "export action for the build result",
			Destination: &settings.Build.Output,
		},
		&cli.StringFlag{
			Name:        "target",
			EnvVars:     []string{"PLUGIN_TARGET"},
			Usage:       "build target to use",
			Destination: &settings.Build.Target,
		},
		&cli.GenericFlag{
			Name:    "cache-from",
			EnvVars: []string{"PLUGIN_CACHE_FROM"},
			Usage:   "images to consider as cache sources",
			Value:   &plugin.StringSliceFlag{},
		},
		&cli.StringFlag{
			Name:        "cache-to",
			EnvVars:     []string{"PLUGIN_CACHE_TO"},
			Usage:       "cache destination for the build cache",
			Destination: &settings.Build.CacheTo,
		},
		&cli.BoolFlag{
			Name:        "pull-image",
			EnvVars:     []string{"PLUGIN_PULL_IMAGE"},
			Usage:       "enforce to pull base image at build time",
			Value:       true,
			Destination: &settings.Build.Pull,
		},
		&cli.BoolFlag{
			Name:        "compress",
			EnvVars:     []string{"PLUGIN_COMPRESS"},
			Usage:       "enable compression of the build context using gzip",
			Value:       false,
			Destination: &settings.Build.Compress,
		},
		&cli.StringFlag{
			Name:        "repo",
			EnvVars:     []string{"PLUGIN_REPO"},
			Usage:       "repository name for the image",
			Destination: &settings.Build.Repo,
		},
		&cli.StringFlag{
			Name:        "docker.registry",
			EnvVars:     []string{"PLUGIN_REGISTRY", "DOCKER_REGISTRY"},
			Usage:       "docker registry to authenticate with",
			Value:       "https://index.docker.io/v1/",
			Destination: &settings.Login.Registry,
		},
		&cli.StringFlag{
			Name:        "docker.username",
			EnvVars:     []string{"PLUGIN_USERNAME", "DOCKER_USERNAME"},
			Usage:       "username for registry authentication",
			Destination: &settings.Login.Username,
		},
		&cli.StringFlag{
			Name:        "docker.password",
			EnvVars:     []string{"PLUGIN_PASSWORD", "DOCKER_PASSWORD"},
			Usage:       "password for registry authentication",
			Destination: &settings.Login.Password,
		},
		&cli.StringFlag{
			Name:        "docker.email",
			EnvVars:     []string{"PLUGIN_EMAIL", "DOCKER_EMAIL"},
			Usage:       "email address for registry authentication",
			Destination: &settings.Login.Email,
		},
		&cli.StringFlag{
			Name:        "docker.config",
			EnvVars:     []string{"PLUGIN_CONFIG", "DOCKER_PLUGIN_CONFIG"},
			Usage:       "content of the docker daemon json config",
			Destination: &settings.Login.Config,
		},
		&cli.BoolFlag{
			Name:        "no-cache",
			EnvVars:     []string{"PLUGIN_NO_CACHE"},
			Usage:       "disable the usage of cached intermediate containers",
			Value:       false,
			Destination: &settings.Build.NoCache,
		},
		&cli.StringSliceFlag{
			Name:        "add-host",
			EnvVars:     []string{"PLUGIN_ADD_HOST"},
			Usage:       "additional host:ip mapping",
			Destination: &settings.Build.AddHost,
		},
		&cli.StringSliceFlag{
			Name:        "platforms",
			EnvVars:     []string{"PLUGIN_PLATFORMS"},
			Usage:       "target platform for build",
			Destination: &settings.Build.Platforms,
		},
		&cli.StringSliceFlag{
			Name:        "labels",
			EnvVars:     []string{"PLUGIN_LABELS"},
			Usage:       "labels to add to image",
			Destination: &settings.Build.Labels,
		},
		&cli.StringFlag{
			Name:        "provenance",
			EnvVars:     []string{"PLUGIN_PROVENANCE"},
			Usage:       "generates provenance attestation for the build",
			Destination: &settings.Build.Provenance,
		},
		&cli.StringFlag{
			Name:        "sbom",
			EnvVars:     []string{"PLUGIN_SBOM"},
			Usage:       "generates sbom attestation for the build",
			Destination: &settings.Build.SBOM,
		},
		&cli.GenericFlag{
			Name:    "secrets",
			EnvVars: []string{"PLUGIN_SECRETS"},
			Usage:   "exposes secrets to the build",
			Value:   &plugin.StringSliceFlag{},
		},
	}
}
