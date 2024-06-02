# variables
# protoc -Iprotos --go_opt=module=microsrvhttpinproto  --go_out=. --go-grpc_opt=module=microsrvhttpinproto --go-grpc_out=.  protos/*.proto
cmp = protoc
go-o = --go_out=.
go-g-o = --go-grpc_out=.

gen-u:
	$(cmp) $(go-o) $(go-g-o) ./protos/user.proto
del-u:
	rm -rf ./userservice/user.pb.go  ./userservice/user_grpc.pb.go 

gen-o:
	$(cmp) $(go-o) $(go-g-o) ./protos/order.proto
del-o:
	rm -rf ./orderservice/order.pb.go  ./orderservice/order_grpc.pb.go 

gen-i:
	$(cmp) $(go-o) $(go-g-o) ./protos/inventory.proto
del-i:
	rm -rf ./inventoryservice/inventory.pb.go  ./inventoryservice/inventory_grpc.pb.go 
