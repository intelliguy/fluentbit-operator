kubectl create ns kubesphere-logging-system
kubectl apply -f manifests/setup
sleep 5
kubectl apply -f our-cr
