/*
* @desc:验证码参数
 */

package common

//测试
import "github.com/gogf/gf/v2/frame/g"

type CaptchaReq struct {
	g.Meta `path:"/get" tags:"验证码" method:"get" summary:"获取验证码"`
}
type CaptchaRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
}

// server:
//   address: ":8808"
//   serverRoot: "resource/public"
//   dumpRouterMap: true
//   routeOverWrite: true
//   openapiPath: "/api.json"
//   swaggerPath: "/swagger"
//   NameToUriType: 3
//   maxHeaderBytes: "20KB"
//   clientMaxBodySize: "50MB"
//   # Logging配置
//   logPath: "resource/log/server" # 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
//   logStdout: false # 日志是否输出到终端。默认为true
//   errorStack: true # 当Server捕获到异常时是否记录堆栈信息到日志中。默认为true
//   errorLogEnabled: true # 是否记录异常日志信息到日志中。默认为true
//   errorLogPattern: "error-{Ymd}.log" # 异常错误日志文件格式。默认为"error-{Ymd}.log"
//   accessLogEnabled: true # 是否记录访问日志。默认为false
//   accessLogPattern: "access-{Ymd}.log" # 访问日志文件格式。默认为"access-{Ymd}.log"

// domain: "cf"

// logger:
//   path: "resource/log/run"
//   file: "{Y-m-d}.log"
//   level: "all"
//   stdout: true

// # Database.
// database:
//   logger:
//     level: "all"
//     stdout: true
//     Path: "resource/log/sql"

//   default:
//     link: "mysql:root:1234qwer@tcp(127.0.0.1:3306)/ceshi"
//     # link: "mysql:root:admin8548!@tcp(127.0.0.1:3306)/callcenter"
//     debug: true
//     charset: "utf8mb4" #数据库编码
//     dryRun: false #空跑
//     maxIdle: 30 #连接池最大闲置的连接数
//     maxOpen: 100 #连接池最大打开的连接数
//     maxLifetime: 10 #(单位秒)连接对象可重复使用的时间长度
// fujian:
//   type: "db" #  db 数据库， file 文件
//   path: "./temp"
// gfToken:
//   cacheKey: "gfToken_"
//   timeOut: 10800
//   maxRefresh: 5400
//   multiLogin: true
//   encryptKey: "49c54195e750b04e74a8429b17896586"
//   cacheModel: "memory" #"redis"
//   excludePaths:
//     - "/api/v1/system/xxx"
//     - "/api/v1/system/xxx2"

// # Redis 配置示例
// redis:
//   # 单实例配置
//   default:
//     address: 127.0.0.1:6379
//     db: 1
//     idleTimeout: 600
//     maxActive: 100

// system:
//   notCheckAuthAdminIds: [1] #无需验证后台权限的用户id
//   dataDir: "./resource/data"
//   cache:
//     model: "memory" #缓存模式 memory OR redis
//     prefix: "tstKeFuV3Cache_" #缓存前缀

// #casbin配置
// casbin:
//   modelFile: "./resource/casbin/rbac_model.conf"
//   policyFile: "./resource/casbin/rbac_policy.csv"

// # CLI.
// gfcli:
//   gen:
//     dao:
//       - link: "mysql:root:1234qwer@tcp(127.0.0.1:3306)/callcenter"
//         tables: "cc_gong_dan_lei_bie_ti_shi"
//         removePrefix: "cc_"
//         group: "default"
//         descriptionTag: true
//         noModelComment: true
//         path: "./internal/app/kefu"
