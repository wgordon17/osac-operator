# API Reference

## Packages
- [osac.openshift.io/v1alpha1](#osacopenshiftiov1alpha1)


## osac.openshift.io/v1alpha1

Package v1alpha1 contains API Schema definitions for the osac v1alpha1 API group

### Resource Types
- [ClusterOrder](#clusterorder)
- [ClusterOrderList](#clusterorderlist)
- [ComputeInstance](#computeinstance)
- [ComputeInstanceList](#computeinstancelist)
- [PublicIP](#publicip)
- [PublicIPList](#publiciplist)
- [PublicIPPool](#publicippool)
- [PublicIPPoolList](#publicippoollist)
- [SecurityGroup](#securitygroup)
- [SecurityGroupList](#securitygrouplist)
- [Subnet](#subnet)
- [SubnetList](#subnetlist)
- [Tenant](#tenant)
- [TenantList](#tenantlist)
- [VirtualNetwork](#virtualnetwork)
- [VirtualNetworkList](#virtualnetworklist)



#### ClusterOrder



ClusterOrder is the Schema for the clusterorders API



_Appears in:_
- [ClusterOrderList](#clusterorderlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `ClusterOrder` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ClusterOrderSpec](#clusterorderspec)_ |  |  |  |
| `status` _[ClusterOrderStatus](#clusterorderstatus)_ |  |  |  |


#### ClusterOrderClusterReferenceType



ClusterOrderClusterReferenceType contains a reference to the namespace created by this ClusterOrder



_Appears in:_
- [ClusterOrderStatus](#clusterorderstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `namespace` _string_ | Namespace that contains the HostedCluster resource |  |  |
| `hostedClusterName` _string_ |  |  |  |
| `serviceAccountName` _string_ |  |  |  |
| `roleBindingName` _string_ |  |  |  |




#### ClusterOrderList



ClusterOrderList contains a list of ClusterOrder





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `ClusterOrderList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ClusterOrder](#clusterorder) array_ |  |  |  |


#### ClusterOrderPhaseType

_Underlying type:_ _string_

ClusterOrderPhaseType is a valid value for .status.phase



_Appears in:_
- [ClusterOrderStatus](#clusterorderstatus)

| Field | Description |
| --- | --- |
| `Progressing` | ClusterOrderPhaseProgressing means an update is in progress<br /> |
| `Failed` | ClusterOrderPhaseFailed means the cluster deployment or update has failed<br /> |
| `Ready` | ClusterOrderPhaseReady means the cluster and all associated resources are ready<br /> |
| `Deleting` | ClusterOrderPhaseDeleting means there has been a request to delete the ClusterOrder<br /> |


#### ClusterOrderSpec



ClusterOrderSpec defines the desired state of ClusterOrder



_Appears in:_
- [ClusterOrder](#clusterorder)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `templateID` _string_ | TemplateID is the unique identigier of the cluster template to use when creating this cluster |  | Pattern: `^[a-zA-Z_][a-zA-Z0-9._]*$` <br />Required: \{\} <br />Type: string <br /> |
| `templateParameters` _string_ | TemplateParameters is a JSON-encoded map of the parameter values for the<br />selected cluster template. |  | Optional: \{\} <br /> |
| `nodeRequests` _[NodeRequest](#noderequest) array_ | NodeRequests defines the types of nodes and number of each type of node that will be used<br />to build the cluster. This value is optional and if not provided will be filled in with template-provided<br />defaults. The selected template may limit what node types you can request. |  | Optional: \{\} <br /> |


#### ClusterOrderStatus



ClusterOrderStatus defines the observed state of ClusterOrder



_Appears in:_
- [ClusterOrder](#clusterorder)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[ClusterOrderPhaseType](#clusterorderphasetype)_ | Phase provides a single-value overview of the state of the ClusterOrder |  | Enum: [Progressing Failed Ready Deleting] <br />Optional: \{\} <br />Type: string <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the ClusterOrder |  | Optional: \{\} <br /> |
| `clusterReference` _[ClusterOrderClusterReferenceType](#clusterorderclusterreferencetype)_ | Reference to the namespace that contains the HostedCluster resource |  | Optional: \{\} <br /> |
| `nodeRequests` _[NodeRequest](#noderequest) array_ | NodeRequests reflects how many nodes are currently associated with the ClusterOrder |  |  |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs tracks the history of provision and deprovision operations<br />Ordered chronologically, with latest operations at the end<br />Limited to the last N jobs (configurable via OSAC_MAX_JOB_HISTORY, default 10) |  | Optional: \{\} <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is a hash of the current spec, used to detect spec changes<br />that require re-provisioning. |  | Optional: \{\} <br /> |


#### ComputeInstance



ComputeInstance is the Schema for the computeinstances API



_Appears in:_
- [ComputeInstanceList](#computeinstancelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `ComputeInstance` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  | Optional: \{\} <br /> |
| `spec` _[ComputeInstanceSpec](#computeinstancespec)_ | spec defines the desired state of ComputeInstance |  | Required: \{\} <br /> |
| `status` _[ComputeInstanceStatus](#computeinstancestatus)_ | status defines the observed state of ComputeInstance |  | Optional: \{\} <br /> |




#### ComputeInstanceList



ComputeInstanceList contains a list of ComputeInstance





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `ComputeInstanceList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ComputeInstance](#computeinstance) array_ |  |  |  |


#### ComputeInstancePhaseType

_Underlying type:_ _string_

ComputeInstancePhaseType is a valid value for .status.phase



_Appears in:_
- [ComputeInstanceStatus](#computeinstancestatus)

| Field | Description |
| --- | --- |
| `Starting` | ComputeInstancePhaseStarting means the compute instance is starting<br /> |
| `Running` | ComputeInstancePhaseRunning means the compute instance is running<br /> |
| `Failed` | ComputeInstancePhaseFailed means the compute instance deployment or update has failed<br /> |
| `Deleting` | ComputeInstancePhaseDeleting means there has been a request to delete the ComputeInstance<br /> |
| `Stopping` | ComputeInstancePhaseStopping means the compute instance is in the process of being stopped<br /> |
| `Stopped` | ComputeInstancePhaseStopped means the compute instance is stopped<br /> |
| `Paused` | ComputeInstancePhasePaused means the compute instance is paused<br /> |


#### ComputeInstanceSpec



ComputeInstanceSpec defines the desired state of ComputeInstance



_Appears in:_
- [ComputeInstance](#computeinstance)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `templateID` _string_ | TemplateID is the unique identifier of the compute instance template to use when creating this compute instance |  | Pattern: `^[a-zA-Z_][a-zA-Z0-9._]*$` <br />Required: \{\} <br /> |
| `templateParameters` _string_ | TemplateParameters allows passing additional template-specific parameters as JSON-encoded key-value pairs.<br />This complements the explicit fields (cores, memoryGiB, etc.) and is used for:<br />- Template-specific parameters not covered by explicit fields (e.g., exposed_ports)<br />- Custom parameters defined by specific templates |  | Optional: \{\} <br /> |
| `image` _[ImageSpec](#imagespec)_ | Image defines the VM image configuration |  | Required: \{\} <br /> |
| `cores` _integer_ | Cores is the number of CPU cores |  | Maximum: 128 <br />Minimum: 1 <br />Required: \{\} <br /> |
| `memoryGiB` _integer_ | MemoryGiB is the memory in gibibytes |  | Minimum: 1 <br />Required: \{\} <br /> |
| `bootDisk` _[DiskSpec](#diskspec)_ | BootDisk is the primary boot disk |  | Required: \{\} <br /> |
| `additionalDisks` _[DiskSpec](#diskspec) array_ | AdditionalDisks are supplementary disks |  | Optional: \{\} <br /> |
| `runStrategy` _[RunStrategyType](#runstrategytype)_ | RunStrategy controls VM running state (MUTABLE) |  | Enum: [Always Halted] <br />Required: \{\} <br /> |
| `userDataSecretRef` _[LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#localobjectreference-v1-core)_ | UserDataSecretRef references cloud-init user data |  | Optional: \{\} <br /> |
| `sshKey` _string_ | SSHKey is the SSH public key |  | Optional: \{\} <br /> |
| `subnetRef` _string_ | SubnetRef is the name of the Subnet CR in the hub cluster<br />This references the Kubernetes CR name (not the fulfillment ID) |  | Optional: \{\} <br /> |
| `restartRequestedAt` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ | RestartRequestedAt is a timestamp signal to request a VM restart (MUTABLE).<br />Set this field to the current time (usually NOW) to request a restart.<br />The controller will execute the restart if this timestamp is greater than<br />status.lastRestartedAt.<br />This is a declarative signal mechanism - the timestamp is a monotonically<br />increasing value to detect new restart requests, not a scheduled time.<br />Typically set to the current time for immediate restarts.<br />External schedulers can set this field on a schedule to implement<br />scheduled maintenance windows if needed. |  | Format: date-time <br />Optional: \{\} <br /> |


#### ComputeInstanceStatus



ComputeInstanceStatus defines the observed state of ComputeInstance.



_Appears in:_
- [ComputeInstance](#computeinstance)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[ComputeInstancePhaseType](#computeinstancephasetype)_ | Phase provides a single-value overview of the state of the ComputeInstance |  | Enum: [Starting Running Failed Deleting Stopping Stopped Paused] <br />Optional: \{\} <br />Type: string <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the ComputeInstance |  | Optional: \{\} <br /> |
| `virtualMachineReference` _[VirtualMachineReferenceType](#virtualmachinereferencetype)_ | Reference to the KubeVirt VirtualMachine CR created by this ComputeInstance |  | Optional: \{\} <br /> |
| `tenantReference` _[TenantReferenceType](#tenantreferencetype)_ | Reference to the tenant that contains the ComputeInstance resources |  | Optional: \{\} <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is the version (hash) of the desired configuration of the ComputeInstance |  | Optional: \{\} <br />Type: string <br /> |
| `lastRestartedAt` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ | LastRestartedAt records when the last restart was initiated by the controller.<br />This is set to spec.restartRequestedAt when the controller processes a restart request.<br />It will be empty if no restart has been performed yet. |  | Format: date-time <br />Optional: \{\} <br />Type: string <br /> |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs tracks the history of provision and deprovision operations<br />Ordered chronologically, with latest operations at the end<br />Limited to the last N jobs (configurable via OSAC_MAX_JOB_HISTORY, default 10) |  | Optional: \{\} <br /> |
| `ipAddress` _string_ | IPAddress is the primary IP address of the running instance, taken from the KubeVirt VirtualMachineInstance.<br />Populated when the instance is ready (phase Running). |  | Optional: \{\} <br /> |


#### DiskSpec



DiskSpec defines disk configuration



_Appears in:_
- [ComputeInstanceSpec](#computeinstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `sizeGiB` _integer_ | SizeGiB is the size of the disk in gibibytes |  | Minimum: 1 <br />Required: \{\} <br /> |


#### ImageSourceType

_Underlying type:_ _string_

ImageSourceType defines valid image source types

_Validation:_
- Enum: [registry]

_Appears in:_
- [ImageSpec](#imagespec)

| Field | Description |
| --- | --- |
| `registry` | ImageSourceTypeRegistry indicates the image is from an OCI registry<br /> |


#### ImageSpec



ImageSpec defines the VM image configuration



_Appears in:_
- [ComputeInstanceSpec](#computeinstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `sourceType` _[ImageSourceType](#imagesourcetype)_ | SourceType specifies the type of image source (currently only "registry" supported) |  | Enum: [registry] <br />Required: \{\} <br /> |
| `sourceRef` _string_ | SourceRef is the OCI image reference for the VM<br />Example: "quay.io/fedora/fedora-coreos:stable" |  | MinLength: 1 <br />Required: \{\} <br /> |


#### JobState

_Underlying type:_ _string_

JobState represents the current state of a job

_Validation:_
- Enum: [Pending Waiting Running Succeeded Failed Canceled Unknown]

_Appears in:_
- [JobStatus](#jobstatus)

| Field | Description |
| --- | --- |
| `Pending` | JobStatePending indicates the job is pending execution<br /> |
| `Waiting` | JobStateWaiting indicates the job is waiting for dependencies<br /> |
| `Running` | JobStateRunning indicates the job is currently running<br /> |
| `Succeeded` | JobStateSucceeded indicates the job completed successfully<br /> |
| `Failed` | JobStateFailed indicates the job failed<br /> |
| `Canceled` | JobStateCanceled indicates the job was canceled<br /> |
| `Unknown` | JobStateUnknown indicates the job state is unknown<br /> |


#### JobStatus



JobStatus represents the status of a provisioning or deprovisioning job



_Appears in:_
- [ClusterOrderStatus](#clusterorderstatus)
- [ComputeInstanceStatus](#computeinstancestatus)
- [PublicIPPoolStatus](#publicippoolstatus)
- [PublicIPStatus](#publicipstatus)
- [SecurityGroupStatus](#securitygroupstatus)
- [SubnetStatus](#subnetstatus)
- [VirtualNetworkStatus](#virtualnetworkstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `jobID` _string_ | JobID is the job identifier from the provisioning provider<br />For AAP Direct: job ID from AAP API response<br />For EDA: auto-incremented "eda-webhook-N" |  | Required: \{\} <br />Type: string <br /> |
| `type` _[JobType](#jobtype)_ | Type indicates the operation type |  | Enum: [provision deprovision] <br />Required: \{\} <br /> |
| `timestamp` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ | Timestamp when this job was created/triggered |  | Format: date-time <br />Required: \{\} <br />Type: string <br /> |
| `state` _[JobState](#jobstate)_ | State is the current state of the job |  | Enum: [Pending Waiting Running Succeeded Failed Canceled Unknown] <br />Required: \{\} <br /> |
| `message` _string_ | Message provides human-readable status or error information |  | Optional: \{\} <br />Type: string <br /> |
| `blockDeletionOnFailure` _boolean_ | BlockDeletionOnFailure indicates whether CR deletion should be blocked if this job fails<br />AAP Direct sets this to true to prevent orphaned cloud resources<br />EDA sets this to false as webhook handles cleanup |  | Optional: \{\} <br /> |
| `configVersion` _string_ | ConfigVersion is the DesiredConfigVersion at the time this job was triggered.<br />Used to determine retry behavior on failure: if ConfigVersion differs from<br />the current DesiredConfigVersion, a new job is triggered immediately.<br />If they match, the controller retries with exponential backoff. |  | Optional: \{\} <br /> |


#### JobType

_Underlying type:_ _string_

JobType represents the type of job operation

_Validation:_
- Enum: [provision deprovision]

_Appears in:_
- [JobStatus](#jobstatus)

| Field | Description |
| --- | --- |
| `provision` | JobTypeProvision indicates a provisioning operation<br /> |
| `deprovision` | JobTypeDeprovision indicates a deprovisioning operation<br /> |


#### NodeRequest







_Appears in:_
- [ClusterOrderSpec](#clusterorderspec)
- [ClusterOrderStatus](#clusterorderstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `resourceClass` _string_ | ResourceClass describes the type of node you are requesting |  | Required: \{\} <br /> |
| `numberOfNodes` _integer_ | NumberOfNodes describes the number of nodes you want of the given resource class |  | Required: \{\} <br /> |


#### PublicIP



PublicIP is the Schema for the publicips API



_Appears in:_
- [PublicIPList](#publiciplist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `PublicIP` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  | Optional: \{\} <br /> |
| `spec` _[PublicIPSpec](#publicipspec)_ | spec defines the desired state of PublicIP |  | Required: \{\} <br /> |
| `status` _[PublicIPStatus](#publicipstatus)_ | status defines the observed state of PublicIP |  | Optional: \{\} <br /> |




#### PublicIPList



PublicIPList contains a list of PublicIP





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `PublicIPList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[PublicIP](#publicip) array_ |  |  |  |


#### PublicIPPhaseType

_Underlying type:_ _string_

PublicIPPhaseType is a valid value for .status.phase



_Appears in:_
- [PublicIPStatus](#publicipstatus)

| Field | Description |
| --- | --- |
| `Progressing` | PublicIPPhaseProgressing means an update is in progress<br /> |
| `Failed` | PublicIPPhaseFailed means the IP provisioning has failed<br /> |
| `Ready` | PublicIPPhaseReady means the IP and all associated resources are ready<br /> |
| `Deleting` | PublicIPPhaseDeleting means there has been a request to delete the PublicIP<br /> |


#### PublicIPPool



PublicIPPool is the Schema for the publicippools API



_Appears in:_
- [PublicIPPoolList](#publicippoollist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `PublicIPPool` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  | Optional: \{\} <br /> |
| `spec` _[PublicIPPoolSpec](#publicippoolspec)_ | spec defines the desired state of PublicIPPool |  | Required: \{\} <br /> |
| `status` _[PublicIPPoolStatus](#publicippoolstatus)_ | status defines the observed state of PublicIPPool |  | Optional: \{\} <br /> |




#### PublicIPPoolList



PublicIPPoolList contains a list of PublicIPPool





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `PublicIPPoolList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[PublicIPPool](#publicippool) array_ |  |  |  |


#### PublicIPPoolPhaseType

_Underlying type:_ _string_

PublicIPPoolPhaseType is a valid value for .status.phase



_Appears in:_
- [PublicIPPoolStatus](#publicippoolstatus)

| Field | Description |
| --- | --- |
| `Progressing` | PublicIPPoolPhaseProgressing means an update is in progress<br /> |
| `Failed` | PublicIPPoolPhaseFailed means the pool provisioning has failed<br /> |
| `Ready` | PublicIPPoolPhaseReady means the pool and all associated resources are ready<br /> |
| `Deleting` | PublicIPPoolPhaseDeleting means there has been a request to delete the PublicIPPool<br /> |


#### PublicIPPoolSpec



PublicIPPoolSpec defines the desired state of PublicIPPool



_Appears in:_
- [PublicIPPool](#publicippool)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cidrs` _string array_ | CIDRs is the list of CIDR blocks for this pool. All CIDRs must match the declared IPFamily. |  | MinItems: 1 <br />Required: \{\} <br /> |
| `ipFamily` _string_ | IPFamily indicates the IP address family for this pool (IPv4 or IPv6) |  | Enum: [IPv4 IPv6] <br />Required: \{\} <br />Type: string <br /> |
| `implementationStrategy` _string_ | ImplementationStrategy determines the backend used to advertise IPs (e.g., metallb-l2).<br />Defaults to metallb-l2 for v7.0. |  | Enum: [metallb-l2] <br />Optional: \{\} <br />Type: string <br /> |


#### PublicIPPoolStatus



PublicIPPoolStatus defines the observed state of PublicIPPool



_Appears in:_
- [PublicIPPool](#publicippool)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[PublicIPPoolPhaseType](#publicippoolphasetype)_ | Phase provides a single-value overview of the state of the PublicIPPool |  | Enum: [Progressing Failed Ready Deleting] <br />Optional: \{\} <br />Type: string <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is a hash of the spec, used to detect spec changes and control retry behavior. |  | Optional: \{\} <br /> |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs holds an array of JobStatus tracking provisioning and deprovisioning operations |  | Optional: \{\} <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the PublicIPPool |  | Optional: \{\} <br /> |
| `total` _integer_ | Total is the total number of usable IP addresses across all CIDRs in this pool.<br />Uses int64 to accommodate large IPv6 CIDR ranges. |  | Optional: \{\} <br /> |
| `allocated` _integer_ | Allocated is the number of IPs currently allocated from the pool. |  | Optional: \{\} <br /> |
| `available` _integer_ | Available is the number of IPs available for allocation. |  | Optional: \{\} <br /> |


#### PublicIPSpec



PublicIPSpec defines the desired state of PublicIP



_Appears in:_
- [PublicIP](#publicip)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `pool` _string_ | Pool is the name of the PublicIPPool this IP is allocated from.<br />This field is immutable after creation. |  | MinLength: 1 <br />Required: \{\} <br /> |
| `computeInstance` _string_ | ComputeInstance is the optional name of the ComputeInstance this IP is attached to.<br />Setting this field triggers attachment of the IP to the referenced instance. |  | MinLength: 1 <br />Optional: \{\} <br /> |


#### PublicIPStateType

_Underlying type:_ _string_

PublicIPStateType is a valid value for .status.state



_Appears in:_
- [PublicIPStatus](#publicipstatus)

| Field | Description |
| --- | --- |
| `Pending` | PublicIPStatePending means the IP allocation is pending<br /> |
| `Allocated` | PublicIPStateAllocated means the IP has been allocated from the pool<br /> |
| `Attached` | PublicIPStateAttached means the IP is attached to a ComputeInstance<br /> |
| `Releasing` | PublicIPStateReleasing means the IP is being released back to the pool<br /> |
| `Failed` | PublicIPStateFailed means provisioning or release failed<br /> |


#### PublicIPStatus



PublicIPStatus defines the observed state of PublicIP



_Appears in:_
- [PublicIP](#publicip)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[PublicIPPhaseType](#publicipphasetype)_ | Phase provides a single-value overview of the state of the PublicIP |  | Enum: [Progressing Failed Ready Deleting] <br />Optional: \{\} <br />Type: string <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is a hash of the spec, used to detect spec changes and control retry behavior. |  | Optional: \{\} <br /> |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs holds an array of JobStatus tracking provisioning and deprovisioning operations |  | Optional: \{\} <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the PublicIP |  | Optional: \{\} <br /> |
| `address` _string_ | Address is the allocated public IP address |  | Optional: \{\} <br /> |
| `state` _[PublicIPStateType](#publicipstatetype)_ | State tracks the attachment lifecycle of the PublicIP |  | Enum: [Pending Allocated Attached Releasing Failed] <br />Optional: \{\} <br />Type: string <br /> |


#### RunStrategyType

_Underlying type:_ _string_

RunStrategyType defines valid VM run strategies

_Validation:_
- Enum: [Always Halted]

_Appears in:_
- [ComputeInstanceSpec](#computeinstancespec)

| Field | Description |
| --- | --- |
| `Always` | RunStrategyAlways means the VM should always be running<br /> |
| `Halted` | RunStrategyHalted means the VM should be stopped<br /> |


#### SecurityGroup



SecurityGroup is the Schema for the securitygroups API



_Appears in:_
- [SecurityGroupList](#securitygrouplist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `SecurityGroup` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  | Optional: \{\} <br /> |
| `spec` _[SecurityGroupSpec](#securitygroupspec)_ | spec defines the desired state of SecurityGroup |  | Required: \{\} <br /> |
| `status` _[SecurityGroupStatus](#securitygroupstatus)_ | status defines the observed state of SecurityGroup |  | Optional: \{\} <br /> |


#### SecurityGroupList



SecurityGroupList contains a list of SecurityGroup





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `SecurityGroupList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[SecurityGroup](#securitygroup) array_ |  |  |  |


#### SecurityGroupPhaseType

_Underlying type:_ _string_

SecurityGroupPhaseType is a valid value for .status.phase

_Validation:_
- Enum: [Progressing Ready Failed Deleting]

_Appears in:_
- [SecurityGroupStatus](#securitygroupstatus)

| Field | Description |
| --- | --- |
| `Progressing` | SecurityGroupPhaseProgressing means an update is in progress<br /> |
| `Ready` | SecurityGroupPhaseReady means the security group and all associated resources are ready<br /> |
| `Failed` | SecurityGroupPhaseFailed means the security group provisioning has failed<br /> |
| `Deleting` | SecurityGroupPhaseDeleting means there has been a request to delete the SecurityGroup<br /> |


#### SecurityGroupProtocol

_Underlying type:_ _string_

SecurityGroupProtocol represents network protocol types

_Validation:_
- Enum: [tcp udp icmp all]

_Appears in:_
- [SecurityRule](#securityrule)

| Field | Description |
| --- | --- |
| `tcp` | SecurityGroupProtocolTCP represents TCP protocol<br /> |
| `udp` | SecurityGroupProtocolUDP represents UDP protocol<br /> |
| `icmp` | SecurityGroupProtocolICMP represents ICMP protocol<br /> |
| `all` | SecurityGroupProtocolAll represents all protocols<br /> |


#### SecurityGroupSpec



SecurityGroupSpec defines the desired state of SecurityGroup



_Appears in:_
- [SecurityGroup](#securitygroup)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `virtualNetwork` _string_ | VirtualNetwork is the ID of the parent VirtualNetwork |  | Required: \{\} <br />Type: string <br /> |
| `ingressRules` _[SecurityRule](#securityrule) array_ | IngressRules defines the ingress security rules |  | Optional: \{\} <br /> |
| `egressRules` _[SecurityRule](#securityrule) array_ | EgressRules defines the egress security rules |  | Optional: \{\} <br /> |


#### SecurityGroupStatus



SecurityGroupStatus defines the observed state of SecurityGroup



_Appears in:_
- [SecurityGroup](#securitygroup)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[SecurityGroupPhaseType](#securitygroupphasetype)_ | Phase provides a single-value overview of the state of the SecurityGroup |  | Enum: [Progressing Ready Failed Deleting] <br />Optional: \{\} <br />Type: string <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is a hash of the spec, used to detect spec changes and control retry behavior. |  | Optional: \{\} <br /> |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs holds an array of JobStatus tracking provisioning and deprovisioning operations |  | Optional: \{\} <br /> |
| `backendSecurityGroupId` _string_ | BackendSecurityGroupID stores provider-specific security group identifier |  | Optional: \{\} <br />Type: string <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the SecurityGroup |  | Optional: \{\} <br /> |


#### SecurityRule



SecurityRule defines a single security rule for ingress or egress traffic



_Appears in:_
- [SecurityGroupSpec](#securitygroupspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `protocol` _[SecurityGroupProtocol](#securitygroupprotocol)_ | Protocol specifies the network protocol |  | Enum: [tcp udp icmp all] <br />Required: \{\} <br />Type: string <br /> |
| `portFrom` _integer_ | PortFrom specifies the start of the port range |  | Maximum: 65535 <br />Minimum: 1 <br />Optional: \{\} <br /> |
| `portTo` _integer_ | PortTo specifies the end of the port range |  | Maximum: 65535 <br />Minimum: 1 <br />Optional: \{\} <br /> |
| `sourceCidr` _string_ | SourceCIDR specifies the source CIDR block for this rule |  | Optional: \{\} <br />Type: string <br /> |
| `destinationCidr` _string_ | DestinationCIDR specifies the destination CIDR block for this rule |  | Optional: \{\} <br />Type: string <br /> |


#### Subnet



Subnet is the Schema for the subnets API



_Appears in:_
- [SubnetList](#subnetlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `Subnet` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  | Optional: \{\} <br /> |
| `spec` _[SubnetSpec](#subnetspec)_ | spec defines the desired state of Subnet |  | Required: \{\} <br /> |
| `status` _[SubnetStatus](#subnetstatus)_ | status defines the observed state of Subnet |  | Optional: \{\} <br /> |




#### SubnetList



SubnetList contains a list of Subnet





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `SubnetList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Subnet](#subnet) array_ |  |  |  |


#### SubnetPhaseType

_Underlying type:_ _string_

SubnetPhaseType is a valid value for .status.phase



_Appears in:_
- [SubnetStatus](#subnetstatus)

| Field | Description |
| --- | --- |
| `Progressing` | SubnetPhaseProgressing means an update is in progress<br /> |
| `Failed` | SubnetPhaseFailed means the subnet provisioning has failed<br /> |
| `Ready` | SubnetPhaseReady means the subnet and all associated resources are ready<br /> |
| `Deleting` | SubnetPhaseDeleting means there has been a request to delete the Subnet<br /> |


#### SubnetSpec



SubnetSpec defines the desired state of Subnet



_Appears in:_
- [Subnet](#subnet)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `virtualNetwork` _string_ | VirtualNetwork is the ID of the parent VirtualNetwork |  | Required: \{\} <br />Type: string <br /> |
| `ipv4Cidr` _string_ | IPv4CIDR is the IPv4 CIDR block for this subnet |  | Optional: \{\} <br />Type: string <br /> |
| `ipv6Cidr` _string_ | IPv6CIDR is the IPv6 CIDR block for this subnet |  | Optional: \{\} <br />Type: string <br /> |


#### SubnetStatus



SubnetStatus defines the observed state of Subnet



_Appears in:_
- [Subnet](#subnet)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[SubnetPhaseType](#subnetphasetype)_ | Phase provides a single-value overview of the state of the Subnet |  | Enum: [Progressing Failed Ready Deleting] <br />Optional: \{\} <br />Type: string <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is a hash of the spec, used to detect spec changes and control retry behavior. |  | Optional: \{\} <br /> |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs holds an array of JobStatus tracking provisioning and deprovisioning operations |  | Optional: \{\} <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the Subnet |  | Optional: \{\} <br /> |
| `backendNetworkId` _string_ | BackendNetworkID stores provider-specific network identifier |  | Optional: \{\} <br /> |


#### Tenant



Tenant is the Schema for the tenants API.



_Appears in:_
- [TenantList](#tenantlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `Tenant` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[TenantSpec](#tenantspec)_ |  |  |  |
| `status` _[TenantStatus](#tenantstatus)_ |  |  |  |




#### TenantList



TenantList contains a list of Tenant.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `TenantList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Tenant](#tenant) array_ |  |  |  |


#### TenantPhaseType

_Underlying type:_ _string_





_Appears in:_
- [TenantStatus](#tenantstatus)

| Field | Description |
| --- | --- |
| `Progressing` |  |
| `Ready` |  |


#### TenantReferenceType



TenantReferenceType contains a reference to the tenant that contains the ComputeInstance resources



_Appears in:_
- [ComputeInstanceStatus](#computeinstancestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ | Name of the tenant |  |  |
| `namespace` _string_ | Namespace of the tenant |  |  |


#### TenantSpec



TenantSpec defines the desired state of Tenant.



_Appears in:_
- [Tenant](#tenant)



#### TenantStatus



TenantStatus defines the observed state of Tenant.



_Appears in:_
- [Tenant](#tenant)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[TenantPhaseType](#tenantphasetype)_ | Phase is the phase of the tenant |  |  |
| `namespace` _string_ | Namespace is the namespace allocated to the tenant on the target cluster |  |  |
| `storageClass` _string_ | StorageClass is the StorageClass allocated to the tenant on the target cluster |  |  |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the Tenant |  | Optional: \{\} <br /> |


#### VirtualMachineReferenceType



VirtualMachineReferenceType contains a reference to the KubeVirt VirtualMachine CR created by this ComputeInstance



_Appears in:_
- [ComputeInstanceStatus](#computeinstancestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `namespace` _string_ | Namespace that contains the VirtualMachine resources |  |  |
| `kubeVirtVirtualMachineName` _string_ |  |  |  |


#### VirtualNetwork



VirtualNetwork is the Schema for the virtualnetworks API



_Appears in:_
- [VirtualNetworkList](#virtualnetworklist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `VirtualNetwork` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  | Optional: \{\} <br /> |
| `spec` _[VirtualNetworkSpec](#virtualnetworkspec)_ | spec defines the desired state of VirtualNetwork |  | Required: \{\} <br /> |
| `status` _[VirtualNetworkStatus](#virtualnetworkstatus)_ | status defines the observed state of VirtualNetwork |  | Optional: \{\} <br /> |


#### VirtualNetworkList



VirtualNetworkList contains a list of VirtualNetwork





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `osac.openshift.io/v1alpha1` | | |
| `kind` _string_ | `VirtualNetworkList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  | Optional: \{\} <br /> |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  | Optional: \{\} <br /> |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[VirtualNetwork](#virtualnetwork) array_ |  |  |  |


#### VirtualNetworkPhaseType

_Underlying type:_ _string_

VirtualNetworkPhaseType is a valid value for .status.phase

_Validation:_
- Enum: [Progressing Ready Failed Deleting]

_Appears in:_
- [VirtualNetworkStatus](#virtualnetworkstatus)

| Field | Description |
| --- | --- |
| `Progressing` | VirtualNetworkPhaseProgressing means an update is in progress<br /> |
| `Ready` | VirtualNetworkPhaseReady means the virtual network and all associated resources are ready<br /> |
| `Failed` | VirtualNetworkPhaseFailed means the virtual network provisioning has failed<br /> |
| `Deleting` | VirtualNetworkPhaseDeleting means there has been a request to delete the VirtualNetwork<br /> |


#### VirtualNetworkSpec



VirtualNetworkSpec defines the desired state of VirtualNetwork



_Appears in:_
- [VirtualNetwork](#virtualnetwork)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `region` _string_ | Region is the cloud region where this VirtualNetwork will be provisioned |  | Required: \{\} <br />Type: string <br /> |
| `ipv4Cidr` _string_ | IPv4CIDR is the IPv4 CIDR block for this virtual network |  | Optional: \{\} <br />Type: string <br /> |
| `ipv6Cidr` _string_ | IPv6CIDR is the IPv6 CIDR block for this virtual network |  | Optional: \{\} <br />Type: string <br /> |
| `networkClass` _string_ | NetworkClass is the name of the NetworkClass that defines implementation strategy |  | Required: \{\} <br />Type: string <br /> |
| `implementationStrategy` _string_ | ImplementationStrategy determines the underlying network backend and Ansible role to use.<br />This value is derived from the NetworkClass at creation time and stored here for direct<br />access by controllers and provisioning systems. |  | Optional: \{\} <br />Type: string <br /> |


#### VirtualNetworkStatus



VirtualNetworkStatus defines the observed state of VirtualNetwork



_Appears in:_
- [VirtualNetwork](#virtualnetwork)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `phase` _[VirtualNetworkPhaseType](#virtualnetworkphasetype)_ | Phase provides a single-value overview of the state of the VirtualNetwork |  | Enum: [Progressing Ready Failed Deleting] <br />Optional: \{\} <br />Type: string <br /> |
| `desiredConfigVersion` _string_ | DesiredConfigVersion is a hash of the spec, used to detect spec changes and control retry behavior. |  | Optional: \{\} <br /> |
| `jobs` _[JobStatus](#jobstatus) array_ | Jobs holds an array of JobStatus tracking provisioning and deprovisioning operations |  | Optional: \{\} <br /> |
| `backendNetworkId` _string_ | BackendNetworkID stores provider-specific network identifier |  | Optional: \{\} <br />Type: string <br /> |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions holds an array of metav1.Condition that describe the state of the VirtualNetwork |  | Optional: \{\} <br /> |
