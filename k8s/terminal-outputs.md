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