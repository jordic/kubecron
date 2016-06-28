

VERSION = r1
PREFIX = eu.gcr.io/project/repo

compile:
	CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o cron -a tmpo.io/cron

build_image:
	docker build -t $(PREFIX):$(VERSION) .

push:
	gcloud docker push $(PREFIX):$(VERSION)

create_rc:
	kubectl create -f k8.yml

update_rc:
	kubectl rolling-update tcron --image $(PREFIX):$(VERSION) --update-period=5s

update: build_image push update_rc
