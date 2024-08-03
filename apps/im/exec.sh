# 自定义创建聊天数据库模型，使用mongoDB
goctl model mongo --type chatLog --dir ./apps/im/immodels
# 会话模型
goctl model mongo --type conversations --dir ./apps/im/immodels
# 会话列表
goctl model mongo --type conversation --dir ./apps/im/immodels

# 自定义创建会话rpc模型
goctl rpc protoc apps/im/rpc/im.proto --go_out=./apps/im/rpc --go-grpc_out=./apps/im/rpc --zrpc_out=./apps/im/rpc

# 构建im api服务
goctl api go -api apps/im/api/im.api -dir apps/im/api -style gozero