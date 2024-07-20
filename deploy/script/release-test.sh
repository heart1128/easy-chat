# 所有启动的服务，执行脚本
# 定一个数组包含所有执行脚本
need_start_server_shell=(
 # rpc
  user-rpc-test.sh
  social-rpc-test.sh

  # api
  user-api-test.sh
  social-api-test.sh
)

# for循环加权限，执行
for i in ${need_start_server_shell[*]} ; do
    chmod +x $i
    ./$i
done

# 验证容器运行
docker ps
# 查看容器中的etcd中所有注册的key，验证运行
docker exec -it etcd etcdctl get --prefix ""