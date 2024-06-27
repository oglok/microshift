package config

import (
	"fmt"
	"path/filepath"
)

const (
	// for files managed via management system in /etc, i.e. user applications
	defaultManifestDirEtc     = "/etc/microshift/manifests"
	defaultManifestDirEtcGlob = "/etc/microshift/manifests.d/*"
	// for files embedded in ostree. i.e. cni/other component customizations
	defaultManifestDirLib     = "/usr/lib/microshift/manifests"
	defaultManifestDirLibGlob = "/usr/lib/microshift/manifests.d/*"
)

type Manifests struct {
	// The locations on the filesystem to scan for kustomization
	// files to use to load manifests. Set to a list of paths to scan
	// only those paths. Set to an empty list to disable loading
	// manifests. The entries in the list can be glob patterns to
	// match multiple subdirectories.
	//
	// +kubebuilder:default={"/usr/lib/microshift/manifests","/usr/lib/microshift/manifests.d/*","/etc/microshift/manifests","/etc/microshift/manifests.d/*"}
	KustomizePaths       []string `json:"kustomizePaths"`
	KustomizeCreatePaths []string `json:"kustomizeCreatePaths"`
}

func (m *Manifests) GetAllKustomizationPaths() ([]string, []string, error) {
	kustomizePaths, err := m.findKustomizationPaths(m.KustomizePaths)
	if err != nil {
		return nil, nil, err
	}

	kustomizeCreatePaths, err := m.findKustomizationPaths(m.KustomizeCreatePaths)
	if err != nil {
		return nil, nil, err
	}

	return kustomizePaths, kustomizeCreatePaths, nil
}

func (m *Manifests) findKustomizationPaths(paths []string) ([]string, error) {
	var results []string
	for _, path := range paths {
		matches, err := filepath.Glob(path)
		if err != nil {
			return nil, fmt.Errorf("error matching path %s: %w", path, err)
		}
		results = append(results, matches...)
	}
	return results, nil
}
