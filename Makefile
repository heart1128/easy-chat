user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test  # 编译user-rpc-dev

# 左边目标，右边依赖
release-test:user-rpc-dev

# 执行脚本，给权限，启动
install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh