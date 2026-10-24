package home_api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIndexActionsRenderLocalizedHomePages(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name    string
		path    string
		handler gin.HandlerFunc
		want    []string
	}{
		{
			name:    "English",
			path:    "/",
			handler: IndexAction,
			want: []string{
				`<html lang="en">`,
				"<title>iNeovim — Native Neovim for macOS</title>",
				"Neovim, at home on macOS.",
				"libghostty",
				"Diagnostics at a click",
				"Cmd+P",
				"https://github.com/daqing/iNeovim",
				"https://github.com/daqing/iNeovim-website/releases/tag/v0.5.0",
				"/static/iNeovim-screenshot.png",
			},
		},
		{
			name:    "Chinese",
			path:    "/zh/",
			handler: ChineseIndexAction,
			want: []string{
				`<html lang="zh-CN">`,
				"<title>iNeovim — macOS 原生 Neovim 客户端</title>",
				"Neovim，融入 macOS。",
				"独立终端面板",
				"点击查看诊断详情",
				"Problems 侧栏",
				"Cmd+P",
				"https://github.com/daqing/iNeovim",
				"https://github.com/daqing/iNeovim-website/releases/tag/v0.5.0",
				"/static/iNeovim-screenshot.png",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.GET(tt.path, tt.handler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			r.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("expected status 200, got %d", w.Code)
			}

			if ct := w.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
				t.Fatalf("expected a text/html content type, got %q", ct)
			}

			body := w.Body.String()
			for _, want := range tt.want {
				if !strings.Contains(body, want) {
					t.Errorf("expected body to contain %q", want)
				}
			}
		})
	}
}
