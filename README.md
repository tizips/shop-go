这是一个基于 [Go](https://go.dev/) 语言开发的项目，下面的运行代码将全部基于 [Docker](https://www.docker.com/) 运行演示

## 服务说明

- ADMIN：后台管理接口 
- WEB：官网请求交互接口
- QUEUE：队列服务

## ADMIN 接口

编译镜像

```bash
docker build -t herhe/admin:1.0.0 -f docker/admin.Dockerfile .
```

运行容器

```bash
docker run \
  -d \
  --name herhe-admin \
  --net uper \
  --ip 172.19.0.115 \
  -v $PWD/conf:/app/conf \
  -v $PWD/migration:/app/migration \
  herhe/admin:1.0.0
```

## WEB 接口

编译镜像

```bash
docker build -t herhe/web:1.0.0 -f docker/web.Dockerfile .
```

运行容器

```bash
docker run \
  -d \
  --name herhe-web \
  --net uper \
  --ip 172.19.0.113 \
  -v $PWD/conf:/app/conf \
  herhe/web:1.0.0
```

## QUEUE 服务

编译镜像

```bash
docker build -t herhe/queue:1.0.0 -f docker/queue.Dockerfile .
```

运行容器

```bash
docker run \
  -d \
  --name herhe-queue \
  --net uper \
  --ip 172.19.0.114 \
  -v $PWD/conf:/app/conf \
  herhe/queue:1.0.0
```
