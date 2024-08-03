user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test  # 编译user-rpc-dev

user-api-dev:
	@make -f deploy/mk/user-api.mk release-test  # 编译user-api-dev

social-rpc-dev:
	@make -f deploy/mk/social-rpc.mk release-test  # 编译social-rpc-dev

social-api-dev:
	@make -f deploy/mk/social-api.mk release-test  # 编译social-api-dev

im-rpc-dev:
	@make -f deploy/mk/im-rpc.mk release-test  # 编译user-rpc-dev

im-api-dev:
	@make -f deploy/mk/im-api.mk release-test  # 编译user-api-dev

im-ws-dev:
	@make -f deploy/mk/im-ws.mk release-test  # 编译user-api-dev

task-mq-dev:
	@make -f deploy/mk/task-mq.mk release-test  # 编译user-api-dev

# 左边目标，右边依赖
release-test:user-rpc-dev user-api-dev social-rpc-dev social-api-dev im-rpc-dev im-api-dev im-ws-dev task-mq-dev

# 执行脚本，给权限，启动
install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh

install-server-user-rpc:
	cd ./deploy/script && chmod +x user-rpc-test.sh && ./user-rpc-test.sh