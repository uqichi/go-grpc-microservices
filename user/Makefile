build:
	docker build . -t go-grpc-microservices-user
push:
	gcloud auth configure-docker
	docker tag go-grpc-microservices-user asia.gcr.io/stable-device-253005/go-grpc-microservices-user
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-user
build_push: build push
