package app

import (
	"context"
	"fmt"
	"github.com/karlderkaefer/argocd-ecr-updater/pkg/aws"
	"github.com/karlderkaefer/argocd-ecr-updater/pkg/kube"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const RequiredLabelApp = "argocd-ecr-updater=enabled"
const RequiredLabelArgoCD = "argocd.argoproj.io/secret-type=repository"

type KubernetesAppClient struct {
	kubeClient kube.KubernetesClient
	awsClient  aws.AwsClient
}

func NewKubernetesAppClient(ctx context.Context, client *kube.KubernetesClient, kubeconfig string, namespace string) (*KubernetesAppClient, error) {
	if client == nil {
		defaultClient, err := kube.NewKubernetesClientFromConfig(ctx, namespace, kubeconfig)
		if err != nil {
			return nil, err
		}
		client = defaultClient
	}
	awsClient, err := aws.NewAwsClient()
	if err != nil {
		return nil, err
	}
	return &KubernetesAppClient{kubeClient: *client, awsClient: *awsClient}, nil
}

func (client *KubernetesAppClient) ListSecrets(ctx context.Context) (*v1.SecretList, error) {
	list, err := client.kubeClient.CoreClientset.Secrets(client.kubeClient.Namespace).List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s,%s", RequiredLabelApp, RequiredLabelArgoCD),
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (client *KubernetesAppClient) MutateSecrets(ctx context.Context) error {
	list, err := client.ListSecrets(ctx)
	if err != nil {
		return err
	}
	token, err := client.awsClient.GetToken()
	if len(list.Items) < 1 {
		logrus.Warn("Could not detect any secret matching the required labels")
	}
	for _, secret := range list.Items {
		logrus.Info("Rotate Token for Secret", secret.Name)
		err = client.mutateSecret(ctx, secret, token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *KubernetesAppClient) mutateSecret(ctx context.Context, secret v1.Secret, token string) error {
	secretModified := &secret
	secretModified.Data["password"] = []byte(token)
	_, err := client.kubeClient.CoreClientset.Secrets(client.kubeClient.Namespace).Update(ctx, &secret, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}
