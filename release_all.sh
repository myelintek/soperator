#!/bin/bash

set -e

usage() { echo "usage: ${0} [-s] -u <ssh_user> -k <path_to_ssh_key> -a <address_of_build_agent> [-h]" >&2; exit 1; }

while getopts u:k:a:sh flag
do
    case "${flag}" in
        u) user=${OPTARG};;
        k) key=${OPTARG};;
        a) address=${OPTARG};;
        s) stable="1";;
        h) usage;;
        *) usage;;
    esac
done

if [ -z "$user" ] || [ -z "$key" ] || [ -z "$address" ]; then
    usage
fi



start_time=$(date +%s)

if [ "$stable" == "1" ]; then
    UNSTABLE="false"
    echo "Syncing versions among all files for stable release"
else
    UNSTABLE="true"
    echo "Syncing versions among all files for unstable release"
fi


make sync-version UNSTABLE=${UNSTABLE}
IMAGE_VERSION=$(make get-image-version UNSTABLE=${UNSTABLE})
VERSION=$(make get-version UNSTABLE=${UNSTABLE})

OPERATOR_IMAGE_TAG=$(make get-operator-tag-version UNSTABLE=${UNSTABLE})

echo "Version is ${VERSION}"
echo "Image version is ${IMAGE_VERSION}"
echo "Operator image tag version is ${OPERATOR_IMAGE_TAG}"

echo "Updating CRDs & auto-generated code (included in test step) & run tests"
make test UNSTABLE="${UNSTABLE}"

echo "Uploading images to the build agent"
./images/upload_to_build_agent.sh -u "$user" -k "$key"

remote_command=$(cat <<EOF
set -e
set -x

echo "Entering /usr/src/prototypes/slurm/${user}"
cd "/usr/src/prototypes/slurm/${user}"
sudo su -- <<'ENDSSH'

echo "Building and pushing container images"

make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=worker_slurmd DOCKERFILE=worker/slurmd.dockerfile
make docker-push  UNSTABLE="${UNSTABLE}" IMAGE_NAME=worker_slurmd

make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=controller_slurmctld DOCKERFILE=controller/slurmctld.dockerfile
make docker-push  UNSTABLE="${UNSTABLE}" IMAGE_NAME=controller_slurmctld

make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=login_sshd DOCKERFILE=login/sshd.dockerfile
make docker-push  UNSTABLE="${UNSTABLE}" IMAGE_NAME=login_sshd

make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=munge DOCKERFILE=munge/munge.dockerfile
make docker-push  UNSTABLE="${UNSTABLE}" IMAGE_NAME=munge

make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=nccl_benchmark DOCKERFILE=nccl_benchmark/nccl_benchmark.dockerfile
make docker-push  UNSTABLE="${UNSTABLE}" IMAGE_NAME=nccl_benchmark
echo "Common images were built"

echo "Removing previous jail rootfs tar archive"
rm -rf images/jail_rootfs.tar

echo "Building tarball for jail"
make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=jail DOCKERFILE=jail/jail.dockerfile DOCKER_OUTPUT="--output type=tar,dest=jail_rootfs.tar"
echo "Built tarball jail_rootfs.tar"

make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=populate_jail DOCKERFILE=populate_jail/populate_jail.dockerfile
make docker-push  UNSTABLE="${UNSTABLE}" IMAGE_NAME=populate_jail

echo "Building image of the operator"
make docker-build UNSTABLE="${UNSTABLE}" IMAGE_NAME=slurm-operator DOCKERFILE=Dockerfile IMAGE_VERSION="$OPERATOR_IMAGE_TAG"
echo "Pushing image of the operator"
make docker-push UNSTABLE="${UNSTABLE}" IMAGE_NAME=slurm-operator IMAGE_VERSION="$OPERATOR_IMAGE_TAG"

echo "Pushing Helm charts"
make release-helm UNSTABLE="${UNSTABLE}" OPERATOR_IMAGE_TAG="${OPERATOR_IMAGE_TAG}"

wait

ENDSSH
EOF
)

ssh -i "$key" "$user"@"$address" "${remote_command}"
echo "All images are built successfully"

######################################################
### The code below should be moved to terraform repo
######################################################

echo "Packing new terraform tarball"
VERSION=${OPERATOR_IMAGE_TAG} ./terraform/release_terraform.sh -f

echo "Unpacking the terraform tarball"
version_formatted=$(echo "${OPERATOR_IMAGE_TAG}" | tr '-' '_' | tr '.' '_')
tarball="slurm_operator_tf_$version_formatted.tar.gz"

pushd ./terraform/releases/unstable
    VERSION=${OPERATOR_IMAGE_TAG} TARBALL=${tarball} ./unpack_current_version.sh
popd

GREEN='\033[0;32m'
RESET='\033[0m'

if [ "$stable" == "1" ]; then
    mv "terraform/releases/unstable/$tarball" "terraform/releases/stable/"
    echo -e "${GREEN}Stable version '$OPERATOR_IMAGE_TAG' is successfully released${RESET}"
else
    echo -e "${GREEN}Unstable version '$OPERATOR_IMAGE_TAG' is successfully released and unpacked to releases/unstable/${RESET}"
fi

end_time=$(date +%s)
duration=$((end_time - start_time))

echo "Time taken: ${duration} seconds"
