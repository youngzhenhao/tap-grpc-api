package api

import (
	"context"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"google.golang.org/grpc"
	"log"
)

func AddFederationServer() {}

func AssetLeafKeys() {}

func AssetLeaves() {}

func AssetRoots() {}

func DeleteAssetRoot() {}

func DeleteFederationServer() {}

// Info
//
//	@Description: 返回一组关于宇宙当前状态的信息
//	@return *universerpc.InfoResponse
func Info() *universerpc.InfoResponse {
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
	client := universerpc.NewUniverseClient(conn)
	// 构建请求
	request := &universerpc.InfoRequest{}
	// 得到响应
	response, err := client.Info(context.Background(), request)
	if err != nil {
		log.Fatalf("universerpc Info Error: %v", err)
	}
	// 处理结果
	return response
}

func InsertProof() {}

// ListFederationServers
//
//	@Description: 列出了组成本地 Universe 服务器联盟的服务器集。这些服务器用于推送新的证明，并定期从远程服务器同步调用新的证明
//	@return *universerpc.ListFederationServersResponse
func ListFederationServers() *universerpc.ListFederationServersResponse {
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
	client := universerpc.NewUniverseClient(conn)
	// 构建请求
	request := &universerpc.ListFederationServersRequest{}
	// 得到响应
	response, err := client.ListFederationServers(context.Background(), request)
	if err != nil {
		log.Fatalf("universerpc ListFederationServers Error: %v", err)
	}
	// 处理结果
	return response
}

func MultiverseRoot() {}

func QueryAssetRoots() {}

func QueryAssetStats() {}

func QueryEvents() {}

func QueryFederationSyncConfig() {}

func QueryProof() {}

func SetFederationSyncConfig() {}

func SyncUniverse() {}

func UniverseStats() {}
