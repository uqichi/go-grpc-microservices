build:
	docker build . -t go-grpc-microservices-activity
push:
	gcloud auth configure-docker
	docker tag go-grpc-microservices-activity asia.gcr.io/stable-device-253005/go-grpc-microservices-activity
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-activity
build_push: build push
