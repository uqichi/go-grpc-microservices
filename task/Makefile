build:
	docker build . -t go-grpc-microservices-task
push:
	gcloud auth configure-docker
	docker tag go-grpc-microservices-task asia.gcr.io/${GCP_PROJECT_ID}/go-grpc-microservices-task
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-task
build_push: build push
