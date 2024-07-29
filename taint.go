package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	taints "k8s.io/kubernetes/pkg/util/taints"
	"log"
)

var nodeTaints = [2]v1.Taint{
	{Key: "node.kubernetes.io/out-of-service",
		Value:  "nodeshutdown",
		Effect: "NoExecute"},
	{Key: "node.kubernetes.io/out-of-service",
		Value:  "nodeshutdown",
		Effect: "NoSchedule"},
}

func AddTaints(clientSet *kubernetes.Clientset, node *v1.Node) {
	var nodeUpdated bool
	for _, taint := range nodeTaints {
		var updated bool
		node, updated, _ = taints.AddOrUpdateTaint(node, &taint)
		if !nodeUpdated && updated {
			nodeUpdated = true
		}
	}

	if nodeUpdated {
		if _, err := clientSet.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{}); err != nil {
			log.Printf("Failed to add taints to node %v: %v", node.Name, err)
		} else {
			log.Printf("Added taints to node %v", node.Name)
		}
	}
}

func RemoveTaints(clientSet *kubernetes.Clientset, node *v1.Node) {
	var nodeUpdated bool
	for _, taint := range nodeTaints {
		var updated bool
		node, updated, _ = taints.RemoveTaint(node, &taint)
		if !nodeUpdated && updated {
			nodeUpdated = true
		}
	}

	if nodeUpdated {
		if _, err := clientSet.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{}); err != nil {
			log.Printf("Failed to remove taints from node %v: %v", node.Name, err)
		} else {
			log.Printf("Removed taints from node %v", node.Name)
		}
	}
}
