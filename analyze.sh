#!/usr/bin/env sh

IMAGE=$1

SIZE=$(docker image inspect --format {{.Size}} $IMAGE)

echo "Uncompressed size: $(($SIZE / 1000000))MB"

echo "Using 'docker sbom' to generate Software Bill Of Materials"
docker sbom -o "results/sbom-$IMAGE.txt" "$IMAGE"

echo "Using 'docker scan' to search for vulnerabilities"
docker scan "$IMAGE" > results/scan-$IMAGE.txt