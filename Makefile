protoc:
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/activity/activity.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/task/task.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/user/user.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/project/project.proto
build:
	docker-compose build
up:
	docker-compose up
deploy_activity:
	docker build -f activity/Dockerfile . -t go-grpc-microservices-activity
	gcloud auth configure-docker
	docker tag go-grpc-microservices-activity asia.gcr.io/stable-device-253005/go-grpc-microservices-activity
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-activity
deploy_front:
	docker build -f front/Dockerfile . -t go-grpc-microservices-front
	gcloud auth configure-docker
	docker tag go-grpc-microservices-front asia.gcr.io/stable-device-253005/go-grpc-microservices-front
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-front
deploy_task:
	docker build -f task/Dockerfile . -t go-grpc-microservices-task
	gcloud auth configure-docker
	docker tag go-grpc-microservices-task asia.gcr.io/stable-device-253005/go-grpc-microservices-task
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-task
deploy_user:
	docker build -f user/Dockerfile . -t go-grpc-microservices-user
	gcloud auth configure-docker
	docker tag go-grpc-microservices-user asia.gcr.io/stable-device-253005/go-grpc-microservices-user
	docker push asia.gcr.io/stable-device-253005/go-grpc-microservices-user
