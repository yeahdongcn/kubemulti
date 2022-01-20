package main

import (
	cmd "github.com/kubemulti/cmd/kubemulti"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()
	cmd.Execute()
}
