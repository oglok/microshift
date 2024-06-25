package kustomize

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/openshift/microshift/pkg/config"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
	"k8s.io/kubectl/pkg/cmd/create"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
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
		if err := createKustomizationWithRetries(ctx, config.KustomizePathConfig{Path: kustomization}, s.kubeconfig); err != nil {
			klog.Errorf("Creating kustomization at %v failed: %v. Giving up.", kustomization, err)
		} else {
			klog.Infof("Kustomization at %v created successfully.", kustomization)
		}
	}
}

func createKustomizationWithRetries(ctx context.Context, kustomization config.KustomizePathConfig, kubeconfig string) error {
	return wait.PollUntilContextTimeout(ctx, retryInterval, retryTimeout, true, func(_ context.Context) (done bool, err error) {
		if err := createKustomization(kustomization, kubeconfig); err != nil {
			klog.Infof("Creating kustomization failed: %s. Retrying in %s.", err, retryInterval)
			return false, nil
		}
		return true, nil
	})
}

func createKustomization(kustomization config.KustomizePathConfig, kubeconfig string) error {
	klog.Infof("Using create policy with kustomization at %s", kustomization)

	cmds := &cobra.Command{
		Use:   "kubectl",
		Short: "kubectl",
	}
	persistFlags := cmds.PersistentFlags()
	persistFlags.SetNormalizeFunc(cliflag.WarnWordSepNormalizeFunc)
	persistFlags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.KubeConfig = &kubeconfig
	kubeConfigFlags.AddFlags(persistFlags)

	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	matchVersionKubeConfigFlags.AddFlags(persistFlags)

	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)
	ioStreams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	createCmd := create.NewCmdCreate(f, ioStreams)
	cmds.AddCommand(createCmd)

	// Set input/output for the command
	cmds.SetOut(ioStreams.Out)
	cmds.SetErr(ioStreams.ErrOut)
	cmds.SetIn(ioStreams.In)

	// Set the arguments for the command
	cmds.SetArgs([]string{"create", "-f", kustomization.Path})

	// Execute the command
	if err := cmds.Execute(); err != nil {
		return err
	}

	return nil
}
