package kube

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// KubernetesClient copied from https://github.com/argoproj-labs/argocd-image-updater/blob/master/pkg/kube/kubernetes.go
type KubernetesClient struct {
	Clientset     kubernetes.Interface
	CoreClientset v1.CoreV1Interface
	Context       context.Context
	Namespace     string
}

func NewKubernetesClient(ctx context.Context, client kubernetes.Interface, coreClientset v1.CoreV1Interface, namespace string) *KubernetesClient {
	kc := &KubernetesClient{}
	kc.Context = ctx
	kc.Clientset = client
	kc.CoreClientset = coreClientset
	kc.Namespace = namespace
	return kc
}

func NewKubernetesClientFromConfig(ctx context.Context, namespace string, kubeconfig string) (*KubernetesClient, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	loadingRules.ExplicitPath = kubeconfig
	overrides := clientcmd.ConfigOverrides{}
	clientConfig := clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &overrides, os.Stdin)

	config, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	coreClientset := clientset.CoreV1()

	return NewKubernetesClient(ctx, clientset, coreClientset, namespace), nil
}
