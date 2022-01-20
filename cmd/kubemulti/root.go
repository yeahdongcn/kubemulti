package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func kubectl(flags []string) (string, error) {
	klog.InfoS("kubectl", "flags", flags)
	out, err := exec.Command("kubectl", flags...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func newRootCmd() *cobra.Command {
	o := NewRootOption()
	cmd := &cobra.Command{
		Use:                "x",
		Short:              "short",
		Long:               "long",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, args []string) error {
			var lastKnownError error
			for _, namespace := range o.Namespaces {
				flags := append(args, "-n", namespace)
				out, err := kubectl(flags)
				if err != nil {
					lastKnownError = err
					klog.ErrorS(err, "failed to invoke kubectl")
					continue
				}
				fmt.Println(namespace)
				fmt.Println(out)
			}
			return lastKnownError
		},
	}
	cmd.Flags().StringArrayVarP(&o.Namespaces, "namespace", "n", o.Namespaces, "xyz")
	return cmd
}

// Execute executes the root command.
func Execute() error {
	rootCmd := newRootCmd()
	return rootCmd.Execute()
}
