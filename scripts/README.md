## Usage Guide

The `kubeplugin` script is designed to retrieve resource usage statistics from Kubernetes for `nodes` and `pods` resources.

### Prerequisites

Before using the script, ensure that you have met the following requirements:

-   Install `kubectl` and configure the connection to your Kubernetes cluster.
-   Have the necessary access rights to retrieve resource statistics.
-   The `kubeplugin.sh` script is available in your environment. If you are using it from a Git repository, clone the repository to obtain the script.
    
    git clone https://github.com/vitalibit/kbot
### Setting Execution Permissions

Make sure that the `kubeplugin.sh` script has execution permissions. If you cloned it from a Git repository, run the command `chmod +x kubeplugin` to grant executable permissions.

### Usage

1.  Open a terminal or command prompt in your environment.
    
2.  Navigate to the location of the `kubeplugin` script.

    cd ./scripts
    
4.  Run the following command to execute the script:
    
    bash kubeplugin.sh RESOURCE_TYPE NAMESPACE 
    
    Replace `RESOURCE_TYPE` with the resource type you want to check. Currently, the supported values are `nodes` and `pods`.
    
    Replace `NAMESPACE` with the namespace of the `pods` resource. For the `nodes` resource, the namespace is optional and ignored.
    
5.  The result will be displayed in the terminal, showing the resource usage statistics for the specified resource type.
    

### Examples

1.  Get resource usage statistics for the `pods` resource in the `kube-system` namespace:
    
    bash kubeplugin pods kube-system 
    
2.  Get resource usage statistics for the `nodes` resource (without specifying a namespace):

    bash kubeplugin nodes
    
