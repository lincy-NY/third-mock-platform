# 工作区聚合构建脚本 / Workspace aggregate build script
# third-mock-platform 本身不是 git 仓库，仅承载两个子仓库与聚合发布入口。
# third-mock-platform itself is NOT a git repo; it hosts the two sub-repos
# and the aggregate release entry.
#
# 用法 / Usage:
#   make release      # 全链路发布：web 构建 → 嵌入 → 服务端全平台编译
#   make build-web    # 仅构建前端
#   make embed-web    # 仅拷贝前端产物到服务端 embed 目录
#   make build-server # 仅编译服务端当前平台
#   make clean        # 清理全部构建产物

.PHONY: release build-web embed-web build-server clean

SERVER_DIR := third-mock-server
WEB_DIR    := third-mock-web

## release: 全链路发布（前端产物嵌入服务端后交叉编译全平台）
release: build-web embed-web build-server-all
	@echo "==> release done, artifacts in $(SERVER_DIR)/bin/"

## build-web: 构建 Vue 前端
build-web:
	@echo "==> npm run build ($(WEB_DIR))"
	cd $(WEB_DIR) && npm run build

## embed-web: 将前端产物拷贝到服务端 go:embed 目录
embed-web:
	$(MAKE) -C $(SERVER_DIR) embed-web

## build-server: 编译服务端当前平台二进制
build-server:
	$(MAKE) -C $(SERVER_DIR) build

## build-server-all: 交叉编译服务端全平台二进制
build-server-all:
	$(MAKE) -C $(SERVER_DIR) build-all

## test: 运行服务端测试
test:
	$(MAKE) -C $(SERVER_DIR) test

## clean: 清理前端与服务端构建产物
clean:
	rm -rf $(WEB_DIR)/dist
	$(MAKE) -C $(SERVER_DIR) clean
