#!/bin/bash
kubectl -n kube-system get configmap coredns -o yaml | sed 's/\/etc\/resolv.conf/1.1.1.1/gi' | kubectl apply -f -
PODNAMES=(`kubectl -n kube-system get pods -o jsonpath='{.items[*].metadata.name}'`)
for name in ${PODNAMES[@]}; do
    if echo "$name" | grep -q 'coredns-'; then
        kubectl -n kube-system delete pods "$name"
    fi
done
