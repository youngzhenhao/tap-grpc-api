package api

import (
	"context"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"google.golang.org/grpc"
	"log"
)

func AddrReceives() {

}

func BurnAsset() {

}

func DebugLevel() {

}

func DecodeAddr() {

}

func DecodeProof() {

}

func ExportProof() {

}

func FetchAssetMeta() {

}

// GetInfo
//
//	@Description: 返回节点的信息
//	@return *taprpc.GetInfoResponse
func GetInfo() *taprpc.GetInfoResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.GetInfoRequest{}
	// 得到响应
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc GetInfo Error: %v", err)
	}
	// 处理结果
	return response
}

// ListAssets
//
//	@Description: 列出目标守护程序拥有的资产集，返回的嵌套结构体中的[]byte需要使用get函数单独处理，如hex.EncodeToString
//	请注意，IncludeSpent和IncludeLeased不能同时指定为true
//	@return *taprpc.ListAssetResponse
func ListAssets(withWitness, includeSpent, includeLeased bool) *taprpc.ListAssetResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListAssetRequest{
		WithWitness:   withWitness,
		IncludeSpent:  includeSpent,
		IncludeLeased: includeLeased,
	}
	// 得到响应
	response, err := client.ListAssets(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc ListAssets Error: %v", err)
	}
	// 处理结果
	return response
}

// ListBalances
//
//	@Description: 列出资产余额
//	参数为true按照AssetsId排序，false则按照GroupKey排序，资产和组键过滤器未设置
//	@return *taprpc.ListBalancesResponse
func ListBalances(isListAssetIdNotGroupKey bool) *taprpc.ListBalancesResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListBalancesRequest{
		// 不设置过滤器
		AssetFilter:    nil,
		GroupKeyFilter: nil,
	}
	// 根据给定参数修改请求结构体的排序方式
	if isListAssetIdNotGroupKey {
		request.GroupBy = &taprpc.ListBalancesRequest_AssetId{AssetId: true}
	} else {
		request.GroupBy = &taprpc.ListBalancesRequest_GroupKey{GroupKey: true}
	}
	// 得到响应
	response, err := client.ListBalances(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc ListBalances Error: %v", err)
	}
	// 处理结果
	return response
}

// ListGroups
//
//	@Description: 列出目标守护程序已知的资产组，以及每个组中持有的资产
//	@return *taprpc.ListGroupsResponse
func ListGroups() *taprpc.ListGroupsResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListGroupsRequest{}
	// 得到响应
	response, err := client.ListGroups(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc ListGroups Error: %v", err)
	}
	// 处理结果
	return response
}

// ListTransfers
//
//	@Description: 列出目标守护程序跟踪的出站资产转移
//	@return *taprpc.ListTransfersResponse
func ListTransfers() *taprpc.ListTransfersResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListTransfersRequest{}
	// 得到响应
	response, err := client.ListTransfers(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc ListTransfers Error: %v", err)
	}
	// 处理结果
	return response
}

// ListUtxos
//
//	@Description: 列出目标守护进程管理的 UTXO 及其持有的资产
//	@param includeLeased
//	@return *taprpc.ListUtxosResponse
func ListUtxos(includeLeased bool) *taprpc.ListUtxosResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListUtxosRequest{
		IncludeLeased: includeLeased,
	}
	// 得到响应
	response, err := client.ListUtxos(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc ListUtxos Error: %v", err)
	}
	// 处理结果
	return response
}

func NewAddr(assetId string, amt uint64) *taprpc.Addr {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	// 构建请求
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdByteSlice,
		Amt:     amt,
	}
	// 得到响应
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc NewAddr Error: %v", err)
	}
	// 处理结果
	return response
}

func QueryAddrs() {

}

func SendAsset(tapAddrs []string, feeRate uint32) *taprpc.SendAssetResponse {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.SendAssetRequest{
		TapAddrs: tapAddrs,
		FeeRate:  feeRate,
	}
	// 得到响应
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		log.Fatalf("taprpc SendAsset Error: %v", err)
	}
	// 处理结果
	return response
}

func StopDaemon() {

}

func SubscribeReceiveAssetEventNtfns() {

}

func SubscribeSendAssetEventNtfns() {

}

func VerifyProof() {

}
