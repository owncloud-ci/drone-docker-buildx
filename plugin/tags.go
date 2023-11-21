package plugin

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// DefaultTagSuffix returns a set of default suggested tags
// based on the commit ref with an attached suffix.
func DefaultTagSuffix(ref, suffix string) ([]string, error) {
	tags, err := DefaultTags(ref)
	if err != nil {
		return nil, err
	}

	if len(suffix) == 0 {
		return tags, nil
	}

	for i, tag := range tags {
		if tag == "latest" {
			tags[i] = suffix
		} else {
			tags[i] = fmt.Sprintf("%s-%s", tag, suffix)
		}
	}

	return tags, nil
}

// DefaultTags returns a set of default suggested tags based on
// the commit ref.
func DefaultTags(ref string) ([]string, error) {
	if !strings.HasPrefix(ref, "refs/tags/") {
		return []string{"latest"}, nil
	}

	rawVersion := stripTagPrefix(ref)

	version, err := semver.NewVersion(rawVersion)
	if err != nil {
		return []string{"latest"}, err
	}

	if version.PreRelease != "" {
		return []string{
			version.String(),
		}, nil
	}

	if version.Major == 0 {
		return []string{
			fmt.Sprintf("%v.%v", version.Major, version.Minor),
			fmt.Sprintf("%v.%v.%v", version.Major, version.Minor, version.Patch),
		}, nil
	}

	return []string{
		fmt.Sprintf("%v", version.Major),
		fmt.Sprintf("%v.%v", version.Major, version.Minor),
		fmt.Sprintf("%v.%v.%v", version.Major, version.Minor, version.Patch),
	}, nil
}

// UseDefaultTag to keep only default branch for latest tag.
func UseDefaultTag(ref, defaultBranch string) bool {
	if strings.HasPrefix(ref, "refs/tags/") {
		return true
	}

	if stripHeadPrefix(ref) == defaultBranch {
		return true
	}

	return false
}

func stripHeadPrefix(ref string) string {
	return strings.TrimPrefix(ref, "refs/heads/")
}

func stripTagPrefix(ref string) string {
	ref = strings.TrimPrefix(ref, "refs/tags/")
	ref = strings.TrimPrefix(ref, "v")

	return ref
}
