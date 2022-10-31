## Prerequisites

### Install Dev

- install go
- brew install protobuf
- brew install envoy
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
- go install github.com/envoyproxy/protoc-gen-validate@v0.16.3
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

### Install k8s and dapr runtime

1. k8s cluster with StorageClass
2. install helm `wget https://get.helm.sh/helm-v3.10.1-linux-amd64.tar.gz`
3. install dapr `wget https://github.com/dapr/cli/releases/download/v1.9.1/dapr_linux_amd64.tar.gz`
4. install dapr runtime `dapr init -k --enable-ha=true`
5. install zipkin
   1. `kubectl create deployment zipkin --image openzipkin/zipkin`
   2. `kubectl expose deployment zipkin --type ClusterIP --port 9411`

### Install Redis

1. `helm repo add bitnami https://charts.bitnami.com/bitnami`
2. `helm repo update`
3. `helm install redis bitnami/redis --set image.tag=6.2`

### Install openebs (OPTIONAL)

1. `helm repo add openebs https://openebs.github.io/charts`
2. `helm repo update`
3. `mkdir -pv /data`
4. `helm install openebs openebs/openebs -n openebs --create-namespace --set localprovisioner.basePath=/data`
5. `kubectl patch storageclass openebs-hostpath -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'`

## QuickStart

### Install

1. `kubectl create cm gateway-config --from-file=envoy.yaml=deploy/envoy.yaml`
2. `kubectl create cm gateway-proto --from-file=proto.pb=proto/proto.pb`
3. `kubectl apply -f deploy/`

### Test

All Request will be deny except with Bearer Token `hanamichi` or `spike`

1. call dapr HTTP `echo` handler without token

```
Request:
curl -X POST \
  '192.168.234.3:31484/s/echo' \
  --header 'Accept: */*' \
  --header 'Authorization: Bearer spike2' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "value2"
}'

Respond:
http status: 401
body: PERMISSION_DENIED
```

2. call dapr HTTP `echo` handler without token

```
Request:
curl -X POST \
  '192.168.234.3:31484/s/echo' \
  --header 'Accept: */*' \
  --header 'Authorization: Bearer spike' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "value2"
}'

Respond:
http status: 200
body:
{
  "name": "value2"
}
```