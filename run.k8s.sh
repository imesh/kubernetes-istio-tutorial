pushd customer-service
./run.k8s.sh
popd

pushd delivery-service
./run.k8s.sh
popd

pushd inventory-service
./run.k8s.sh
popd

pushd invoice-service
./run.k8s.sh
popd

pushd order-service
./run.k8s.sh
popd