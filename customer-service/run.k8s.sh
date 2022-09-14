kubectl run customer-service --image=customer-service --image-pull-policy Never
kubectl expose pod customer-service --name=customer-service --type=LoadBalancer --port=8080