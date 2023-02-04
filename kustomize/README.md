# Kustomize

## How to start with kustomize
```
kubectl apply -k ./overlays/sit
```

## How to delete application
```
kubectl delete service sit-go-service
kubectl delete deployment sit-go-service
```

