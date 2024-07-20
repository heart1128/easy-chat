user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test  # 编译user-rpc-dev

user-api-dev:
	@make -f deploy/mk/user-api.mk release-test  # 编译user-api-dev

social-rpc-dev:
	@make -f deploy/mk/social-rpc.mk release-test  # 编译social-rpc-dev

social-api-dev:
	@make -f deploy/mk/social-api.mk release-test  # 编译social-api-dev


# 左边目标，右边依赖
release-test:user-rpc-dev user-api-dev social-rpc-dev social-api-dev

# 执行脚本，给权限，启动
install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh