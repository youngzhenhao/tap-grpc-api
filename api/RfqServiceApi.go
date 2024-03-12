package api

import (
	"context"
	"github.com/lightninglabs/taproot-assets/taprpc/rfqrpc"
	"google.golang.org/grpc"
	"log"
)

func AddAssetBuyOrder() *rfqrpc.AddAssetBuyOrderResponse {
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
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.AddAssetBuyOrderRequest{}
	// 得到响应
	response, err := client.AddAssetBuyOrder(context.Background(), request)
	if err != nil {
		log.Fatalf("rfqrpc  Error: %v", err)
	}
	// 处理结果
	return response
}

func AddAssetSellOffer() *rfqrpc.AddAssetSellOfferResponse {
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
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.AddAssetSellOfferRequest{}
	// 得到响应
	response, err := client.AddAssetSellOffer(context.Background(), request)
	if err != nil {
		log.Fatalf("rfqrpc  Error: %v", err)
	}
	// 处理结果
	return response
}

func QueryRfqAcceptedQuotes() *rfqrpc.QueryRfqAcceptedQuotesResponse {
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
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.QueryRfqAcceptedQuotesRequest{}
	// 得到响应
	response, err := client.QueryRfqAcceptedQuotes(context.Background(), request)
	if err != nil {
		log.Fatalf("rfqrpc QueryRfqAcceptedQuotes Error: %v", err)
	}
	// 处理结果
	return response
}

func SubscribeRfqEventNtfns() *rfqrpc.Rfq_SubscribeRfqEventNtfnsClient {
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
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.SubscribeRfqEventNtfnsRequest{}
	// 得到响应
	response, err := client.SubscribeRfqEventNtfns(context.Background(), request)
	if err != nil {
		log.Fatalf("rfqrpc  Error: %v", err)
	}
	// 处理结果
	return &response
}
