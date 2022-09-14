pushd customer-service
./build.sh
popd

pushd delivery-service
./build.sh
popd

pushd inventory-service
./build.sh
popd

pushd invoice-service
./build.sh
popd

pushd order-service
./build.sh
popd