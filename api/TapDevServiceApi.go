package api

import (
	"context"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc/tapdevrpc"
	"google.golang.org/grpc"
	"log"
)

// ImportProof
//
//	@Description: 尝试将证明文件导入守护进程。如果成功，将在磁盘上插入一个新资产，可使用指定的目标脚本密钥和内部密钥进行花费
//	@param proofFile
//	@param genesisPoint
//	@return *tapdevrpc.ImportProofResponse
func ImportProof(proofFile, genesisPoint string) *tapdevrpc.ImportProofResponse {
	// 读取参数
	grpcHost := getEnv("RPC_SERVER")
	tlsCertPath := getEnv("TLS_CERT_PATH")
	macaroonPath := getEnv("MACAROON_PATH")
	// 处理证书和macaroon
	creds := newTlsCert(tlsCertPath)
	macaroon := getMacaroon(macaroonPath)
	// 连接到grpc服务器
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: grpc.Dial: %v", err)
	}
	// 匿名函数延迟关闭grpc连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("conn Close Error: %v", err)
		}
	}(conn)
	// 创建客户端
	client := tapdevrpc.NewTapDevClient(conn)
	_proofFileByteSlice, _ := hex.DecodeString(proofFile)
	// 构建请求
	request := &tapdevrpc.ImportProofRequest{
		ProofFile:    _proofFileByteSlice,
		GenesisPoint: genesisPoint,
	}
	// 得到响应
	response, err := client.ImportProof(context.Background(), request)
	if err != nil {
		log.Fatalf("tapdevrpc QueryRfqAcceptedQuotes Error: %v", err)
	}
	// 处理结果
	return response
}
