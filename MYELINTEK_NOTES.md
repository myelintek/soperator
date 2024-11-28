# NOTES by MyelinTek

- base:
  - tag: `1.15.3`

## Pre-requisites

1. Kubernetes cluster available.
1. NFS service available, needed in creating some PVs.
1. At least a CPU-only node out of GPU Operator and at least a GPU node within GPU Operator.
1. `local-path` storage class is available, needed in creating worker PVs.
1. `cr.myelintek.com/soperator/slurm-operator:1.15.3-p3` is available, which has be built as follows
   ```bash
   docker build -t cr.myelintek.com/soperator/slurm-operator:1.15.3-p3 .
   ```

## Limitations

Inherited from soperator:

1. No AMD GPU support.
1. Single-partition clusters.
1. The list of software versions we currently support is quite short:
   - Linux: Ubuntu 20.04 and 22.04.
   - Slurm: versions 23.11.6 and 24.05.3.
   - CUDA: version 12.2.2.
   - Kubernetes: >= 1.28.
   - Versions of some preinstalled software packages can't be changed.

Specific to our deployment:

1. No accounting support.
1. Single-controller only.

## Deployment

1. Install the soperator.
   ```bash
   cd $repo/helm/soperator
   helm -n soperator install --create-namespace -f values.yaml soperator .
   ```
1. Prepare resources needed by the slurm cluster.
   - The slurm cluster will be installed in the `slurm1` namespace. To install the slurm cluster in another namespace:
     - Add the `-n $another_namespace` argument in ALL preceding kubectl and helm commands.
     - Change the value of `clusterName` in `helm/slurm-cluster/values.yaml`.
   ```bash
   cd $repo/deployment
   kubectl apply -f cluster.yaml
   kubectl apply -f secrets.yaml
   ```
1. Install the slurm cluster.
   ```bash
   cd $repo/helm/slurm-cluster
   helm -n slurm1 install -f values.yaml slurm .
   ```
   > It takes long time to pull images. All the pods should be ready or completed after a while.

To log in to the Slurm service:

1. Create port forwarding (due to VPN restrictions).
   ```bash
   kubectl -n slurm1 port-forward svc/slurm1-login-svc 8022:22
   ```
1. Access with SSH on the host machine.
   ```console
   $ ssh -p 8022 -i $repo/deployment/id_soperator root@localhost
   Welcome to Ubuntu 22.04.5 LTS (GNU/Linux 5.15.0-91-generic x86_64)
   
    * Documentation:  https://help.ubuntu.com
    * Management:     https://landscape.canonical.com
    * Support:        https://ubuntu.com/pro
   
   This system has been minimized by removing packages and content that are
   not required on a system that users do not log into.
   
   To restore this content, you can run the 'unminimize' command.
   
   The programs included with the Ubuntu system are free software;
   the exact distribution terms for each program are described in the
   individual files in /usr/share/doc/*/copyright.
   
   Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
   applicable law.
   
   root@login-0:~# sinfo
   PARTITION AVAIL  TIMELIMIT  NODES  STATE NODELIST
   main*        up   infinite      2   idle worker-[0-1]
   root@login-0:~#
   ```

## Changes

1. deployment
   1. helm/slurm-cluster/values.yaml
      - Create NFS-based PVs/PVCs needed by jail and controller-spool in slurm-cluster (we need PVCs across nodes).
   1. deployment/secrets.yaml
      - Create the `slurm1-slurmdbd-configs` secret manually, which soperator should create but it doesn't.
        > https://github.com/nebius/soperator/issues/203
   1. deployment/id\_soperator and deployment/id\_soperator.pub
      - A self-generated key pair for accessing the login service.
1. soperator
   1. helm/soperator/templates/deployment.yaml
      - Add node selector.
   1. helm/soperator/values.yaml
      - Replace the operator image from `cr.eu-north1.nebius.cloud/soperator/slurm-operator:1.15.3` to `cr.myelintek.com/soperator/slurm-operator:1.15.3-p3`.
        - We built an operator image to fix bugs (see the next item).
          > https://github.com/nebius/soperator/issues/204
   1. internal/render/login/configmap.go
      - Remove `MaxStartups` and `LoginGraceTime` from sshd config.
        ```
        /mnt/ssh-configs/sshd_config line 18: Directive 'MaxStartups' is not allowed within a Match block
        ```
        Likewise for 'LoginGraceTime'.
      - Add a space between `MaxAuthTries` and `consts.SSHDMaxAuthTries`.
1. slurm-cluster
   1. helm/slurm-cluster/values.yaml
      - Replace gpu/non-gpu node matching label from `nebius.com/node-group-id` to `nvidia.com/gpu.present`.
      - Disable accounting in slurm nodes.
        - *accounting* pod or svc are not created by soperator, which fails the *controller-0* pod.
      - Decrease controller size from 2 to 1.
        - Bi-controller deployment does not work as expected.
          - *controller-0* pod fails to resolve `controller-1.slurm1-controller-svc.slurm1.svc.cluster.local` (since *controller-1* pod has not been running) and `slurm1-accounting-svc` (such a svc does not exist, fixed by disabling accounting).
            ```
            slurmctld: error: xgetaddrinfo: getaddrinfo(controller-1.slurm1-controller-svc.slurm1.svc.cluster.local:6817) failed: Name or service not known
            slurmctld: error: slurm_set_addr: Unable to resolve "controller-1.slurm1-controller-svc.slurm1.svc.cluster.local"
            slurmctld: debug:  Requesting control from backup controller controller-1
            slurmctld: error: slurm_get_port: Address family '0' not supported
            slurmctld: error: Error connecting, bad data: family = 0, port = 0
            slurmctld: error: _shutdown_bu_thread:send/recv controller-1: No error
            slurmctld: error: xgetaddrinfo: getaddrinfo(slurm1-accounting-svc:6819) failed: Name or service not known
            slurmctld: error: slurm_set_addr: Unable to resolve "slurm1-accounting-svc"
            slurmctld: error: slurm_get_port: Address family '0' not supported
            slurmctld: error: Error connecting, bad data: family = 0, port = 0
            slurmctld: error: _open_persist_conn: failed to open persistent connection to host:slurm1-accounting-svc:6819: No error
            slurmctld: error: Sending PersistInit msg: No error
            slurmctld: accounting_storage/slurmdbd: clusteracct_storage_p_register_ctld: Registering slurmctld at port 6817 with slurmdbd
            slurmctld: error: xgetaddrinfo: getaddrinfo(slurm1-accounting-svc:6819) failed: Name or service not known
            slurmctld: error: slurm_set_addr: Unable to resolve "slurm1-accounting-svc"
            slurmctld: error: slurm_get_port: Address family '0' not supported
            slurmctld: error: Error connecting, bad data: family = 0, port = 0
            slurmctld: error: Sending PersistInit msg: No such file or directory
            ```
          - *controller-1- pod stuck at waiting for heartbeat.
            ```
            slurmctld: debug:  get_last_heartbeat: sleeping before attempt 1 to open heartbeat
            slurmctld: debug:  get_last_heartbeat: sleeping before attempt 2 to open heartbeat
            slurmctld: error: get_last_heartbeat: heartbeat open attempt failed from /var/spool/slurmctld/heartbeat.
            slurmctld: warning: Waiting for heartbeat file to exist...
            ```
      - Decrease the resource requests from the *slurmctld* and *munge* containers in *worker* pods (insufficient node resources).
      - Replace the storage class of PVs/PVCs for *worker-spool* from `nebius-network-ssd` to `local-path` and decrease the requested volume size (insufficient node resources).
      - Decrease the login size from 2 to 1 (insufficient node resources).
      - Add a record in `sshRootPublicKeys` to add an authorized key.
