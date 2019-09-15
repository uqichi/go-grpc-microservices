protoc:
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/activity/activity.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/task/task.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/user/user.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:./proto proto/project/project.proto
build:
	docker-compose build
up:
	docker-compose up
