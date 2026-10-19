package utility

import (
	"strings"
)

// 清理字串，專門移除隱藏的 BOM 簽章與不必要的空白
//
// 參數:
//   - text: string -> 原始輸入的字串（例如從檔案、API 讀取出來的原始文字）
//
// 返回值:
//   - string -> 處理完成的乾淨字串
func cleanBOMAndSpace(text string) string {
	text = strings.TrimPrefix(text, "\ufeff")
	return strings.TrimSpace(text)
}
