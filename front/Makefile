build:
	docker build . -t go-grpc-microservices-front
push:
	gcloud auth configure-docker
	docker tag go-grpc-microservices-front asia.gcr.io/${GCP_PROJECT_ID}/go-grpc-microservices-front
	docker push asia.gcr.io/${GCP_PROJECT_ID}/go-grpc-microservices-front
build_push: build push
