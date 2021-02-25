# Kubernetes Zerotier bridge 

### *TL;DR*
A Zerotier gateway to access your non-public k8s services thru ZT subnet 


### Kubernetes

## Prerequisites

* A consul cluster with agents listening on the host (to be made optional)

## Helm chart to deploy a DaemonSet

NOTE: This is way out of date. just checkout the repo and run helm directly from it for now.

`Helm repo add kubernetes-zerotier-bridge https://leunamnauj.github.io/kubernetes-zerotier-bridge/`

`helm repo update`

`helm install --name kubernetes-zerotier-bridge kubernetes-zerotier-bridge/kubernetes-zerotier-bridge`

## Single Deployment

NOTE: This may not work anymore, untested.

Since this docker image expects the subnetIDs as an env variable you need to use something like this
```
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: zerotier-network
data:
  NETWORK_ID: << your subnetid >>
  ZTAUTHTOKEN: << your token >>
  AUTOJOIN: true
---
apiVersion: v1
kind: Pod
metadata:
  name: kubernetes-zerotier-bridge
spec:
  containers:
    - name: ubernetes-zerotier-bridge
      image: << your registry >>
      env:
      - name: NETWORK_ID
        valueFrom:
          configMapKeyRef:
            name: zerotier-network
            key: NETWORK_ID
      - name: ZTHOSTNAME
        valueFrom:
          configMapKeyRef:
            name: zerotier-network
            key: ZTHOSTNAME 
      - name: ZTAUTHTOKEN
        valueFrom:
          configMapKeyRef:
            name: zerotier-network
            key: ZTAUTHTOKEN 
      - name: AUTOJOIN
        valueFrom:
          configMapKeyRef:
            name: zerotier-network
            key: AUTOJOIN 
      securityContext:
          privileged: true
          capabilities:
            add:
            - NET_ADMIN
            - SYS_ADMIN
            - CAP_NET_ADMIN
        volumeMounts:
        - name: dev-net-tun
          mountPath: /dev/net/tun

```
**Important:** Be aware of `securityContext` and `dev-net-tun` volume

## Zerotier level config
In order to route traffic to this POD have to add the proper rule on ZT Managed Routes section, to accomplish that you have to know the ZT address assigned to the pod and your Service and/or PODs subnet.





## Local Run

NOTE: This is also probably broken.

Running this locally will let you test your ZT connection and also use it without install ZT at all

### Usage

Modify docker compose file accordly.

  - `NETWORK_ID` networkID.
  - `ZTAUTHTOKEN` Your network token, required to perform auto join and set hostname.
  - `AUTOJOIN` Automatically accept new host.
  - `ZTHOSTNAME` Hostname to identify this client. If not provided will keep it blank.
```
docker-compose up
```




## Inspired on

* https://github.com/henrist/zerotier-one-docker
* https://github.com/crocandr/docker-zerotier