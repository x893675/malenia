## Prerequisites

1. k8s cluster with StorageClass
2. install helm `wget https://get.helm.sh/helm-v3.10.1-linux-amd64.tar.gz`
3. install dapr `wget https://github.com/dapr/cli/releases/download/v1.9.1/dapr_linux_amd64.tar.gz`
4. install dapr runtime `dapr init -k --enable-ha=true`
5. install zipkin 
   1. `kubectl create deployment zipkin --image openzipkin/zipkin`
   2. `kubectl expose deployment zipkin --type ClusterIP --port 9411`


helm repo add openebs https://openebs.github.io/charts
helm repo update
helm install openebs openebs/openebs -n openebs --create-namespace --set localprovisioner.basePath=/data
kubectl patch storageclass openebs-hostpath -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'

## Sample QuickStart

```bash
kubectl create cm gateway-config --from-file=envoy.yaml=deploy/envoy.yaml

kubectl create cm gateway-proto --from-file=proto.pb=proto/proto.pb

kubectl apply -f deploy/dapr-config.yaml


```