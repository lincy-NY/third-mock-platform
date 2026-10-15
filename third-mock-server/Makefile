# 模拟报文平台服务端构建脚本 / Build script for third-mock-server
# 用法 / Usage:
#   make            # 构建当前平台 / build for current platform
#   make build-all  # 交叉编译全平台 / cross build all platforms
#   make embed-web  # 拷贝前端构建产物到 embed 目录 / copy web dist into embed dir
#   make clean      # 清理产物 / clean artifacts
#   make test       # 运行测试 / run tests
#   make fmt        # 格式化 / gofmt

APP        := third-mock-server
VERSION    ?= dev
BINARY     := bin
LDFLAGS    := -s -w -X main.version=$(VERSION)

# 前端产物目录（third-mock-web 的构建输出）
# Frontend dist directory (build output of third-mock-web)
WEB_DIST   ?= ../third-mock-web/dist

# 嵌入目录（go:embed，占位文件会被真实产物覆盖）
# Embed directory (go:embed; placeholder gets overwritten by real artifacts)
EMBED_DIST := internal/static/dist

# 目标平台：windows / linux / darwin × amd64 / arm64
TARGETS := windows/amd64 windows/arm64 linux/amd64 linux/arm64 darwin/amd64 darwin/arm64

.PHONY: default build build-all embed-web clean test fmt vet

default: build

## build: 构建当前平台二进制（server + mockctl）
build:
	@mkdir -p $(BINARY)
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY)/ ./cmd/...

## build-all: 交叉编译 windows/linux/macos 全平台
build-all:
	@mkdir -p $(BINARY)
	@for t in $(TARGETS); do \
		os=$${t%%/*}; arch=$${t##*/}; \
		ext=""; [ "$$os" = "windows" ] && ext=".exe"; \
		name=$(APP)-$$os-$$arch$$ext; \
		echo "==> build $$name"; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 \
		go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY)/$$name ./cmd/server; \
		name=$(APP)-ctl-$$os-$$arch$$ext; \
		echo "==> build $$name"; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 \
		go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY)/$$name ./cmd/mockctl; \
	done

## embed-web: 拷贝前端构建产物到 embed 目录（保留 .keep 占位，先清空再拷贝）
embed-web:
	@if [ ! -f "$(WEB_DIST)/index.html" ]; then \
		echo "error: $(WEB_DIST)/index.html 不存在，请先在 third-mock-web 执行 npm run build"; \
		exit 1; \
	fi
	@find $(EMBED_DIST) -mindepth 1 ! -name '.keep' -delete
	@cp -R "$(WEB_DIST)/." "$(EMBED_DIST)/"
	@echo "==> embedded web dist from $(WEB_DIST) -> $(EMBED_DIST)"

## test: 运行全部测试
test:
	go test ./...

## vet: 静态检查
vet:
	go vet ./...

## fmt: 格式化代码
fmt:
	gofmt -w .

## clean: 清理构建产物
clean:
	rm -rf $(BINARY)
