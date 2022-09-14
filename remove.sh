pushd customer-service
./remove.sh
popd

pushd delivery-service
./remove.sh
popd

pushd inventory-service
./remove.sh
popd

pushd invoice-service
./remove.sh
popd

pushd order-service
./remove.sh
popd