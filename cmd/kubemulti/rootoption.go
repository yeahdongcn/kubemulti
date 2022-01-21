package cmd

import (
	"github.com/kubemulti/pkg/kubeconfig"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

type rootOptions struct {
	Namespaces []string
}

func getCurrentNamespace() (string, error) {
	kc := new(kubeconfig.Kubeconfig).WithLoader(kubeconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		return "", errors.Wrap(err, "kubeconfig error")
	}

	ctx := kc.GetCurrentContext()
	if ctx == "" {
		return "", errors.New("current-context is not set")
	}
	ns, err := kc.NamespaceOfContext(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read namespace of \"%s\"", ctx)
	}
	return ns, nil
}

func newRootOption() *rootOptions {
	ns, err := getCurrentNamespace()
	if err != nil {
		klog.ErrorS(err, "failed to get current namespace")
	}
	return &rootOptions{
		Namespaces: []string{ns},
	}
}
