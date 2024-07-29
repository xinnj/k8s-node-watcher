package main

import (
	"context"
	"log"
	"math/rand/v2"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

func ProcessNodes(clientSet *kubernetes.Clientset) {
	time.Sleep(time.Duration(rand.IntN(6)) * time.Second)

	nodes, err := clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{Limit: int64(1000)})
	if err != nil {
		log.Printf("Can't list kubernetes nodes: %s", err.Error())
	}

	for _, node := range nodes.Items {
		ProcessOneNode(clientSet, node)
	}
}

func ProcessOneNode(clientSet *kubernetes.Clientset, node v1.Node) {
	if isStuckNode(node) {
		AddTaints(clientSet, &node)
	} else {
		RemoveTaints(clientSet, &node)
	}
}

func isStuckNode(node v1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if strings.EqualFold(string(condition.Type), "Ready") &&
			strings.EqualFold(string(condition.Status), "Unknown") {
			return true
		}
	}

	return false
}

func WatchNodes(clientSet *kubernetes.Clientset) {
	log.Print("Starting kubernetes nodes watch.")

	result, err := clientSet.CoreV1().Nodes().Watch(context.TODO(), metav1.ListOptions{Limit: int64(1000)})
	if err != nil {
		log.Printf("Can't watch nodes: %s", err)
		time.Sleep(30 * time.Second)
		WatchNodes(clientSet)
	}

	if result != nil {
		for event := range result.ResultChan() {
			if event.Type == watch.Modified {
				node := event.Object.(*v1.Node)
				// log.Print("Received modified event of node: " + node.Name)
				ProcessOneNode(clientSet, *node)
			}
		}
	}

	log.Print("Closing watch channel.")
}
