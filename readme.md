## 镜像同步

```bash
GOOS=linux GOARCH=amd64 go build -o push_amd64 cmd/push/main.go
# 由github workflow触发
bash sync.sh
```

github action 设置secret

## 镜像下载

```bash
GOOS=linux GOARCH=amd64 go build -o pull_amd64 cmd/pull/main.go
go build -o pull_mac cmd/pull/main.go
./pull_mac -image k8s.gcr.io/pause:3.4.1
```