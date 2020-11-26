package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

func Md5(str string) string{
	data := []byte(str)
	has := md5.Sum(data)
	return strings.ToUpper(fmt.Sprintf("%x", has)[8:24])
}


//AES advanced encryption standard 高级加密标准
//16,24,32位字符串分别对应 AES-128 AES-192 AES-256 加密方法
var PwdKey = []byte("DFASDFKEWLSKDJFK")

//PKCS7 填充模式
func PKCS7Padding(src []byte, blockSize int) []byte {
	//padNum 是padNum即被加密的文本差多少个字节补满需要的位数
	padNum := blockSize - len(src) % blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padNum),复制padNum个然后合并成新的字节切片返回}
	//int转成字节是byte 但是转换之后不是string转换后的字节数组，所以要[]byte{}加入字节数组
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

//PKCS7 反填充
func PKCS7UnPadding(pad []byte) (src []byte,err error) {
	length := len(pad)
	if  length == 0 {
		err = errors.New("加密字符串为空")
		return
	} else {
		//获取填充的字符串长度，最后一位的数字就是补充的长度
		padNum := int(pad[length-1])
		tmp := length-padNum
		if tmp < 0 {
			return
		}
		src = pad[:tmp]
		return
	}
}

//加密
func EncryptAES(data []byte,key []byte)(encrypted []byte,err error){
	//创建加密算法实例
	block,err := aes.NewCipher(key)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	//对数据填充，使长度满足要求
	data = PKCS7Padding(data,blockSize)
	//采用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block,key[:blockSize])
	//加密
	encrypted = make([]byte,len(data))
	blockMode.CryptBlocks(encrypted,data)
	return
}

//再次加密base64
func EncryptBase64(pwd []byte)(res string,err error){
	result,err := EncryptAES(pwd,PwdKey)
	if err != nil {
		BgLogger.Error(err)
		return
	}
	res = base64.StdEncoding.EncodeToString(result)
	return
}


//aes解密
func DecryptAes(encrypted []byte,key []byte)(data []byte,err error){
	//创建加密实例
	block,err := aes.NewCipher(key)
	if err != nil {
		BgLogger.Error(err)
		return
	}
	blockSize := block.BlockSize()
	//创建解密实例
	blockMode := cipher.NewCBCDecrypter(block,key[:blockSize])
	//解密
	dataPad := make([]byte,len(encrypted))
	blockMode.CryptBlocks(dataPad,encrypted)
	//反填充
	if data,err = PKCS7UnPadding(dataPad);err != nil {
		BgLogger.Error(err)
		return
	}

	return
}

//base64解密
func DecryptBase64(pwd string)(data []byte,err error){
	pwdByte,err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		BgLogger.Error(err)
		return
	}
	return DecryptAes(pwdByte,PwdKey)
}