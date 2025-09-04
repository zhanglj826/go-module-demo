package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// 计算文本的MD5值
func TextMD5(text, salt string) string {
	// 组合盐值和文本（盐值在前，文本在后）
	data := append([]byte(salt), []byte(text)...)

	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}

// 计算文件的MD5值
func FileMD5(filePath, salt string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()

	// 先写入盐值
	hasher.Write([]byte(salt))

	// 分块读取文件内容并更新哈希
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hasher.Write(buffer[:n])
	}

	// 转换为十六进制字符串
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
