package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"log"
)

func main() {
	var outAuth bool
	flag.BoolVar(&outAuth, "out", false, "Auth from outside of cluster")
	flag.Parse()

	var clientSet *kubernetes.Clientset
	var err error

	if outAuth {
		log.Print("Using outside of cluster")
		clientSet, err = AuthenticateOutOfCluster()
	} else {
		log.Print("Using inside of cluster")
		clientSet, err = AuthenticateInCluster()
	}
	if err != nil {
		panic("Failed to authenticate on kubernetes: " + err.Error())
	}

	ProcessNodes(clientSet)

	for {
		WatchNodes(clientSet)
	}
}
