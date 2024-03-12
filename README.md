# tap-api

Replace `.envExample` and `api/.envExample` with the real `.env` file

用真正的 `.env` 文件替换 `.envExample` 和 `api/.envExample`

## Official API Document

https://lightning.engineering/api-docs/api/taproot-assets/

## Taproot Assets Protocol

### AssetWallet Service
- `AnchorVirtualPsbts`: 锚定虚拟PSBTs。
- `FundVirtualPsbt`: 为虚拟PSBT提供资金。
- `NextInternalKey`: 获取下一个内部密钥。
- `NextScriptKey`: 获取下一个脚本密钥。
- `ProveAssetOwnership`: 证明资产所有权。
- `QueryInternalKey`: 查询内部密钥。
- `QueryScriptKey`: 查询脚本密钥。
- `RemoveUTXOLease`: 移除UTXO租约。
- `SignVirtualPsbt`: 签署虚拟PSBT。
- `VerifyAssetOwnership`: 验证资产所有权。

### Mint Service
- `CancelBatch`: 取消批次。
- `FinalizeBatch`: 完成批次。
- `ListBatches`: 列出批次。
- `MintAsset`: 铸造资产。

### Rfq Service
- `AddAssetBuyOrder`: 添加资产购买订单。
- `AddAssetSellOffer`: 添加资产销售报价。
- `QueryRfqAcceptedQuotes`: 查询已接受的报价。
- `SubscribeRfqEventNtfns`: 订阅RFQ事件通知。

### TapDev Service
- `ImportProof`: 导入证明。

### TaprootAssets Service
- `AddrReceives`: 地址接收。
- `BurnAsset`: 销毁资产。
- `DebugLevel`: 调试级别。
- `DecodeAddr`: 解码地址。
- `DecodeProof`: 解码证明。
- `ExportProof`: 导出证明。
- `FetchAssetMeta`: 获取资产元数据。
- `GetInfo`: 获取信息。
- `ListAssets`: 列出资产。
- `ListBalances`: 列出余额。
- `ListGroups`: 列出组。
- `ListTransfers`: 列出转账。
- `ListUtxos`: 列出UTXO。
- `NewAddr`: 新建地址。
- `QueryAddrs`: 查询地址。
- `SendAsset`: 发送资产。
- `StopDaemon`: 停止守护进程。
- `SubscribeReceiveAssetEventNtfns`: 订阅接收资产事件通知。
- `SubscribeSendAssetEventNtfns`: 订阅发送资产事件通知。
- `VerifyProof`: 验证证明。

### Universe Service
- `AddFederationServer`: 添加联盟服务器。
- `AssetLeafKeys`: 资产叶子密钥。
- `AssetLeaves`: 资产叶子。
- `AssetRoots`: 资产根。
- `DeleteAssetRoot`: 删除资产根。
- `DeleteFederationServer`: 删除联盟服务器。
- `Info`: 信息。
- `InsertProof`: 插入证明。
- `ListFederationServers`: 列出联盟服务器。
- `MultiverseRoot`: 多元宇宙根。
- `QueryAssetRoots`: 查询资产根。
- `QueryAssetStats`: 查询资产统计。
- `QueryEvents`: 查询事件。
- `QueryFederationSyncConfig`: 查询联盟同步配置。
- `QueryProof`: 查询证明。
- `SetFederationSyncConfig`: 设置联盟同步配置。
- `SyncUniverse`: 同步宇宙。
- `UniverseStats`: 宇宙统计。

### REST Endpoints

提供了一系列RESTful API端点，用于与Taproot Assets Protocol进行交互。

