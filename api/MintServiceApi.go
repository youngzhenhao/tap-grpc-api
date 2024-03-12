package api

import (
	"context"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"google.golang.org/grpc"
	"log"
)

// CancelBatch
//
//	@Description: 将尝试取消当前待处理批次
//	@return *mintrpc.CancelBatchResponse
func CancelBatch() *mintrpc.CancelBatchResponse {
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
	client := mintrpc.NewMintClient(conn)
	// 构建请求
	request := &mintrpc.CancelBatchRequest{}
	// 得到响应
	response, err := client.CancelBatch(context.Background(), request)
	if err != nil {
		log.Fatalf("mintrpc CancelBatch Error: %v", err)
	}
	// 处理结果
	return response
}

// FinalizeBatch
//
//	@Description: 将尝试完成当前待处理的批次
//	@return *mintrpc.FinalizeBatchResponse
func FinalizeBatch(shortResponse bool, feeRate uint32) *mintrpc.FinalizeBatchResponse {
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
	client := mintrpc.NewMintClient(conn)
	// 构建请求
	request := &mintrpc.FinalizeBatchRequest{
		ShortResponse: shortResponse,
		FeeRate:       feeRate,
		// 暂时忽略 BatchSibling
		BatchSibling: nil,
	}
	// 得到响应
	response, err := client.FinalizeBatch(context.Background(), request)
	if err != nil {
		log.Fatalf("mintrpc FinalizeBatch Error: %v", err)
	}
	// 处理结果
	return response
}

// ListBatches
//
//	 @Description: 列出提交到守护程序的批处理集，包括待处理和已取消的批处理
//		过滤器设置为空
//	 @return *mintrpc.ListBatchResponse
func ListBatches() *mintrpc.ListBatchResponse {
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
	client := mintrpc.NewMintClient(conn)
	// 构建请求
	request := &mintrpc.ListBatchRequest{}
	// 得到响应
	response, err := client.ListBatches(context.Background(), request)
	if err != nil {
		log.Fatalf("mintrpc ListBatches Error: %v", err)
	}
	// 处理结果
	return response
}

// MintAsset
//
//	@Description:  将尝试铸模请求中指定的资产集（默认为异步，以确保正确的批处理）。
//	返回的待处理批次将显示属于下一批次的其他待处理资产。此调用将阻塞，直至操作成功（资产已在批次中分期）或失败
//	@return *mintrpc.MintAssetResponse
func MintAsset(assetVersionIsV1 bool, assetTypeIsCollectible bool, name string, assetMetaData string, AssetMetaTypeIsJsonNotOpaque bool, amount uint64, newGroupedAsset bool, groupedAsset bool, groupKey string, groupAnchor string, shortResponse bool) *mintrpc.MintAssetResponse {
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
	client := mintrpc.NewMintClient(conn)
	// 指定资产版本是否为V1，默认V0
	var _assetVersion taprpc.AssetVersion
	if assetVersionIsV1 {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V1
	} else {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V0
	}
	// 指定资产类型是否为收藏品，默认正常
	var _assetType taprpc.AssetType
	if assetTypeIsCollectible {
		_assetType = taprpc.AssetType_COLLECTIBLE
	} else {
		_assetType = taprpc.AssetType_NORMAL
	}
	// AssetMeta Data
	_assetMetaDataByteSlice, _ := hex.DecodeString(assetMetaData)
	// 指定资产元数据类型
	var _assetMetaType taprpc.AssetMetaType
	if AssetMetaTypeIsJsonNotOpaque {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_JSON
	} else {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_OPAQUE
	}
	// GroupKey
	_groupKeyByteSlice, _ := hex.DecodeString(groupKey)
	// 构建请求
	request := &mintrpc.MintAssetRequest{
		Asset: &mintrpc.MintAsset{
			AssetVersion: _assetVersion,
			AssetType:    _assetType,
			Name:         name,
			AssetMeta: &taprpc.AssetMeta{
				Data: _assetMetaDataByteSlice,
				Type: _assetMetaType,
			},
			Amount:          amount,
			NewGroupedAsset: newGroupedAsset,
			GroupedAsset:    groupedAsset,
			GroupKey:        _groupKeyByteSlice,
			GroupAnchor:     groupAnchor,
		},
		ShortResponse: shortResponse,
	}
	// 得到响应
	response, err := client.MintAsset(context.Background(), request)
	if err != nil {
		log.Fatalf("mintrpc MintAsset Error: %v", err)
	}
	// 处理结果
	return response
}
