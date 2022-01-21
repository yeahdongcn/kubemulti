package cmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/klog/v2"
)

func kubectl(namespace string, flags []string) ([]metav1.TableColumnDefinition, []metav1.TableRow, error) {
	klog.InfoS("kubectl", "namespace", namespace, "flags", flags)

	columns := []metav1.TableColumnDefinition{}
	rows := []metav1.TableRow{}

	out, err := exec.Command("kubectl", append(flags, "-n", namespace)...).CombinedOutput()
	if err != nil {
		return columns, rows, err
	}

	first := true

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if first {
			first = false
			columns = append(columns, metav1.TableColumnDefinition{
				Name: "namespace",
				Type: "string",
			})
			for _, field := range fields {
				columns = append(columns, metav1.TableColumnDefinition{
					Name: field,
					Type: "string",
				})
			}
		} else {
			row := metav1.TableRow{}
			row.Cells = append(row.Cells, namespace)
			for _, field := range fields {
				row.Cells = append(row.Cells, field)
			}
			rows = append(rows, row)
		}
	}
	return columns, rows, nil
}

func newRootCmd() *cobra.Command {
	o := newRootOption()
	cmd := &cobra.Command{
		Use:                "kubemulti",
		Short:              "short",
		Long:               "long",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, args []string) error {
			var lastKnownError error
			columns := []metav1.TableColumnDefinition{}
			rows := []metav1.TableRow{}

			for _, namespace := range o.Namespaces {
				nsColumns, nsRows, err := kubectl(namespace, args)
				if err != nil {
					lastKnownError = err
					klog.ErrorS(err, "failed to execute kubectl")
					continue
				}
				columns = nsColumns
				rows = append(rows, nsRows...)
			}
			table := &metav1.Table{
				ColumnDefinitions: columns,
				Rows:              rows,
			}
			printer := printers.NewTablePrinter(printers.PrintOptions{})
			err := printer.PrintObj(table, os.Stdout)
			if err != nil {
				lastKnownError = err
				klog.ErrorS(err, "failed to print table")
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
