# micro-store 服务中心
  主要提供基础服务

# 开发环境选择
    1.本地开发
    2.测试环境
    3.预发布环境
    4.生产环境
# 环境配置文件读取
>运行环境下config.toml 文件为配置文件（此文件为忽略文件）可在conf目录进行查看
# 阿里云 ACM 配置中心(任意格式)
```
  环境变量读取
  # 配置地址
  CFG_ENDPOINT=acm.aliyun.com:8080
  # 空间命名
  CFG_NAMESPACEID=a0630038-0d1c-4002-8854-0c08c47fa3e3
  # ak 密钥
  CFG_ACCESSKEY=123129234324
  # sk 密钥
  CFG_SECRETKEY=23432432432
  # 服务名
  CFG_DATA=39383232
  # 环境判断
  CFG_GROUP=dev
```
备注：配置文件和ACM只能使用其一
# 配置说明
```
[db]
   name = "store_user"
   user = "root"
   password = "password"
   host = "192.168.3.2:3306"
   charset = "utf8mb4"
   debug =  true
```
# ETCD 注册中心使用
```
  ./app --registry=etcd --registry_address=11.11.11.111:2379
```
```
  export MICRO_REGISTRY=etcd MICRO_REGISTRY_ADDRESS=11.11.11.111:2379 && ./app
```