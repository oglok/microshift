package config

import (
	"fmt"
	"path/filepath"
	"sort"

	"k8s.io/klog/v2"
	"sigs.k8s.io/kustomize/api/konfig"
)

const (
	// for files managed via management system in /etc, i.e. user applications
	defaultManifestDirEtc     = "/etc/microshift/manifests"
	defaultManifestDirEtcGlob = "/etc/microshift/manifests.d/*"
	// for files embedded in ostree. i.e. cni/other component customizations
	defaultManifestDirLib     = "/usr/lib/microshift/manifests"
	defaultManifestDirLibGlob = "/usr/lib/microshift/manifests.d/*"
)

type KustomizePathConfig struct {
	Path   string `json:"path"`
	Policy string `json:"policy"` // "apply" as default, "create" as an alternative
}

type Manifests struct {
	// The locations on the filesystem to scan for kustomization
	// files to use to load manifests. Set to a list of paths to scan
	// only those paths. Set to an empty list to disable loading
	// manifests. The entries in the list can be glob patterns to
	// match multiple subdirectories.
	//
	// +kubebuilder:default={{Path: "/usr/lib/microshift/manifests", Policy: "apply"}, {Path: "/usr/lib/microshift/manifests.d/*", Policy: "apply"}, {Path: "/etc/microshift/manifests", Policy: "apply"}, {Path: "/etc/microshift/manifests.d/*", Policy: "apply"}}
	KustomizePathConfigs []KustomizePathConfig `json:"kustomizePathConfigs"`
}

// GetKustomizationPaths returns the list of configured paths for
// which there are actual kustomization files to be loaded. The paths
// are returned in the order given in the configuration file. The
// results of any glob patterns are sorted lexicographically.
func (m *Manifests) GetKustomizationPathConfigs() ([]KustomizePathConfig, error) {
	kustomizationFileNames := konfig.RecognizedKustomizationFileNames()
	results := []KustomizePathConfig{}

	for _, kustomizePathConfig := range m.KustomizePathConfigs {
		for _, filename := range kustomizationFileNames {
			pattern := filepath.Join(kustomizePathConfig.Path, filename)
			matches, err := filepath.Glob(pattern)
			if err != nil {
				return nil, fmt.Errorf("could not understand kustomizePath value %v: %w", kustomizePathConfig.Path, err)
			}
			if len(matches) == 0 {
				klog.Infof("No kustomize path matches %v", pattern)
				continue
			}
			// We add the filename to the pattern so we only return
			// directories where there is something to apply, but the
			// results we need are the directory names, so convert the
			// full match string back to a directory.
			//
			// Glob() does not explicitly say it sorts its return
			// value, so we do it to ensure deterministic behavior.
			sort.Strings(matches)
			for _, match := range matches {
				dir := filepath.Dir(match)
				klog.Infof("Adding kustomize path %v with policy %v", dir, kustomizePathConfig.Policy)
				results = append(results, KustomizePathConfig{Path: dir, Policy: kustomizePathConfig.Policy})
			}
		}
	}
	return results, nil
}
