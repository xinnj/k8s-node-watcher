# k8s-node-watcher

When a node is shutdown but not detected by kubelet's Node Shutdown Manager, the pods that are part of a StatefulSet 
will be stuck in terminating status on the shutdown node and cannot move to a new running node. 

From v1.26 [beta] / v1.28 [stable], Kubernetes introduces a new feature named "non-graceful node shutdown".
A user can manually add the taint `node.kubernetes.io/out-of-service` with either `NoExecute` or `NoSchedule` effect to 
a Node marking it out-of-service. The pods on the node will be forcefully deleted if there are no matching tolerations 
on it and volume detach operations for the pods terminating on the node will happen immediately. 
This allows the Pods on the out-of-service node to recover quickly on a different node.
Please find the detail from this [post](https://kubernetes.io/docs/concepts/cluster-administration/node-shutdown/#non-graceful-node-shutdown).   

This application will add (remove) taint `node.kubernetes.io/out-of-service` to (from) a node automatically by watching node status changing.

# Installation
Download the helm chart and deploy to a k8s cluster.