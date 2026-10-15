package server

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"third-mock-server/internal/static"
)

// ResolveStatic 依据配置解析前端静态资源来源。
// ResolveStatic resolves the frontend static asset source by configuration.
//
// 优先级 / Priority:
//  1. webDist 非空且目录下存在 index.html → 使用外部目录（开发调试可整目录替换）
//  2. 否则回退 go:embed 嵌入产物（单二进制部署）
//
// 返回值同时给出来源标记，便于启动日志提示当前模式。
// The returned source flag indicates which mode is active for startup logging.
func ResolveStatic(webDist string) (assetFS fs.FS, external bool, err error) {
	if trimmed := strings.TrimSpace(webDist); trimmed != "" {
		indexPath := filepath.Join(trimmed, "index.html")
		if info, statErr := os.Stat(indexPath); statErr == nil && !info.IsDir() {
			return os.DirFS(trimmed), true, nil
		}
	}
	return static.Dist(), false, nil
}
