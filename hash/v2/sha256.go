package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// CalculateTextSHA256 计算文本的 SHA-256 哈希值
func TextSHA256(text, salt string) string {
	// 组合盐值和文本
	data := append([]byte(salt), []byte(text)...)

	hasher := sha256.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

// CalculateFileSHA256 计算文件的 SHA-256 哈希值
func FileSHA256(filePath, salt string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()

	// 先写入盐值
	hasher.Write([]byte(salt))

	buffer := make([]byte, 4096) // 分块读取，适合大文件

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // 读取完毕
			}
			return "", err // 其他读取错误
		}
		hasher.Write(buffer[:n]) // 写入实际读取的字节数
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
