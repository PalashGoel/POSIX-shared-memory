build:
	env GOOS=linux GOARCH=amd64 go build -o writer writer.go
	env GOOS=linux GOARCH=amd64 go build -o read read.go

image:
	docker build -t posix-sm-writer -f Dockerfile_writer --build-arg binary=writer . 
	docker build -t posix-sm-read -f Dockerfile_read --build-arg binary=read .


load:
	minikube image load posix-sm-writer:latest
	minikube image load posix-sm-read:latest

unload:
	minikube image unload posix-sm-writer:latest
	minikube image unload posix-sm-read:latest

create:
	kubectl create -f /home/osboxes/c2c/POSIX-shared-memory/pod.yaml

delete:
	kubectl delete -f /home/osboxes/c2c/POSIX-shared-memory/pod.yaml

get:
	kubectl get po posix-sm-pod

describe:
	kubectl describe po posix-sm-pod

exec:
	kubectl exec -it posix-sm-pod -c posix-sm-writer -- sh

log:
	kubectl logs -f posix-sm-pod -c posix-sm-read
