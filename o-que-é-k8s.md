O Kubernetes (k8s) é uma plataforma de orquestração de contêineres de código aberto que automatiza a implantação,
o dimensionamento e o gerenciamento de aplicativos em contêineres. Ele fornece um ambiente escalável e resiliente
para executar aplicativos em produção.

Alguns dos principais conceitos e contextos importantes do Kubernetes incluem:

Pods: Um pod é a menor unidade de implantação no Kubernetes. Ele contém um ou mais contêineres que são co-localizados
e compartilham recursos.

Replication Controllers/ReplicaSets: Eles garantem que um número específico de réplicas de um pod esteja em execução
em todos os momentos, garantindo a alta disponibilidade dos aplicativos.

Services: Os serviços fornecem uma maneira de expor aplicativos em execução em um conjunto de pods para outros serviços
ou usuários externos. Eles podem ser acessados por meio de um IP e porta fixos.

Deployments: Os deployments fornecem uma maneira declarativa de atualizar e gerenciar aplicativos no Kubernetes. Eles
permitem a implantação de novas versões de aplicativos com facilidade, garantindo a disponibilidade contínua.

Namespaces: Os namespaces fornecem uma maneira de dividir recursos do cluster em grupos lógicos. Eles ajudam a organizar
e isolar aplicativos e recursos em um ambiente compartilhado.

Volumes: Os volumes permitem que os contêineres acessem e armazenem dados persistentes. Eles podem ser usados para
compartilhar dados entre contêineres ou para armazenar dados além do ciclo de vida de um pod.

ConfigMaps e Secrets: ConfigMaps são usados para armazenar configurações de aplicativos como pares chave-valor, enquanto
Secrets são usados para armazenar informações sensíveis, como senhas e chaves de API.

Esses são apenas alguns dos conceitos e contextos importantes do Kubernetes. Dominar esses conceitos ajudará a garantir
que você possa implantar e gerenciar aplicativos de forma eficaz no Kubernetes em projetos reais.

De onde veio o k8s ?
O Kubernetes foi originalmente desenvolvido pelo Google e foi doado à Cloud Native Computing Foundation (CNCF) em 2015.

Borg:
O Kubernetes foi inspirado no sistema de gerenciamento de contêineres interno do Google, chamado Borg. O Borg foi projetado
para executar cargas de trabalho de produção em escala, com alta disponibilidade e eficiência. Ele foi projetado para
executar milhões de contêineres em milhares de nós.

Omega:
O Borg foi substituído pelo Omega, que foi projetado para ser mais flexível e escalável. O Omega foi projetado para
executar cargas de trabalho de produção em escala, com alta disponibilidade e eficiência. Ele foi projetado para
executar milhões de contêineres em milhares de nós.

Kubernetes:

Pontos importantes sobre o Kubernetes:

Kuberbetes é disponibilizado atravez de um conjunto de APIs:
Normalmente acessamos a API do Kubernetes por meio da ferramenta de linha de comando kubectl.
Tudo é baseado em estado. Voce configura o estado de cada objeto e o Kubernetes garante que o estado real corresponda
Kubernetes Master:
O Kubernetes Master é o cérebro do cluster. Ele é responsável por tomar decisões sobre o estado do cluster e também
por detectar e responder a eventos no cluster. O Kubernetes Master é composto por vários componentes, incluindo:

kube-apiserver: O kube-apiserver é o componente central do Kubernetes Master. Ele expõe a API do Kubernetes. Ele é
responsável por validar e configurar dados para os objetos do Kubernetes (como pods, serviços, replicação de controladores
e assim por diante) e atualizar o estado do cluster de acordo.

kube-controller-manager: O kube-controller-manager é um processo que executa vários controladores do Kubernetes. Ele
é responsável por monitorar o estado do cluster por meio do kube-apiserver e garantir que o estado real corresponda
ao estado desejado.

kube-scheduler: O kube-scheduler é responsável por atribuir pods a nós. Ele observa novos pods que não têm um nó
atribuído e seleciona um nó para eles. Ele considera requisitos de recursos, requisitos de hardware/software, afinidade
e anti-afinidade, localização de dados, limites de recursos e assim por diante.

Outros nodes:
kubelet: O kubelet é o principal agente executado em cada nó. Ele é responsável por garantir que os contêineres estejam
em execução em um pod. Ele faz isso garantindo que os contêineres em um pod estejam em execução e saudáveis. Ele também
é responsável por garantir que o estado do pod corresponda ao estado desejado.

kube-proxy: O kube-proxy é responsável por encaminhar o tráfego de rede para os pods. Ele faz isso configurando as
tabelas de regras de encaminhamento do sistema operacional. Ele também é responsável por expor serviços no cluster.

Dinamica superficial do k8s:
Cluster: Conjunto de nodes que executam aplicativos em contêineres gerenciados pelo Kubernetes. (nodes são maquinas)
Cada node possui uma quantidade de recursos (CPU, memória, armazenamento e assim por diante).

Pods: Um pod é a menor unidade de implantação no Kubernetes. Ele contém um ou mais contêineres que são co-localizados
e compartilham recursos. Os pods são executados em nodes. Um node pode executar vários pods. Os pods são efêmeros e
não mantêm um estado. Eles podem ser criados, destruídos e recriados a qualquer momento. Os pods são geralmente criados
por meio de um controlador de replicação.
recomendado usar um pod por container, mas é possivel ter mais de um container por pod.

Deployment: Os deployments fornecem uma maneira declarativa de atualizar e gerenciar aplicativos no Kubernetes. Eles provisionam
pods e os mantêm em execução. Eles fornecem recursos como implantação de novas versões de aplicativos com facilidade,
implantação de aplicativos em vários nós, dimensionamento de aplicativos e assim por diante.

ReplicaSet: Os ReplicaSets são usados para garantir que um número específico de réplicas de um pod esteja em execução
em todos os momentos. Eles são usados para garantir a alta disponibilidade dos aplicativos. Eles são usados por deployments
para garantir que o número desejado de pods esteja em execução.

Services: Os serviços fornecem uma maneira de expor aplicativos em execução em um conjunto de pods para outros serviços
ou usuários externos. Eles podem ser acessados por meio de um IP e porta fixos. Eles fornecem um ponto de extremidade
estável para acessar os pods. Eles podem ser usados para balanceamento de carga, descoberta de serviço e assim por diante.

Namespaces: Os namespaces fornecem uma maneira de dividir recursos do cluster em grupos lógicos. Eles ajudam a organizar
e isolar aplicativos e recursos em um ambiente compartilhado. Eles podem ser usados para dividir um cluster em vários
ambientes (como desenvolvimento, teste e produção).

Kind: Kind é uma ferramenta para executar clusters Kubernetes usando contêineres Docker como "nós". Kind foi principalmente
projetado para testes, CI e desenvolvimento. Ele pode ser usado para executar clusters Kubernetes em um laptop ou em
um servidor de desenvolvimento.

Minikube: Minikube é uma ferramenta que facilita a execução do Kubernetes localmente. Ele executa um único nó Kubernetes
em uma máquina virtual (VM) local para fins de desenvolvimento e teste. Ele é projetado para ser executado em laptops
e máquinas de desenvolvimento.

Kubectl: Kubectl é uma ferramenta de linha de comando para interagir com o Kubernetes. Ele é usado para implantar e
gerenciar aplicativos no Kubernetes. Ele pode ser usado para executar várias tarefas, como criar, atualizar, excluir
e obter recursos do Kubernetes.

Kubernetes Dashboard: O Kubernetes Dashboard é uma interface da Web que pode ser usada para gerenciar aplicativos em
um cluster Kubernetes. Ele fornece uma interface gráfica para visualizar aplicativos em execução no cluster, bem como
gerenciar recursos do cluster.

Criando um cluster com kind:
kind create cluster --name kind-cluster

Output e comandos que ocorreram durante os estudos do kind:

```bash
  ➜  k8s-estudos git:(master) ✗ kubectl cluster-info --context kind-kind-cluster
Kubernetes control plane is running at https://127.0.0.1:39021
CoreDNS is running at https://127.0.0.1:39021/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
➜  k8s-estudos git:(master) ✗ kubectl get nodes
NAME                         STATUS   ROLES           AGE     VERSION
kind-cluster-control-plane   Ready    control-plane   4m42s   v1.27.3
➜  k8s-estudos git:(master) ✗ kind
kind creates and manages local Kubernetes clusters using Docker container 'nodes'

Usage:
  kind [command]

Available Commands:
  build       Build one of [node-image]
  completion  Output shell completion code for the specified shell (bash, zsh or fish)
  create      Creates one of [cluster]
  delete      Deletes one of [cluster]
  export      Exports one of [kubeconfig, logs]
  get         Gets one of [clusters, nodes, kubeconfig]
  help        Help about any command
  load        Loads images into nodes
  version     Prints the kind CLI version

Flags:
  -h, --help              help for kind
      --loglevel string   DEPRECATED: see -v instead
  -q, --quiet             silence all stderr output
  -v, --verbosity int32   info log verbosity, higher value produces more output
      --version           version for kind

Use "kind [command] --help" for more information about a command.
➜  k8s-estudos git:(master) ✗ kind get clusters
kind-cluster
➜  k8s-estudos git:(master) ✗ kind delete kind-cluster
ERROR: unknown command "kind-cluster" for "kind delete"
➜  k8s-estudos git:(master) ✗ kind delete clusters kind-cluster
Deleted nodes: ["kind-cluster-control-plane"]
Deleted clusters: ["kind-cluster"]
➜  k8s-estudos git:(master) ✗
```
