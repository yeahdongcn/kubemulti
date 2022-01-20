package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func kubectl(args []string) error {
	klog.InfoS("kubectl", "args", args)
	out, err := exec.Command("kubectl", args...).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func newRootCmd() *cobra.Command {
	o := NewRootOption()
	cmd := &cobra.Command{
		Use:                "x",
		Short:              "short",
		Long:               "long",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(o.Namespaces) == 0 {
				return kubectl(args)
			} else {
				for _, namespace := range o.Namespaces {
					fmt.Println(namespace)
					parameters := append(args, "-n", namespace)
					if err := kubectl(parameters); err != nil {
						klog.ErrorS(err, "failed to invoke kubectl")
						continue
					}
				}
			}
			return nil
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
