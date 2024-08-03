#!/bin/bash
reso_addr='registry.cn-hangzhou.aliyuncs.com/my_easy-chat/task-mq-dev'
tag='latest'

container_name="easy-chat-task-mq-test"

# 删除所有rpc容器和镜像，拉取新的镜像
docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-im -v /easy-im/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run --name=${container_name} -d ${reso_addr}:${tag}