package main

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	loader := clientcmd.NewDefaultClientConfigLoadingRules()

	clientCfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, nil)
	cfg, err := clientCfg.ClientConfig()
	failOnError(err)

	ns, _, err := clientCfg.Namespace()
	failOnError(err)

	dyn := dynamic.NewForConfigOrDie(cfg)

	// use the dynamic client to list the pods
	replicasets, err := dyn.Resource(schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}).Namespace(ns).List(context.TODO(), v1.ListOptions{})
	failOnError(err)

	log.Println("Pods: ")
	for _, rs := range replicasets.Items {
		log.Println(rs.GetName())
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
