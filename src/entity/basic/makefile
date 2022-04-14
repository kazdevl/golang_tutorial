setupDb:
	docker pull mysql:8.0

runDb:
	docker run --name sample -p 3306:3306 -e MYSQL_ROOT_PASSWORD=samplepw -e MYSQL_DATABASE=sampledb -d mysql:8.0

stopDb:
	docker stop sample

OBJ=Sample
createEntitySource:
	ent init ${OBJ}

createEntity:
	go generate ./ent