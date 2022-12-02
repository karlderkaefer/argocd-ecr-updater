package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"testing"
)

type testObject struct {
	expected      *RepositoryInfo
	input         v1.Secret
	expectedError error
	description   string
}

func TestGetRepositoryInfo(t *testing.T) {
	cases := []testObject{
		{
			input: v1.Secret{
				Data: map[string][]byte{
					"url": []byte("123456789012.dkr.ecr.eu-central-1.amazonaws.com"),
				},
			},
			expected: &RepositoryInfo{
				region:    "eu-central-1",
				accountId: "123456789012",
			},
			expectedError: nil,
			description:   "passing valid registry",
		},
		{
			input: v1.Secret{
				Data: map[string][]byte{
					"url": []byte("023456789012.dkr.ecr.us-east-1.amazonaws.com"),
				},
			},
			expected: &RepositoryInfo{
				region:    "us-east-1",
				accountId: "023456789012",
			},
			expectedError: nil,
			description:   "passing valid registry with other region",
		},
		{
			input: v1.Secret{
				Data: map[string][]byte{
					"url": []byte("1234567.dkr.ecr.eu-central-1.amazonaws.com"),
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("ecr repository is not valid: %s", "1234567.dkr.ecr.eu-central-1.amazonaws.com"),
			description:   "passing wrong account with too less numbers",
		},
		{
			input: v1.Secret{
				Data: map[string][]byte{
					"url": []byte("123456789012.dkr.ecr..amazonaws.com"),
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("ecr repository is not valid: %s", "123456789012.dkr.ecr..amazonaws.com"),
			description:   "passing invalid region",
		},
	}
	for _, c := range cases {
		t.Log(c.description)
		result, err := GetRepositoryInfo(c.input)
		assert.Equal(t, c.expectedError, err)
		assert.Equal(t, c.expected, result)
	}
}
