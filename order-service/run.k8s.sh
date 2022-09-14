kubectl run order-service --image=order-service --image-pull-policy Never
kubectl expose pod order-service --name=order-service --type=LoadBalancer --port=8081