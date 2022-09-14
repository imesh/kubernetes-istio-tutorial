kubectl run invoice-service --image=invoice-service --image-pull-policy Never
kubectl expose pod invoice-service --name=invoice-service --type=LoadBalancer --port=8082