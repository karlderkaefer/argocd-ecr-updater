package app

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/karlderkaefer/argocd-ecr-updater/pkg/aws"
	"github.com/karlderkaefer/argocd-ecr-updater/pkg/kube"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const RequiredLabelApp = "argocd-ecr-updater=enabled"

// https://github.com/argoproj/argo-cd/issues/19881
// https://argo-cd.readthedocs.io/en/latest/operator-manual/argocd-repo-creds-yaml/
const RequiredLabelArgoCD = "argocd.argoproj.io/secret-type=repo-creds"

type KubernetesAppClient struct {
	kubeClient kube.KubernetesClient
}

type RepositoryInfo struct {
	region    string
	accountId string
}

func NewKubernetesAppClient(ctx context.Context, client *kube.KubernetesClient, kubeconfig string, namespace string) (*KubernetesAppClient, error) {
	if client == nil {
		defaultClient, err := kube.NewKubernetesClientFromConfig(ctx, namespace, kubeconfig)
		if err != nil {
			return nil, err
		}
		client = defaultClient
	}
	return &KubernetesAppClient{kubeClient: *client}, nil
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

func GetRepositoryInfo(secret v1.Secret) (*RepositoryInfo, error) {
	registry := string(secret.Data["url"])
	if registry == "" {
		return nil, errors.New("required field url is empty")
	}
	reg := regexp.MustCompile(`(?P<AccountId>\d{12}).dkr\.ecr\.(?P<region>[A-Za-z0-9-]+)\.amazonaws\.com`)
	if !reg.MatchString(registry) {
		return nil, fmt.Errorf("ecr repository is not valid: %s", registry)
	}
	matches := reg.FindStringSubmatch(registry)
	if len(matches) < 3 {
		return nil, fmt.Errorf("ecr repository is not valid: %s", registry)
	}
	repository := &RepositoryInfo{
		region:    matches[2],
		accountId: matches[1],
	}
	return repository, nil
}

func (client *KubernetesAppClient) MutateSecrets(ctx context.Context) error {
	list, err := client.ListSecrets(ctx)
	if err != nil {
		return err
	}
	if len(list.Items) < 1 {
		logrus.Warn("Could not detect any secret matching the required labels")
	}
	for _, secret := range list.Items {
		registry, err := GetRepositoryInfo(secret)
		if err != nil {
			logrus.Errorf("Could not read required region info out of secret %s %v", secret.Name, err)
			break
		}
		token, err := aws.GetToken(registry.region)
		if err != nil {
			logrus.Errorf("Error when authenticate to registry %s in region %s for secret %s %v", registry.accountId, registry.region, secret.Name, err)
			break
		}
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
