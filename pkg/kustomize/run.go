package kustomize

import (
	"context"
	"fmt"
	"time"

	"github.com/openshift/microshift/pkg/config"
)

const (
	retryInterval = 10 * time.Second
	retryTimeout  = 1 * time.Minute
)

type Kustomizer struct {
	cfg        *config.Config
	kubeconfig string
}

func NewKustomizer(cfg *config.Config) *Kustomizer {
	return &Kustomizer{
		cfg:        cfg,
		kubeconfig: cfg.KubeConfigPath(config.KubeAdmin),
	}
}

func (s *Kustomizer) Name() string           { return "kustomizer" }
func (s *Kustomizer) Dependencies() []string { return []string{"kube-apiserver"} }

func (s *Kustomizer) Run(ctx context.Context, ready chan<- struct{}, stopped chan<- struct{}) error {
	defer close(stopped)
	defer close(ready)

	kustomizationConfigs, err := s.cfg.Manifests.GetKustomizationConfigs()
	if err != nil {
		return fmt.Errorf("failed to find any kustomization paths: %w", err)
	}

	for _, config := range kustomizationConfigs {
		switch config.Policy {
		case "apply":
			s.applyKustomizationPath(ctx, config.Path)
		case "create":
			s.createKustomizationPath(ctx, config.Path)
		default:
			return fmt.Errorf("unknown policy %s for path %s", config.Policy, config.Path)
		}
	}

	return ctx.Err()
}
