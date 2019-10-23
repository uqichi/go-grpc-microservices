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
	docker build -f activity/Dockerfile . -t micro-activity
	gcloud auth configure-docker
	docker tag micro-activity asia.gcr.io/${GOOGLE_PROJECT}/micro-activity
	docker push asia.gcr.io/${GOOGLE_PROJECT}/micro-activity
deploy_front:
	docker build -f front/Dockerfile . -t micro-front
	gcloud auth configure-docker
	docker tag micro-front asia.gcr.io/${GOOGLE_PROJECT}/micro-front
	docker push asia.gcr.io/${GOOGLE_PROJECT}/micro-front
deploy_project:
	docker build -f project/Dockerfile . -t micro-project
	gcloud auth configure-docker
	docker tag micro-project asia.gcr.io/${GOOGLE_PROJECT}/micro-project
	docker push asia.gcr.io/${GOOGLE_PROJECT}/micro-project
deploy_task:
	docker build -f task/Dockerfile . -t micro-task
	gcloud auth configure-docker
	docker tag micro-task asia.gcr.io/${GOOGLE_PROJECT}/micro-task
	docker push asia.gcr.io/${GOOGLE_PROJECT}/micro-task
deploy_user:
	docker build -f user/Dockerfile . -t micro-user
	gcloud auth configure-docker
	docker tag micro-user asia.gcr.io/${GOOGLE_PROJECT}/micro-user
	docker push asia.gcr.io/${GOOGLE_PROJECT}/micro-user
