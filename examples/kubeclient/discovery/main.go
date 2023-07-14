package main

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()

	clientCfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, nil)
	cfg, err := clientCfg.ClientConfig()
	failOnError(err)

	d := memory.NewMemCacheClient(discovery.NewDiscoveryClientForConfigOrDie(cfg))

	// use discovery client to discover replica sets
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(d)
	r, err := mapper.ResourceFor(schema.GroupVersionResource{Resource: "ClusterServiceVersion"})
	failOnError(err)

	log.Println(r)
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
