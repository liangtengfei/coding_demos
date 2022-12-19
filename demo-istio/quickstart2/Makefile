build:
	docker build -t quickstart2/homepage:v1 -f Dockerfile .
buildv2:
	docker build -t quickstart2/homepage:v2 -f Dockerfile .

serivce:
	kubectl apply -f .\deploys\service.yaml
gateway:
	kubectl apply -f .\deploys\istio_gateway.yaml

.PHONY: build buildv2