kubectl run delivery-service --image=delivery-service --image-pull-policy Never
kubectl expose pod delivery-service --name=delivery-service --type=LoadBalancer --port=8084