package main

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()

	clientCfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, nil)
	cfg, err := clientCfg.ClientConfig()
	failOnError(err)

	clientset := kubernetes.NewForConfigOrDie(cfg)

	ns, _, err := clientCfg.Namespace()
	failOnError(err)

	// get deployments using standard k8s client
	deps, err := clientset.AppsV1().Deployments(ns).List(context.TODO(), v1.ListOptions{})
	failOnError(err)

	log.Println("Deployments: ")
	for _, dep := range deps.Items {
		log.Println(dep.Name)
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
