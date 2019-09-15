build:
	docker build . -t go-grpc-microservices-project
push:
	gcloud auth configure-docker
	docker tag go-grpc-microservices-project asia.gcr.io/${GCP_PROJECT_ID}/go-grpc-microservices-project
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-project
build_push: build push
