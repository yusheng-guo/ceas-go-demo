package utils

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

	// 2.转为 pem 格式
	// 杉德pfx私钥格式(私钥 证书1 证书2 证书3 )
	// golang pkcs12 解码格式( 私钥 证书 )
	blocks, err := pkcs12.ToPEM(pfxData, "nft123456")
	if err != nil {
		return nil, err
	}

	// 3.获取私钥
	var priv any
	for _, block := range blocks {
		if block.Type == "PRIVATE KEY" {
			priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("parse pkcs private key, err: %w", err)
			}
		}
	}
	// 4.验证私钥合法性
	if err = priv.(*rsa.PrivateKey).Validate(); err != nil {
		return nil, fmt.Errorf("validate private key, err: %w", err)
	}
	return priv.(*rsa.PrivateKey), nil
}

// LoadPublicKey 加载公钥
//
//	func LoadPublicKey(pemPath string) (*rsa.PublicKey, error) {
//		// 1.读取公钥文件
//		certData, err := os.ReadFile(pemPath)
//		if err != nil {
//			return nil, fmt.Errorf("read file, err: %w", err)
//		}
//		//block, _ := pem.Decode(certData)
//		//if block == nil {
//		//	return nil, fmt.Errorf("decode PEM block, err: %w", err)
//		//}
//		// 2.解析证书
//		cert, err := x509.ParseCertificate(certData)
//		if err != nil {
//			return nil, fmt.Errorf("parse cert, err: %w", err)
//		}
//
//		// 验证证书的签名
//		//err = cert.CheckSignature(cert.SignatureAlgorithm, cert.RawTBSCertificate, cert.Signature)
//		//if err != nil {
//		//	return nil, fmt.Errorf("check signature, err: %w", err)
//		//}
//
//		// 验证证书的有效期限
//		now := time.Now()
//		if now.Before(cert.NotBefore) || now.After(cert.NotAfter) {
//			// log.Fatal("Certificate is expired or not yet valid")
//			return nil, errors.New("certificate is expired or not yet valid")
//		}
//
//		// 验证证书的主题和颁发者
//		fmt.Println(cert.Subject.CommonName)
//		fmt.Println(cert.Issuer.CommonName)
//		if cert.Subject.CommonName != cert.Issuer.CommonName {
//			return nil, errors.New("certificate subject and issuer do not match")
//		}
//
//		// 3.获取公钥
//		publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
//		if !ok {
//			return nil, errors.New("failed to get public key")
//		}
//		return publicKey, nil
//	}
func LoadPublicKey(pemPath string) (*rsa.PublicKey, error) {
	// 1.读取公钥文件
	certData, err := os.ReadFile(pemPath)
	if err != nil {
		return nil, fmt.Errorf("read file, err: %w", err)
	}
	//block, _ := pem.Decode(certData)
	//if block == nil {
	//	return nil, fmt.Errorf("failed decode PEM block")
	//}
	// 2.解析证书
	cert, err := x509.ParseCertificate(certData)
	if err != nil {
		return nil, fmt.Errorf("parse cert, err: %w", err)
	}

	// 3.获取公钥
	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("failed to get public key")
	}
	return publicKey, nil
}
