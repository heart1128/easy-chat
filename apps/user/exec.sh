goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc/ --go-grpc_out=./apps/user/rpc --zrpc_out=./apps/user/rpc

# goctl user数据库模型构建 -c加缓存
goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c
# goctl 构建api服务
goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero