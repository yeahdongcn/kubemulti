# kubemulti

A `kubectl` plugin to query multiple namespace at the same time.

```
$ kubemulti -n cdi -n default get pods
NAMESPACE   NAME                                                   READY   STATUS    RESTARTS   AGE
cdi         cdi-apiserver-b9984b66c-d8l4z                          1/1     Running   0          76m
cdi         cdi-deployment-695bb6698f-dsllh                        1/1     Running   0          76m
cdi         cdi-operator-66d96fbb8b-t4565                          1/1     Running   0          83m
cdi         cdi-uploadproxy-6bd6cb85d7-fm584                       1/1     Running   0          76m
default     nfs-nfs-subdir-external-provisioner-546b74ccb7-9sn5w   1/1     Running   0          82m
```