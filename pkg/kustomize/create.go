package kustomize

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	"k8s.io/klog/v2"
	"k8s.io/kubectl/pkg/cmd/create"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"

	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/kustomize/api/konfig"
)

func (s *Kustomizer) createKustomizationPath(ctx context.Context, path string) {
	klog.Infof("Using create policy to kustomization at %v", path)

	kustomizationFileNames := konfig.RecognizedKustomizationFileNames()
	for _, filename := range kustomizationFileNames {
		kustomization := filepath.Join(path, filename)

		if _, err := os.Stat(kustomization); errors.Is(err, os.ErrNotExist) {
			klog.Infof("No kustomization found at " + kustomization)
			continue
		}

		klog.Infof("Using create policy to kustomization at %v ", kustomization)
		if err := createKustomizationWithRetries(ctx, path, s.kubeconfig); err != nil {
			klog.Errorf("Creating kustomization at %v failed: %v. Giving up.", kustomization, err)
		} else {
			klog.Infof("Kustomization at %v created successfully.", kustomization)
		}
	}
}

func createKustomizationWithRetries(ctx context.Context, kustomization string, kubeconfig string) error {
	return wait.PollUntilContextTimeout(ctx, retryInterval, retryTimeout, true, func(_ context.Context) (done bool, err error) {
		if err := createKustomization(kustomization, kubeconfig); err != nil {
			klog.Infof("Creating kustomization failed: %s. Retrying in %s.", err, retryInterval)
			return false, nil
		}
		return true, nil
	})
}

func createKustomization(kustomization string, kubeconfig string) error {
	klog.Infof("Using create policy with kustomization at %s", kustomization)

	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.KubeConfig = &kubeconfig

	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)

	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
	ioStreams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	o := create.NewCreateOptions(ioStreams)
	o.FilenameOptions.Kustomize = kustomization
	dummyCmd := &cobra.Command{}
	dummyCmd.Flags().Bool("save-config", false, "If true, save the configuration after applying")
	err := o.RunCreate(f, dummyCmd)
	if err != nil {
		klog.Errorf("Creating kustomization failed: %s", err)
		return nil
	}

	klog.Infof("Kustomization applied successfully")
	return nil
}
