package aws

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/sirupsen/logrus"
	"strings"
)

type AwsClient struct {
	ecrClient *ecr.Client
}

type RepositoryInfo struct {
	Name string
	Id   string
}

func NewAwsClient(region string) (*AwsClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}
	client := ecr.NewFromConfig(cfg)
	return &AwsClient{ecrClient: client}, nil
}

func GetToken(region string) (string, error) {
	client, err := NewAwsClient(region)
	if err != nil {
		return "", err
	}
	resp, err := client.ecrClient.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		logrus.Errorf("unable to load SDK config, %v", err)
		return "", err
	}
	if len(resp.AuthorizationData) < 1 {
		logrus.Errorln("length of auhtorization data must be greater than 0")
		return "", errors.New("length of auhtorization data must be greater than 0")
	}
	token := resp.AuthorizationData[0].AuthorizationToken

	expires := resp.AuthorizationData[0].ExpiresAt
	if token == nil {
		return "", errors.New("invalid token")
	}
	// returns b64 decoded string AWS:ey...
	decToken, err := base64.StdEncoding.DecodeString(*token)
	if err != nil {
		return "", err
	}
	if !strings.Contains(string(decToken), ":") {
		return "", errors.New("invalid token")
	}
	password := strings.Split(string(decToken), ":")[1]
	logrus.Debug("found token. token will expire at ", expires)
	return password, nil
}

// ValidateAccess unused
func (client *AwsClient) ValidateAccess(repository RepositoryInfo) error {
	images, err := client.ecrClient.ListImages(context.TODO(), &ecr.ListImagesInput{
		RegistryId:     aws.String(repository.Id),
		RepositoryName: aws.String(repository.Name),
	})
	if err != nil {
		return fmt.Errorf("unable to authenticate to repository id %s repository name %s, %v", repository.Id, repository.Name, err)
	}
	if len(images.ImageIds) < 1 {
		logrus.Warn("detected no image tags in repository")
	}
	return nil
}
