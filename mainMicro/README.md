# go-zero

## Getting started
### API
透過已撰寫好的 `.api` 生成檔案
例：透過 `user.api` 生成

先 cd 到 /api
```text
goctl api go -api <想生成檔案的.api位置> -dir <想生成的位置> -style <檔案名稱格式>
```

```text
goctl api go -api ./doc/user.api -dir ./user -style goZero
```


### model

透過已撰寫好的.sql生成檔案
例：透過user.sql生成

先 cd 到 /model
```text
goctl model mysql ddl -src= <想生成檔案的.sql位置> -dir= <想生成的位置> -style <檔案名稱格式>
```

```text
goctl model mysql ddl -src=./sql/user.sql  -dir ./user  -c --style goZero
```

`-c` 代表該 model 啟用 `redis cache` 每次查詢後會把相關結果存入 redis 中，於下次查詢時優先查找

### rpc

透過已撰寫好的 `.proto` 生成檔案

例：透過 `userctl.proto` 生成

先 cd 到 /rpc
```text
goctl rpc protoc <想生成檔案的.proto位置> --go_out= <.pd.go想生成的位置> --go-grpc_out= <_grpc.pd.go想生成的位置> --zrpc_out= <想生成的位置> -style <檔案名稱格式>
```

```text
goctl rpc protoc ./proto/userctl.proto --go_out ./userctl --go-grpc_out ./userctl --zrpc_out ./userctl --style goZero
```

## 串接
### DB 連線
於 rpc 服務中
#### add config struct
```go
type Config struct {
	rest.RestConf

        // 加上DB結構體
	DB struct {
		DsnString string
	}
}
```
#### add config yaml
```yaml
# DB
DB:
  DsnString: root:yourPwd@tcp(127.0.0.1:3306)/Demo?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
  
```
#### add svc
```go
import usermodel "go_zero/demo/model/user"

type ServiceContext struct {
	Config config.Config

	// 定義 UserModel，有多個 model 也是相同的加法
	UserModel usermodel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// db 連線
	sqlConn := sqlx.NewMysql(c.DB.DsnString)

	return &ServiceContext{
		Config: c,
		// new UserModel 物件，將連線物件注入
		UserModel: usermodel.NewUserModel(sqlConn),
	}
}

```

#### use it
in logic
```go
_, err = l.svcCtx.UserModel.Insert(l.ctx, &usermodel.User{
		...
	})
```
### zRPC 服務發現
### 測試 rpc 服務
```
grpcurl -plaintext 0.0.0.0:8089 list
```

### etcd
預設透過 etcd 做群集管理，需要預先起一個 etcd 服務在本地，詳見 docker > etcd 提供的 docker-compose 腳本。

### 修改 API Gateway 來調用 rpc 服務
`api/user/etc/user.yaml`
```yaml
# user rpc 設定
UserRpcConf:
  Etcd:
    Hosts:
      - localhost:2379
    Key: userctl.rpc # 與 userctl.yaml 中的 Key 相符
```
`api/user/internal/config/config.go`
```go
type Config struct {
	rest.RestConf

	// user rpc
	UserRpcConf zrpc.RpcClientConf
}
```
#### add svc

`api/user/internal/svc/serviceContext.go`
```go
type ServiceContext struct {
	Config config.Config
	// userCtl rpc
	UserRpc userctl.UserCtl
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userctl.NewUserCtl(zrpc.MustNewClient(c.UserRpcConf)), // new UserCtl 帶入 rpc 連線資訊
	}
}
```
#### use it
`api/user/internal/logic/registerLogic.go`
```go
func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// call rpc func
	r, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterReq{
		...
	})
	...
```

## 執行
#### 運行 rpc 服務
cd rpc/userctl
```
go run userctl.go
```
#### 運行 api服務
cd api/user
```
go run user.go
```
#### call api
路徑
```
localhost:8888/user/register
```
參數 body
```
{
    "username": "abcd",
    "password": "qwe123",
    "mobile": "0900000087",
    "gender": "male",
    "nickname": "abc"
}
```
回傳
```
{
    "result": "ok"
}
```
## 其他 goctl 指令
### 生成前端 ts
cd api
```
goctl api ts --api ./doc/user.api --dir ts
```
### 生成 Dockerfile
cd api/user
```
goctl docker --go user.go --exe user
```
cd rpc/userctl
```
goctl docker --go userctl.go --exe userctl
```

## 單元測試
### in api
#### mock rpc
cd `doing/rpc/userctl/userctl`
```
mockgen -destination userCtl_mock.go  -package userctl -source userCtl.go 
```
#### use mock in test func
```go
	...

	ctx := context.TODO()

	// 宣告 mock 物件
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRPC := userctl.NewMockUserCtl(ctrl)

	mockLogicFields := fields{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: &svc.ServiceContext{
			UserRpc: mockRPC,
		},
	}

	// 定義 mock 值
	mockRPC.
		EXPECT().
		Login(ctx, &pb.LoginReq{
			Username: "testUser",
			Password: "qwe123",
		}).
		Return(&pb.LoginResp{
			Nickname: "testUser01",
		}, nil)
```
cases 請參考 `doing/api/user/internal/logic/group1/loginLogic_test.go`
### in rpc
#### mock model ( mysql )
cd `doing/model/user`
```
mockgen -destination userModel_gen_mock.go -package user -source userModel_gen.go
```
#### use mock in test func
```go
	...

	ctx := context.TODO()

	// 宣告 mock 物件
	dbMock := user.NewMockuserModel(gomock.NewController(t))

	mockLogicFields := fields{
		ctx: ctx,
		svcCtx: &svc.ServiceContext{
			UserModel:   dbMock, // 帶入 mock 物件
		},
		Logger: logx.WithContext(ctx),
	}

	// 設定 mock 值
	dbMock.
		EXPECT().
		FindOneByName(ctx, "testUser").
		Return(&user.User{
			Password: "qwe123",
			Nickname: "testUser01",
		}, nil)
```
cases 請參考 `doing/rpc/userctl/internal/logic/loginLogic_test.go`
## tracing
### 準備
本地 docker 啟動一個 jaeger 服務
```shell
$ docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.24

```
### 啟用

於設定檔中加上設定值後，重新啟動服務

`api/user/etc/user.yaml`
```yaml
# tracing
Telemetry:
  Name: user # Service Name
  Endpoint: http://127.0.0.1:14268/api/traces
  Batcher: jaeger
  Sampler: 1.0
```
### 查看
1. 打任意一隻 API 
2. 瀏覽器前往 `http://localhost:16686/` 開啟 jeager UI
Service 選擇 `user`
3. 可以查看到第一步的 API 經過哪些 span 的鏈路

![](https://i.imgur.com/XsdQyaT.png)
