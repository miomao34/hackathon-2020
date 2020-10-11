NETWORK_NAME=servicenetwork

SERVER_IMAGE=serverimage
DB_IMAGE=tarantoolimage

SERVER_NAME=serverinst
DB_NAME=tarantoolinst


all:
	go build server.go

local:
	go build server.go

server:
	go build server.go

run:
	mkdir -p stored

	docker rm -f ${DB_NAME} ${SERVER_NAME} || true
	docker network rm ${NETWORK_NAME} || true
	
	docker network create --driver bridge ${NETWORK_NAME}

	docker build -t ${SERVER_IMAGE} -f docker/server/Dockerfile .
	docker build -t ${DB_IMAGE} -f docker/tarantool/Dockerfile .

	docker run --name ${DB_NAME} --network=${NETWORK_NAME} -d -p 3301:3301 -v $(PWD)/stored:/var/lib/tarantool ${DB_IMAGE}
	docker run --name ${SERVER_NAME} --network=${NETWORK_NAME} -d -p 8080:8080 ${SERVER_IMAGE}

start:
	docker ps -qa -f name=${DB_NAME} | xargs -r docker start
	docker ps -qa -f name=${SERVER_NAME} | xargs -r docker start

stop:
	docker ps -q -f name=${DB_NAME} | xargs -r docker stop
	docker ps -q -f name=${SERVER_NAME} | xargs -r docker stop
	