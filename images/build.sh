#!/bin/bash

set -e # Exit immediately if any command returns a non-zero error code

usage() { echo "usage: ${0} -d <dockerfile> -t <docker-target> -v <version> [-i] [-n] [-s] [-h]" >&2; exit 1; }

while getopts d:t:v:insh flag
do
    case "${flag}" in
        d) dockerfile=${OPTARG};;
        t) target=${OPTARG};;
        v) version=${OPTARG};;
        i) ignore_cache=--no-cache;;
        n) no_push=1;;
        s) stable=1;;
        h) usage;;
        *) usage;;
    esac
done

if [ -z "${dockerfile}" ] || [ -z "${target}" ] || [ -z "${version}" ]; then
    usage
fi

docker build --tag "${target}:${version}" --target "${target}" ${ignore_cache} --load --platform=linux/amd64 -f "${dockerfile}" .
echo "Built image ${target}:${version}"

if [ -z "${no_push}" ]; then
    # Do not touch! These 3 lines are generated by make sync-version
    # Stable: https://console.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnlu9nio77sg3p8n5bi/overview.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnlu9nio77sg3p8n5bi/overview.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnlu9nio77sg3p8n5bi/overview.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnlu9nio77sg3p8n5bi/overview
    # Unstable: https://console.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnbu823dealq64cp1s6/overview.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnbu823dealq64cp1s6/overview.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnbu823dealq64cp1s6/overview.nebius.ai/folders/bje82q7sm8njm3c4rrlq/container-registry/registries/crnbu823dealq64cp1s6/overview
    if [ -n "${stable}" ]; then container_registry_id=crnlu9nio77sg3p8n5bi; else container_registry_id=crnbu823dealq64cp1s6; fi

    docker tag "${target}:${version}" "cr.ai.nebius.cloud/${container_registry_id}/${target}:${version}"
    docker push "cr.ai.nebius.cloud/${container_registry_id}/${target}:${version}"
    echo "Pushed image ${target}:${version}"
fi
