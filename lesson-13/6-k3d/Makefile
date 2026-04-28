create:
	k3d registry create registry --port 12345
	k3d cluster create -c k3d.yml

delete:
	k3d cluster delete -c k3d.yml
	k3d registry delete k3d-registry

config:
	cat ~/.kube/config

myapp-install:
	docker build -t k3d-registry:12345/myapp:v0.1.0 ./my-app
	docker push k3d-registry:12345/myapp:v0.1.0
	kubectl apply -f ./manifest/01-deploy.yaml
	kubectl apply -f ./manifest/02-service.yaml
	kubectl apply -f ./manifest/03-ingress.yaml
	kubectl rollout restart deployment myapp

myapp-remove:
	kubectl delete -f ./manifest/01-deploy.yaml
	kubectl delete -f ./manifest/02-service.yaml
	kubectl delete -f ./manifest/03-ingress.yaml