kubectl run inventory-service --image=inventory-service --image-pull-policy Never
kubectl expose pod inventory-service --name=inventory-service --type=LoadBalancer --port=8083