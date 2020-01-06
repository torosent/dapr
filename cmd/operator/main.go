// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main

import (
	"flag"
	"time"

	log "github.com/sirupsen/logrus"
	scheme "dapr/pkg/client/clientset/versioned"
	k8s "dapr/pkg/kubernetes"
	"dapr/pkg/operator"
	"dapr/pkg/signals"
	"dapr/pkg/version"
	"dapr/utils"
	"k8s.io/klog"
)

var (
	logLevel = flag.String("log-level", "info", "Options are debug, info, warning, error, fatal, or panic. (default info)")
)

func main() {
	log.Infof("starting Dapr Operator -- version %s -- commit %s", version.Version(), version.Commit())

	ctx := signals.Context()

	kubeClient := utils.GetKubeClient()

	config := utils.GetConfig()
	daprClient, err := scheme.NewForConfig(config)

	if err != nil {
		log.Fatalf("error building Kubernetes clients: %s", err)
	}

	kubeAPI := k8s.NewAPI(kubeClient, daprClient)

	operator.NewOperator(kubeAPI).Run(ctx)

	shutdownDuration := 5 * time.Second
	log.Infof("allowing %s for graceful shutdown to complete", shutdownDuration)
	<-time.After(shutdownDuration)
}

func init() {
	// This resets the flags on klog, which will otherwise try to log to the FS.
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	klogFlags.Set("logtostderr", "true")

	flag.Parse()

	parsedLogLevel, err := log.ParseLevel(*logLevel)
	if err == nil {
		log.SetLevel(parsedLogLevel)
		log.Infof("log level set to: %s", parsedLogLevel)
	} else {
		log.Fatalf("invalid value for --log-level: %s", *logLevel)
	}
}
