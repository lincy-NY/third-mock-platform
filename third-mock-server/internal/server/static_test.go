package server

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// TestResolveStatic_EmbedFallback 未配置外部目录时回退嵌入产物。
// TestResolveStatic_EmbedFallback falls back to embedded assets when webDist is empty.
func TestResolveStatic_EmbedFallback(t *testing.T) {
	assetFS, external, err := ResolveStatic("")
	if err != nil {
		t.Fatalf("ResolveStatic() error = %v", err)
	}
	if external {
		t.Fatal("expected embedded mode, got external")
	}
	// 嵌入 FS 至少包含 .keep 占位，保证非空可服务。
	// The embedded FS contains at least the .keep placeholder.
	entries, readErr := fs.ReadDir(assetFS, ".")
	if readErr != nil {
		t.Fatalf("read embedded root: %v", readErr)
	}
	if len(entries) == 0 {
		t.Fatal("embedded FS is empty")
	}
}

// TestResolveStatic_ExternalDir 配置含 index.html 的外部目录时优先外部模式。
// TestResolveStatic_ExternalDir prefers the external directory when it contains index.html.
func TestResolveStatic_ExternalDir(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte("external"), 0o644); err != nil {
		t.Fatalf("write index.html: %v", err)
	}
	assetFS, external, err := ResolveStatic(dir)
	if err != nil {
		t.Fatalf("ResolveStatic() error = %v", err)
	}
	if !external {
		t.Fatal("expected external mode, got embedded")
	}
	data, readErr := fs.ReadFile(assetFS, "index.html")
	if readErr != nil {
		t.Fatalf("read external index.html: %v", readErr)
	}
	if string(data) != "external" {
		t.Fatalf("external index.html content = %q, want %q", data, "external")
	}
}

// TestResolveStatic_ExternalDirWithoutIndex 外部目录缺 index.html 时回退嵌入产物。
// TestResolveStatic_ExternalDirWithoutIndex falls back to embedded assets when the external dir lacks index.html.
func TestResolveStatic_ExternalDirWithoutIndex(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "assets"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	_, external, err := ResolveStatic(dir)
	if err != nil {
		t.Fatalf("ResolveStatic() error = %v", err)
	}
	if external {
		t.Fatal("expected embedded fallback, got external")
	}
}
