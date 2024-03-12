package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

type macaroonCredential struct {
	macaroon string
}

func newMacaroonCredential(macaroon string) *macaroonCredential {
	return &macaroonCredential{macaroon: macaroon}
}

func (c *macaroonCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"macaroon": c.macaroon}, nil
}

func (c *macaroonCredential) RequireTransportSecurity() bool {
	return true
}

// getEnv
//
//	@Description: 读取.env中指定参数的值
//	@param key
//	@param filename
//	@return string
func getEnv(key string, filename ...string) string {
	// 读取.env配置
	err := godotenv.Load(filename...)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	value := os.Getenv(key)
	return value
}

// newTlsCert
//
//	@Description: 给定tls.cert证书路径，创建grpc的传输凭证并返回
//	@param tlsCertPath
//	@return credentials.TransportCredentials
func newTlsCert(tlsCertPath string) credentials.TransportCredentials {
	// 读取tls.cert证书
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}
	// 创建证书池
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}
	// 创建tls配置
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	// 创建新的TLS凭证
	creds := credentials.NewTLS(config)
	return creds
}

// getMacaroon
//
//	@Description: 给定macaroon文件存储路径，将其读取为十六进制并返回
//	@param macaroonPath
//	@return string
func getMacaroon(macaroonPath string) string {
	// macaroon读取为Bytes
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	// macaroon转为十六进制
	macaroon := hex.EncodeToString(macaroonBytes)
	return macaroon
}
