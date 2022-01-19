package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

type RootOptions struct {
	Namespaces []string
}

func NewRootOption() *RootOptions {
	return &RootOptions{
		Namespaces: []string{},
	}
}

func kubectl(args []string) error {
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
					if err := kubectl(args); err != nil {
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
