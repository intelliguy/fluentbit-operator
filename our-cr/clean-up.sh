kubectl delete -f our-cr
kubectl delete -f manifests/setup
for svr in ct1l ctl2 ctl3 com1 com2
do
	ssh $svr "sudo docker rmi siim/fluentbit-operator:v0.0.1"
done
kubectl delete ns kubesphere-logging-system
