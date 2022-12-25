serv-api:
	go run .\api\bookstore.go -f .\api\etc\bookstore-api.yaml

serv-rpc1:
	go run .\rpc\add\add.go -f .\rpc\add\etc\add.yaml
serv-rpc2:
	go run .\rpc\check\check.go -f .\rpc\check\etc\check.yaml
serv-rpc3:
	go run .\rpc\borrow\borrow.go -f .\rpc\borrow\etc\borrow.yaml

.PHONY: serv-api serv-rpc1 serv-rpc2 serv-rpc3