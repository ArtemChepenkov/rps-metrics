kubectl apply -f deployments/client-deployment.yaml
kubectl apply -f deployments/server-deployment.yaml
kubectl apply -f deployments/prometheus-config.yaml
kubectl apply -f deployments/prometheus-deployment.yaml
kubectl apply -f deployments/blackbox-deployment.yaml
kubectl apply -f deployments/grafana-deployment.yaml
kubectl apply -f deployments/node-exporter.yaml
