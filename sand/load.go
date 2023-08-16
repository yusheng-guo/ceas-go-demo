package sand

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"os"
)

func LoadPrivateKey(pemPath, password string) (*rsa.PrivateKey, error) {
	// 1.读取 PFX 文件
	pfxData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, fmt.Errorf("read file %s, err: %w", pemPath, err)
	}

	// 杉德pfx私钥格式(私钥 证书1 证书2 证书3 )
	// golang pkcs12 解码格式( 私钥 证书 )
	blocks, err := pkcs12.ToPEM(pfxData, "nft123456")
	if err != nil {
		return nil, err
	}
	var priv any
	for _, block := range blocks {
		if block.Type == "PRIVATE KEY" {
			priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("parse pkcs private key, err: %w", err)
			}
		}
	}
	if err := priv.(*rsa.PrivateKey).Validate(); err != nil {
		return nil, fmt.Errorf("validate private key, err: %w", err)
	}
	return priv.(*rsa.PrivateKey), nil
}

// LoadPrivateKey 加载私钥 私钥密码: nft123456
//func LoadPrivateKey(pemPath, password string) (*rsa.PrivateKey, error) {
//	// 1.读取 PFX 文件
//	pfxData, err := os.ReadFile(pemPath)
//	if err != nil {
//		return nil, fmt.Errorf("read file %s, err: %w", pemPath, err)
//	}
//	// 2.解码 pfx 格式数据
//	privateKey, _, err := pkcs12.Decode(pfxData, password)
//	if err != nil {
//		return nil, fmt.Errorf("pkcs12 decode, err: %w", err)
//	}
//	// 4.获取 RSA 私钥
//	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
//	if !ok {
//		return nil, errors.New("failed to parse RSA private key")
//	}
//	return rsaPrivateKey, nil
//}

//func LoadPrivateKey(privateKeyName, privatePassword string) (*rsa.PrivateKey, error) {
//	bytes, err := os.ReadFile(privateKeyName)
//	if err != nil {
//		return nil, err
//	}
//	// 因为pfx证书公钥和密钥是成对的，所以要先转成pem.Block
//	blocks, err := pkcs12.ToPEM(bytes, privatePassword)
//	if err != nil {
//		return nil, err
//	}
//	if len(blocks) != 2 {
//		return nil, errors.New("解密错误")
//	}
//	// 拿到第一个block，用x509解析出私钥（当然公钥也是可以的）
//	privateKey, err := x509.ParsePKCS1PrivateKey(blocks[0].Bytes)
//	if err != nil {
//		return nil, err
//	}
//	return privateKey, nil
//}

// LoadPublicKey 加载公钥
func LoadPublicKey(pemPath string) (*rsa.PublicKey, error) {
	// 1.读取公钥文件
	fmt.Println("读取公钥文件")
	key, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, fmt.Errorf("read file, err: %w", err)
	}
	// 2.解码 PEM 格式数据
	//block, _ := pem.Decode(key)
	//if block == nil {
	//	return nil, errors.New("failed to parse PEM block")
	//}
	// 3.解析证书
	//certBody, err := x509.ParseCertificate(block.Bytes)
	certBody, err := x509.ParseCertificate(key)
	if err != nil {
		return nil, fmt.Errorf("parse cert, err: %w", err)
	}
	// 4.获取公钥
	publicKey, ok := certBody.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to get public key")
	}
	return publicKey, nil
}
