
# sentinel

sentinel is a Kubernetes executor to help to manage multiple, multi-CSP backed Kubernetes clusters using standardized templates such as Kubernetes APis, Prometheus APIs and Helm commands. Using schedulers (such as cron or other cron services), users can automate the cluser management activities such as collecting Kubernetes events, status of nodes and pods, metrics & alerts of Prometheus. Also, we are working hard to provide addtional features such as custom-template with multiple steps, users can create thier own template with several Kuberntes APIs and http API calls to automate & standardize the DevOps operations. 

## Why sentinel is different than others?

sentinel consists of a manager and a agent for an each Kubernetes cluster, however, sentinel's agent is not a typical agent. No direct access from the manager to the agent API calls. The agent always polls the manager to inquire whether there is any assignment for the agent and fetch the assigned taks and execute/send the result back to the manager. Users can control how often the agent polls the manager using manager management API. 

sentinel maintains standard templates to use Kubernetes & Prometheus APIs and users just need to pick the template and cluster(s) to get what they want. It's simple and resuable. 


## What are the available templates at this moment? 

The list below is the template list sentinel is supporting as of today and "create", "delete", "apply" will be added soon. 

```
alertmanager_silences_create
alertmanager_silences_delete
alertmanager_silences_get
alertmanager_silences_list
alertmanager_silences_update
helm_get_values
helm_history
helm_install
helm_repo_add
helm_repo_list
helm_repo_update
helm_rollback
helm_uninstall
helm_upgrade
kubernetes_alertmanagerconfigs_get
kubernetes_alertmanagerconfigs_list
kubernetes_alertmanagers_get
kubernetes_alertmanagers_list
kubernetes_apiservices_get
kubernetes_apiservices_list
kubernetes_certificatesigningrequests_get
kubernetes_certificatesigningrequests_list
kubernetes_clusterrolebindings_get
kubernetes_clusterrolebindings_list
kubernetes_clusterroles_get
kubernetes_clusterroles_list
kubernetes_configmaps_get
kubernetes_configmaps_list
kubernetes_configmaps_patch
kubernetes_controllerrevisions_get
kubernetes_controllerrevisions_list
kubernetes_cronjobs_get
kubernetes_cronjobs_list
kubernetes_csidrivers_get
kubernetes_csidrivers_list
kubernetes_csinodes_get
kubernetes_csinodes_list
kubernetes_csistoragecapacities_get
kubernetes_csistoragecapacities_list
kubernetes_customresourcedefinitions_get
kubernetes_customresourcedefinitions_list
kubernetes_daemonsets_get
kubernetes_daemonsets_list
kubernetes_deployments_get
kubernetes_deployments_list
kubernetes_endpoints_get
kubernetes_endpoints_list
kubernetes_endpointslices_get
kubernetes_endpointslices_list
kubernetes_events_get
kubernetes_events_list
kubernetes_horizontalpodautoscalers_get
kubernetes_horizontalpodautoscalers_list
kubernetes_ingressclasses_get
kubernetes_ingressclasses_list
kubernetes_ingresses_get
kubernetes_ingresses_list
kubernetes_jobs_get
kubernetes_jobs_list
kubernetes_leases_get
kubernetes_leases_list
kubernetes_limitranges_get
kubernetes_limitranges_list
kubernetes_mutatingwebhookconfigurations_get
kubernetes_mutatingwebhookconfigurations_list
kubernetes_namespaces_get
kubernetes_namespaces_list
kubernetes_networkpolicies_get
kubernetes_networkpolicies_list
kubernetes_nodes_get
kubernetes_nodes_list
kubernetes_persistentvolumeclaims_get
kubernetes_persistentvolumeclaims_list
kubernetes_persistentvolumeclaims_patch
kubernetes_persistentvolumes_get
kubernetes_persistentvolumes_list
kubernetes_persistentvolumes_patch
kubernetes_poddisruptionbudgets_get
kubernetes_poddisruptionbudgets_list
kubernetes_podmonitors_get
kubernetes_podmonitors_list
kubernetes_pods_delete
kubernetes_pods_exec
kubernetes_pods_get
kubernetes_pods_list
kubernetes_podtemplates_get
kubernetes_podtemplates_list
kubernetes_priorityclasses_get
kubernetes_priorityclasses_list
kubernetes_probes_get
kubernetes_probes_list
kubernetes_prometheuses_get
kubernetes_prometheuses_list
kubernetes_prometheusrules_get
kubernetes_prometheusrules_list
kubernetes_replicasets_get
kubernetes_replicasets_list
kubernetes_replicationcontrollers_get
kubernetes_replicationcontrollers_list
kubernetes_resourcequotas_get
kubernetes_resourcequotas_list
kubernetes_rolebindings_get
kubernetes_rolebindings_list
kubernetes_roles_get
kubernetes_roles_list
kubernetes_runtimeclasses_get
kubernetes_runtimeclasses_list
kubernetes_secrets_get
kubernetes_secrets_list
kubernetes_serviceaccounts_get
kubernetes_serviceaccounts_list
kubernetes_servicemonitors_get
kubernetes_servicemonitors_list
kubernetes_services_get
kubernetes_services_list
kubernetes_statefulsets_get
kubernetes_statefulsets_list
kubernetes_storageclasses_get
kubernetes_storageclasses_list
kubernetes_thanosrulers_get
kubernetes_thanosrulers_list
kubernetes_validatingwebhookconfigurations_get
kubernetes_validatingwebhookconfigurations_list
kubernetes_volumeattachments_get
kubernetes_volumeattachments_list
prometheus_alertmanagers
prometheus_alerts
prometheus_query
prometheus_query_range
prometheus_rules
prometheus_targets
prometheus_targets/metadata
sentinel_agent_pod_rebounce
sentinel_agent_upgrade
sentinel_credential_add
sentinel_credential_get
sentinel_credential_remove
sentinel_credential_update

```

## How to install?

Use manifest files in this github to install Manager & Agent. 

### Manager

You need to have MariaDB 10.0 and above to install sentinel manager. Recommand using bitnami Mariadb helm chart to install Mariadb in your Kubernetes cluster. Once you install Mariadb and get host/port/user informatin to configre in "environment.yaml". 

```yaml
data:
  db_host: "XXX.XXX.XXX.XXX"
  db_port: "3306"
  db_scheme: "sentinel"
  db_export_path: "."
  db_server_username: "sentinel"
```

Also use BASE64 to encrypt your password and populate in the Secret in the "environment.yaml" file. Also, configure "x_auth_token" value to use the token in the header of API request. 
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: sentinel-secret
  namespace: sentinel
type: Opaque
data:
  db_server_password: 

```

Run the following commands to install the manager. 


```shell
$ kubectl apply -f application.yaml
$ kubectl apply -f environment.yaml
```

Check your installed sentinel manager in sentinel namespace. 
```shell
$ kubectl get pods -n sentinel
$ kubectl get service -n sentinel
$ kubectl get deployment -n sentinel
```

### Agent 

To install sentinel agent, you need to get cluster uuid and bear's token from sentinel manager. 

```shell
POST http://<sentinel_manager_url/manager/cluster
```
with the follwoing body - 
```json
{
  "name": "cluster name",
  "polling_option": {
    "additionalProp1": {}
  },
  "polling_limit": 10
  "summary": "cluser description"
}
```

Then, you will get uuid of your cluster. Here is the sample body of response - 

```json
    "id": 3,
    "uuid": "8a331d8d913d47e39946b32dc70e77f7",
    "name": "cluster name",
    "summary": "cluster description",
    "polling_option": {
        "additionalProp1": {}
     },
    "polling_lumit": 10   
```

Next step is to get a bearer' token with the following APIs. 

```shell
POST http://<sentinel_manager_url/manager/cluster_token
```
with the following request body - 

```json
{
  "name": "cluster name",
  "cluster_uuid": "8a331d8d913d47e39946b32dc70e77f7",
  "summary": "cluster summary"
}
```
Here is the sample response from the api. 

```json
{
    "id": 2,
    "uuid": "2e11ec0976354bfdb593af4fa6120db4",
    "name": "cluster name",
    "cluster_uuid": "8a331d8d913d47e39946b32dc70e77f7",
    "token": "8e3c93cedc0b4c57964836b4f065c84d",
    "issued_at_time": "2022-04-20T07:38:48.49479352Z",
    "expiration_time": "2022-10-21T00:00:00Z"
}
```

sentinel agent shall be installed at the Kubernetes cluster that you want to manage. 
Configure these in environment.yaml for uuid and token for agent. You can limit sentinel's roles by configuring cluster role in sa.yaml. With the following configuration, sentinel will perform only for namespace and pods. You can configure "*" for all the resources for sentinel to access & execute the commands. 

```yaml
rules:
- apiGroups: ["*"]
  resources: ["namespaces", "pods"]
  verbs: ["get", "list", "watch"]
```

Run these commands to install sentinel agent once the configuration is done.

```shell
$ kubectl apply -f environment.yaml
$ kubectl apply -f application.yaml
$ kubectl apply -f sa.yaml
```



## Make Example


swagger build 
```shell
$ make swagger
```

source build
```shell
$ make build target=manager
```

image build(manager / agent)  
```shell
$ make image base=jaehoon/sentinel target=manager
or
$ make image base=jaehoon/sentinel target=agent
```

image push
```shell
$ make push base=jaehoon/sentinel target=manager
```

## Design

how to use swagger?
```shell
$ kubectl port-forward svc/sentinel-sentinel-manager -n sentinel 8099
http://127.0.0.1:8099/swagger/index.html
```

database design

![](asset/erd.png)
