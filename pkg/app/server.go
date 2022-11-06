// Package app
// Logic mainly copied from ArgoCD Image Updater
// https://github.com/argoproj-labs/argocd-image-updater/blob/master/cmd/run.go
package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Run(ctx context.Context, kubeconfig string, namespace string, interval string) error {
	var hsErrCh chan error
	logrus.Infof("Starting health probe server TCP port=%d", 8080)
	hsErrCh = startHealthServer(8080)

	refreshInterval, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	logrus.Infof("starting refreshing token every %s", refreshInterval)

	for {
		select {
		case err := <-hsErrCh:
			if err != nil {
				logrus.Errorf("Health probe server exited with error: %v", err)
			} else {
				logrus.Infof("Health probe server exited gracefully")
			}
			return nil
		default:
			logrus.Infof("Refreshing Token Interval started")
			client, err := NewKubernetesAppClient(ctx, nil, kubeconfig, namespace)
			if err != nil {
				return err
			}
			err = client.MutateSecrets(ctx)
			if err != nil {
				return err
			}
		}
		time.Sleep(refreshInterval)
	}
}

func startHealthServer(port int) chan error {
	errCh := make(chan error)
	go func() {
		sm := http.NewServeMux()
		sm.HandleFunc("/healthz", healthProbe)
		errCh <- http.ListenAndServe(fmt.Sprintf(":%d", port), sm)
	}()
	return errCh
}

func healthProbe(w http.ResponseWriter, r *http.Request) {
	logrus.Tracef("/healthz ping request received, replying with pong")
	_, _ = fmt.Fprintf(w, "OK\n")
}
