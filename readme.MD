## Building an Operator with the operator-sdk

```bash
# build operaotr-sdk

$ cd $GOPATH/src/github.com/operator-framework/operator-sdk
$ git checkout master
$ make dep
$ make install
```

```bash
# generate project
OPERATOR_NAME=echo-operator
GROUP=echo.redhat.com
VERSION=v1
KIND=EchoApp

$ cd $GOPATH/src/github.com/{user name}/
$ operator-sdk new ${OPERATOR_NAME} --api-version=${GROUP}/${VERSION} --kind=${KIND}
$ cd ${OPERATOR_NAME}

# build and push image
REGISTRY=docker.io
ORG=tohinkashem
TAG=v0.0.1

$ operator-sdk build ${REGISTRY}/${ORG}/${OPERATOR_NAME}:${TAG}

$ docker push ${REGISTRY}/${ORG}/${OPERATOR_NAME}:${TAG}

```


```go
# add fields to custom types

type EchoApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              EchoAppSpec   `json:"spec"`
	Status            EchoAppStatus `json:"status,omitempty"`
}

type EchoAppSpec struct {
	Size  int32  `json:"size"`
	Image string `json:"image"`
}
type EchoAppStatus struct {
	Nodes []string `json:"nodes"`
}
```

```bash
$ operator-sdk generate k8s
```

## Deploying the operator
```bash
# Create the custom resource definition
$ kubectl create -f deploy/crd.yaml

# Create the Roleand create the Role Binding
$ kubectl create -f deploy/rbac.yaml

# Deploy the operator
$ kubectl create -f deploy/operator.yaml
```

## Deploying a single node echo cluster
```bash
# Deploy the cluster
$ kubectl create -f deploy/example-single-node-cluster.yaml
```
