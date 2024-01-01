#!/usr/bin/env bash

# 创建网络
docker network create my-network

# 启动 Redis 1
docker run -d --name go-server-redis --network my-network redis

# 启动 Redis 2
docker run -d --name my-redis --network my-network redis

# 等5秒
sleep 5

# 启动 MySQL
docker run -d --name tok-persistor -e MYSQL_ROOT_PASSWORD=fuckyou --network my-network -p 3307:3306 mysql

# 启动 tok-server-image
docker run -d --name tok_container --network my-network -p 8080:2222 tok-server-image

# 启动 tok_crawler_image
docker run -d --name tok_crawler_container --network my-network tok_crawler_image

# 输出部署完成信息
echo "Deployment completed successfully."
