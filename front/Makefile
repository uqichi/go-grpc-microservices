build:
	docker build . -t go-grpc-microservices-front
push:
	gcloud auth configure-docker
	docker tag go-grpc-microservices-front asia.gcr.io/stable-device-253005/go-grpc-microservices-front
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-front
build_push: build push
