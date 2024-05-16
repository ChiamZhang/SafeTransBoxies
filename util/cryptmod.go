package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/alexmullins/zip"
	"io"
	"os"
)

// 加密函数
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充明文数据
	plaintext = pkcs7Pad(plaintext, aes.BlockSize)

	// 创建一个与填充后的明文长度相同的字节数组用于存储加密后的密文
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// 生成一个随机的初始化向量
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 创建一个 CBC 模式的加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密数据
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// 解密函数
func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, io.ErrUnexpectedEOF
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, io.ErrUnexpectedEOF
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext := pkcs7Unpad(ciphertext)
	if plaintext == nil {
		return nil, io.ErrUnexpectedEOF
	}

	return plaintext, nil
}

// pkcs7Pad 使用 PKCS7 填充方案填充数据
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad 使用 PKCS7 解填充方案
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil
	}
	return data[:(length - unpadding)]
}

// 加密文件名
func encryptFileName(fileName string, key []byte) (string, error) {
	ciphertext, err := encrypt([]byte(fileName), key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// 解密文件名
func decryptFileName(encryptedFileName string, key []byte) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedFileName)
	if err != nil {
		return "", err
	}
	plaintext, err := decrypt(ciphertext, key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// 创建加密并带密码保护的zip文件
func CreateEncryptedZipFile(middlewares map[string][]byte, key []byte, zipPassword string, outputPath string) error {
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for fileName, content := range middlewares {
		// 加密文件名
		encryptedFileName, err := encryptFileName(fileName, key)
		if err != nil {
			return err
		}

		// 加密文件内容
		ciphertext, err := encrypt(content, key)
		if err != nil {
			return err
		}

		// 创建加密文件并设置密码
		writer, err := zipWriter.Encrypt(encryptedFileName, zipPassword)
		if err != nil {
			return err
		}

		_, err = writer.Write(ciphertext)
		if err != nil {
			return err
		}
	}

	return nil
}

// 读取并解密 ZIP 文件内容
func readEncryptedZipFile(zipPath string, key []byte, zipPassword string) (map[string][]byte, error) {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()

	decryptedFiles := make(map[string][]byte)

	for _, file := range zipReader.File {
		// 打开文件并解密内容
		file.SetPassword(zipPassword)
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		var buffer bytes.Buffer
		if _, err := io.Copy(&buffer, rc); err != nil {
			rc.Close()
			return nil, err
		}
		rc.Close()

		ciphertext := buffer.Bytes()
		plaintext, err := decrypt(ciphertext, key)
		if err != nil {
			return nil, err
		}

		// 解密文件名
		decryptedFileName, err := decryptFileName(file.Name, key)
		if err != nil {
			return nil, err
		}

		decryptedFiles[decryptedFileName] = plaintext
	}

	return decryptedFiles, nil
}

/** 以下为调用方式
func main() {
	// 中间件名称及内容
	middlewares := map[string][]byte{
		"middleware1.txt": []byte("这是中间件1的内容"),
		"middleware2.txt": []byte("这是中间件2的内容"),
	}

	// AES密钥，注意密钥长度必须是16、24或32字节
	key := []byte("0123456789abcdef0123456789abcdef")

	// ZIP 文件密码
	zipPassword := "123456"

	// 输出加密的zip文件路径
	outputPath := "encrypted_middlewares.zip"

	// 创建加密并带密码保护的zip文件
	err := createEncryptedZipFile(middlewares, key, zipPassword, outputPath)
	if err != nil {
		panic("错误: " + err.Error())
	}

	fmt.Println("加密ZIP文件创建成功")

	// 读取并解密 ZIP 文件内容
	decryptedFiles, err := readEncryptedZipFile("encrypted_middlewares.safebox", key, zipPassword)
	if err != nil {
		panic("错误: " + err.Error())
	}

	// 打印解密后的文件内容
	for fileName, content := range decryptedFiles {
		fmt.Println("文件名:", fileName)
		fmt.Println("内容:", string(content))
	}
}
**/
