```bash
➜  k8s-estudos git:(master) ✗ kind create cluster --config=k8s/kind.yaml --name=fullcycle
Creating cluster "fullcycle" ...
 ✓ Ensuring node image (kindest/node:v1.27.3) 🖼
 ✓ Preparing nodes 📦 📦 📦 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-fullcycle"
You can now use your cluster with:

kubectl cluster-info --context kind-fullcycle

Have a nice day! 👋
➜  k8s-estudos git:(master) ✗ kubectl cluster-info --context kind-fullcycle
Kubernetes control plane is running at https://127.0.0.1:42533
CoreDNS is running at https://127.0.0.1:42533/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
➜  k8s-estudos git:(master) ✗
```



Erro com o configmap family
```bash
➜  k8s-estudos git:(dev) ✗ kubectl port-forward svc/book-api-service 8080:80
Forwarding from 127.0.0.1:8080 -> 3000
Forwarding from [::1]:8080 -> 3000
Handling connection for 8080
Handling connection for 8080
Handling connection for 8080
Handling connection for 8080
E0112 11:53:58.945966  210729 portforward.go:409] an error occurred forwarding 8080 -> 3000: error forwarding port 3000 to pod 94190324cd2b3c6d397aa5a340d8469e6230f8395e83af9d28c34fdf1ce03b64, uid : failed to execute portforward in network namespace "/var/run/netns/cni-b2988414-5909-6eec-b4b7-317276a30a0d": failed to connect to localhost:3000 inside namespace "94190324cd2b3c6d397aa5a340d8469e6230f8395e83af9d28c34fdf1ce03b64", IPv4: dial tcp4 127.0.0.1:3000: connect: connection refused IPv6 dial tcp6 [::1]:3000: connect: connection refused
error: lost connection to pod
```

