FROM ubuntu

COPY k8s-node-watcher /usr/bin/
CMD ["k8s-node-watcher"]