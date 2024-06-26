package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKustomizationPaths(t *testing.T) {
	dataDir, cleanup := setupSuiteDataDir(t)
	defer cleanup()

	kustomizeDirName := func(path string) string {
		return filepath.Join(dataDir, path)
	}

	makeTestKustomizeYaml := func(path string) error {
		filename := filepath.Join(kustomizeDirName(path), "kustomization.yaml")
		return os.WriteFile(filename, []byte{}, 0600)
	}

	makeTestKustomizeYml := func(path string) error {
		filename := filepath.Join(kustomizeDirName(path), "kustomization.yml")
		return os.WriteFile(filename, []byte{}, 0600)
	}

	makeTestKustomize := func(path string) error {
		filename := filepath.Join(kustomizeDirName(path), "Kustomization")
		return os.WriteFile(filename, []byte{}, 0600)
	}

	makeTestKustomizeDir := func(path string) error {
		return os.Mkdir(kustomizeDirName(path), 0700)
	}

	assert.NoError(t, makeTestKustomizeDir("empty"))
	assert.NoError(t, makeTestKustomizeDir("yaml"))
	assert.NoError(t, makeTestKustomizeYaml("yaml"))
	assert.NoError(t, makeTestKustomizeDir("yml"))
	assert.NoError(t, makeTestKustomizeYml("yml"))
	assert.NoError(t, makeTestKustomizeDir("no-ext"))
	assert.NoError(t, makeTestKustomize("no-ext"))

	assert.NoError(t, makeTestKustomizeDir("parent"))
	assert.NoError(t, makeTestKustomizeDir("parent/a"))
	assert.NoError(t, makeTestKustomizeYaml("parent/a"))
	assert.NoError(t, makeTestKustomizeDir("parent/b"))
	assert.NoError(t, makeTestKustomizeYaml("parent/b"))

	var ttests = []struct {
		name        string
		manifests   *Manifests
		results     []string
		expectError bool
	}{
		{
			name: "empty",
			manifests: &Manifests{
				KustomizePathConfigs: []KustomizePathConfig{},
			},
			results: []string{},
		},
		{
			name: "all",
			manifests: &Manifests{
				KustomizePathConfigs: []KustomizePathConfig{
					{
						Path: kustomizeDirName("empty"),
					},
					{
						Path: kustomizeDirName("no-ext"),
					},
					{
						Path: kustomizeDirName("yaml"),
					},
					{
						Path: kustomizeDirName("yml"),
					},
				},
			},
			results: []string{
				kustomizeDirName("no-ext"),
				kustomizeDirName("yaml"),
				kustomizeDirName("yml"),
			},
		},
		{
			name: "o*",
			manifests: &Manifests{
				KustomizePathConfigs: []KustomizePathConfig{
					{
						Path: kustomizeDirName("ya*"),
					},
				},
			},
			results: []string{
				kustomizeDirName("yaml"),
			},
		},
		{
			name: "*o*",
			manifests: &Manifests{
				KustomizePathConfigs: []KustomizePathConfig{
					{
						Path: kustomizeDirName("*m*"),
					},
				},
			},
			results: []string{
				kustomizeDirName("yaml"),
				kustomizeDirName("yml"),
			},
		},
		{
			name: "nomatch",
			manifests: &Manifests{
				KustomizePathConfigs: []KustomizePathConfig{
					{
						Path: kustomizeDirName("nomatch"),
					},
				},
			},
			results: []string{},
		},
		{
			name: "glob-sort",
			manifests: &Manifests{
				KustomizePathConfigs: []KustomizePathConfig{
					{
						Path: kustomizeDirName("parent/*"),
					},
				},
			},
			results: []string{
				kustomizeDirName("parent/a"),
				kustomizeDirName("parent/b"),
			},
		},
	}

	for _, tt := range ttests {
		t.Run(tt.name, func(t *testing.T) {
			paths, err := tt.manifests.GetKustomizationPathConfigs()
			if tt.expectError {
				assert.EqualError(t, err, "")
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, paths, tt.results)
		})
	}
}
