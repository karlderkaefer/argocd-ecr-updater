# ArgoCD ECR updater

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/argocd-aws-ecr-updater)](https://artifacthub.io/packages/search?repo=argocd-aws-ecr-updater)

If you are using a private [AWS ECR repository](https://docs.aws.amazon.com/AmazonECR/latest/userguide/push-oci-artifact.html) 
to store helm charts, the stored password will become expired at latest in `12h`.  
The argocd-ecr-updater will refresh the token in defined interval.

The updater will only consider secrets with these labels for update.
```yaml
kind: Secret
metadata:
  labels:
    argocd-ecr-updater: enabled
    argocd.argoproj.io/secret-type: repository
```
The data field `password` will be updated with a fresh token from AWS ECR.

To give argocd-ecr-updater permission to get the ECR token create [IRSA role](https://github.com/terraform-aws-modules/terraform-aws-iam/tree/v5.5.5/modules/iam-assumable-role-with-oidc) with following permissions
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": [
        "ecr:ListTagsForResource",
        "ecr:ListImages",
        "ecr:GetRepositoryPolicy",
        "ecr:GetLifecyclePolicyPreview",
        "ecr:GetLifecyclePolicy",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetAuthorizationToken",
        "ecr:DescribeRepositories",
        "ecr:DescribeImages",
        "ecr:DescribeImageScanFindings",
        "ecr:BatchGetImage",
        "ecr:BatchCheckLayerAvailability"
      ],
      "Resource": "*"
    }
  ]
}
```
As alternative, you can also use managed AWS role `AmazonEC2ContainerRegistryReadOnly`.

If the ECR is not in the same account as the kubernetes cluster, you may want to adapt ECR Access Permission. 
In this example, I have one dedicated user `ecr.circleci` that can write (publish charts) to ECR.
While argocd just need to pull helm charts, you can create separate read permission and allow accounts to read from ECR where the kubernetes clusters runs (account `555555555555`).

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:user/ecr.circleci"
      },
      "Action": [
        "ecr:BatchCheckLayerAvailability",
        "ecr:BatchGetImage",
        "ecr:CompleteLayerUpload",
        "ecr:DescribeImages",
        "ecr:DescribeRepositories",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:InitiateLayerUpload",
        "ecr:ListImages",
        "ecr:PutImage",
        "ecr:UploadLayerPart"
      ]
    },
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": [
          "arn:aws:iam::555555555555:root",
          ""
        ]
      },
      "Action": [
        "ecr:BatchCheckLayerAvailability",
        "ecr:BatchGetImage",
        "ecr:DescribeImages",
        "ecr:DescribeRepositories",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:ListImages"
      ]
    }
  ]
}
```

The token is valid to authenticate against any registry id, the user has access. 
Depending on how you set the trust relationship on ECR repository policy. 

Finally annotate the service account with your role arn
```yaml
serviceAccount:
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::123456:role/argocd-ecr-updater
```

## Use Helm Charts from ECR

* [ArgoCD App](example/example-argocd-app.yaml)
* [Helm Chart](example/Chart.yaml)


## Install with Helm
```bash
helm repo add argocd-ecr-updater https://karlderkaefer.github.io/argocd-ecr-updater
helm search repo argocd-ecr-updater 
helm upgrade --install argocd-ecr-updater -n argocd argocd-ecr-updater/argocd-ecr-updater
```


## Usage CLI
```bash
Usage:
  argocd-ecr-updater [flags]

Flags:
  -h, --help                help for argocd-ecr-updater
      --interval string     interval to refresh token (default "6h")
      --kubeconfig string   kubernetes config file
      --namespace string    kubernetes namespace
```
* `--kubeconfig`: will use in-cluster config by default, optional you can provide own kubeconfig for testing
* `--interval`: defined in which interval the token will refreshed
* `--namespace`: if empty, then mutate secrets from all namespaces matching the label

You can also set these values by providing environment variable with prefix `ARGOCD_ECR_UPDATER`
```bash
ARGOCD_ECR_UPDATER_NAMESPACE="argocd"
ARGOCD_ECR_UPDATER_INTERVAL="6h0m0s"
ARGOCD_ECR_UPDATER_KUBECONFIG="/home/user/.kube/config"
```

## Related GitHub Issues
* https://github.com/argoproj/argo-cd/issues/8097
* https://github.com/argoproj/argo-cd/issues/8952

## Alternative

As alternative, you can use [External Secrets](https://external-secrets.io/v0.6.1/guides/common-k8s-secret-types/)
