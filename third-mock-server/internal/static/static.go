// Package static 承载前端构建产物的嵌入文件系统。
// Package static hosts the embedded filesystem of the frontend build artifacts.
//
// third-mock-web 仓库执行 `npm run build` 后，将 dist/ 拷贝至本包 dist/ 目录
// （make embed-web），即可随服务端编译为单二进制发布。
// After `npm run build` in the third-mock-web repository, copy dist/ into this
// package's dist/ directory (via `make embed-web`) to ship a single binary.
//
// dist/ 内仅提交 .keep 占位：保证前端产物缺失时 go build 依然通过，
// 真实构建产物经 make embed-web 覆盖写入且不纳入版本管理。
// Only a .keep placeholder is committed under dist/ so that `go build`
// succeeds even without real frontend artifacts; real artifacts are written
// by `make embed-web` and kept out of version control.
package static

import (
	"embed"
	"io/fs"
)

// distFS 嵌入 dist/ 目录全部内容（含占位文件）。
// distFS embeds everything under dist/ (placeholder included).
//
//go:embed all:dist
var distFS embed.FS

// Dist 返回以 dist/ 为根的前端静态资源文件系统。
// Dist returns the frontend asset filesystem rooted at dist/.
func Dist() fs.FS {
	sub, err := fs.Sub(distFS, "dist")
	if err != nil {
		// embed 路径编译期已校验，理论上不会失败；防御性返回空 FS。
		// The embed path is validated at compile time and cannot fail;
		// return an empty FS defensively.
		return emptyFS{}
	}
	return sub
}

// emptyFS 空文件系统，用于防御性兜底。
// emptyFS is an empty filesystem used as a defensive fallback.
type emptyFS struct{}

func (emptyFS) Open(string) (fs.File, error) {
	return nil, fs.ErrNotExist
}
