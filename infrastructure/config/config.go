package config

import (
	"encoding/json"
	"fmt"

	"github.com/caarlos0/env/v6"
)

var Cfg *Config

func init() {
	//Cfg = getConfig()
	Cfg = new(Config)
	if err := env.Parse(Cfg); err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	b, _ := json.MarshalIndent(Cfg, "", "  ")
	fmt.Println("config:", string(b))
}

type Config struct {
	//服务相关配置
	ServerName   string `env:"SERVER_NAME" envDefault:"mop-open-api"`
	HttpPort     string `env:"HTTP_PORT" envDefault:"8080"` //服务端口
	GrpcPort     string `env:"GRPC_PORT" envDefault:"9090"`
	ConsulAddr   string `env:"CONSUL_ADDR" envDefault:""`
	ConsulHttpIp string `env:"CONSUL_HTTP_IP" envDefault:""`
	ConsulPort   string `env:"CONSUL_PORT" envDefault:"8500"`
	ConsulTag    string `env:"CONSUL_TAG" envDefault:"mop-finstore"`
	ReqType      string `env:"REQ_TYPE" envDefault:"http"`
	Mode         string `env:"MODE" envDefault:"debug"`           //路由模式 test、realese、debug
	LogMode      string `env:"LOG_MODE" envDefault:"debug"`       //日志级别 默认debug
	OpenLogColor bool   `env:"OPEN_LOG_COLOR" envDefault:"false"` //是否开启日志颜色
	OpenMonitor  bool   `env:"OPEN_MONITOR" envDefault:"false"`   //是否开启监控

	//数据库配置相关
	MongoURL string `env:"MONGO_URL" envDefault:"mongo-cluster.storage.svc.cluster.local"` //mongo url
	//for local test
	//MongoURL      string `env:"MONGO_URL" envDefault:"mongodb://127.0.0.1:27017"` //mongo url
	DBName    string `env:"MONGO_DB_NAME" envDefault:"mop-open-api"`                     //数据库名
	RedisAddr string `env:"REDIS_ADDR" envDefault:"redis://proxy.redis-cluster:6379/14"` //redis url
	//for local test
	//RedisAddr     string `env:"REDIS_ADDR" envDefault:"redis://user:finochat@redis1:6379/0"`    //redis url
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""` //pwd

	//应用市场配置
	AppManageHost         string `env:"APP_MANAGE_HOST" envDefault:"http://mop-app-manage-svr:8080/api/v1/mop/finstore"`
	GetBindingIdURL       string `env:"GET_BINDING_ID_URL" envDefault:"/openapi/bindingId"`
	GetOrganIdBySdkKeyURL string `env:"GET_ORGAN_ID_SDK_URL" envDefault:"/organ-id"`
	GetSecretURL          string `env:"GET_SECRET_URL" envDefault:"/openapi/secret"`
	GetBindInfoURL        string `env:"GET_BINDING_INFO_URL" envDefault:"/openapi/bindingInfo"`
	GetAppListURL         string `env:"GET_APP_LIST_URL" envDefault:"/openapi/app/list"`
	GetAppInfoURL         string `env:"GET_APP_INFO_URL" envDefault:"/openapi/apps/info"`
	GetAppVerInfoURL      string `env:"GET_APP_VER_INFO_URL" envDefault:"/openapi/appver"`
	SearchAppURL          string `env:"SEARCH_APP_URL" envDefault:"/openapi/app/search"`
	ControlManagerHost    string `env:"CONTROL_MANAGER_HOST" envDefault:"http://mop-control-manager:8080"`
	ControlManagerInfoURL string `env:"CONTROL_MANAGER_INFO_URL" envDefault:"/api/v1/mop/mop-control-manager/mopSDK/config/info"`
	BasicPackHost         string `env:"BASIC_PAC_HOST" envDefault:"http://mop-basic-pack-svr:8080"`
	GetLatestBasicURL     string `env:"GET_LATEST_BASIC_URL" envDefault:"/api/v1/mop/mop-basic-pack-svr/latestBasic"`
	VerifyCodeHost        string `env:"VERIFYCODE_HOST" envDefault:"http://mop-verify-code-gateway:8080"`
	//初始化服务
	CommonConfURL string `env:"COMMON_CONF_URL" envDefault:"http://mop-private-init-server:8080/api/v1/mop/mop-private-init-server/common/config"`
	//license服务
	LicenseInfoURL string `env:"LICENSE_INFO_URL" envDefault:"http://mop-license-checker:8080/api/v1/mop/mop-license-checker/license"`

	//规则引擎配置相关
	RuleEngHost            string `env:"RULE_ENG_HOST" envDefault:"http://mop-rule-engine-svr:8080"`
	RuleEngGrayAppVerUrl   string `env:"RE_GRAY_APPVER_URL" envDefault:"/api/v1/mop/ruleEngine/gray/appVer"`
	RuleEngGrayBatchGetApp string `env:"RE_GRAY_BATCH_APP_URL" envDefault:"/api/v1/mop/rule-engine/gray/batch/app"`
	//apm上报
	ApmReportHost          string `env:"APM_REPORT_HOST" envDefault:"http://mop-data-report:8080"`
	CommonApmReportURL     string `env:"COM_APM_REPORT_URL" envDefault:"/api/v1/mop/data-report/reportmsg/report/commonMsg/report"`
	PrivateApmReportURL    string `env:"PRI_APM_REPORT_URL" envDefault:"/api/v1/mop/data-report/reportmsg/report"`
	RegulatorsApmReportURL string `env:"REGULATORS_APM_REPORT_URL" envDefault:"/api/v1/mop/data-report/report-msg/report/regulators"`

	//业务需要相关配置
	TimestampExpire int    `env:"TIMESTAMP_EXPIRE" envDefault:"60"`
	SignTest        string `env:"SIGN_TEST" envDefault:""`
	//加密算法
	EncryType string `env:"ENTRY_TYPE" envDefault:""`
	//redis
	RedisSentinelAddr     string `env:"REDIS_SENTINEL_ADDR" envDefault:"redis://redis-cluster-redis-cluster-sentinel.redis-cluster:26379/"`
	RedisMasterName       string `env:"REDIS_MASTER_NAME" envDefault:"mymaster"`
	RedisDatabase         int    `env:"REDIS_INDEX" envDefault:"11"`
	MaxRedisClientNum     int    `env:"MAX_REDIS_CLIENT_NUM" envDefault:"60"`
	RedisMode             string `env:"REDIS_MODE" envDefault:"single"`
	RedisSentinelPassword string `env:"REDIS_SENTINEL_PASSWORD" envDefault:""`

	//frq
	OpenFrq                bool   `env:"OPEN_FRQ" envDefault:"false"`
	FrqSpan                string `env:"FRQ_SPAN" envDefault:"sec"` //frq的粒度 sec-秒，min-分钟，hour-小时，day-天，week-周，mon-月，year-年
	FrqType                string `env:"FRQ_TYPE"  envDefault:"incr"`
	FrqLimit               int64  `env:"FRQ_LIMIT" envDefault:"99999999"` //frq 限制数量
	FrqGetAppInfoLimit     int64  `env:"FRQ_GET_APP_INFO_LIMIT" envDefault:"99999999"`
	FrqGetLatestBasicLimit int64  `env:"FRQ_GET_LATEST_BASIC_LIMIT" envDefault:"99999999"`

	PubEnv string `env:"PUB_ENV" envDefault:"mop-uat"`

	KafkaAddr string `env:"KAFKA_ADDR" envDefault:"kafka-service.kafka:9093"`
	//for local test
	//KafkaAddr             string `env:"KAFKA_ADDR" envDefault:"kafka:9092"`
	KafkaVersion        string `env:"KAFKA_VERSION" envDefault:"2.3.0"`
	KafkaMopReportTopic string `env:"DATA_REPORT_TOPIC" envDefault:""`
	//for local test
	//KafkaMopReportTopic   string `env:"DATA_REPORT_TOPIC" envDefault:"mop_data_report"`
	OpenPrivateDataReport bool `env:"OPEN_PRIVATE_DATA_REPORT" envDefault:"true"`
	OpenCommonDataReport  bool `env:"OPEN_COMMON_DATA_REPORT" envDefault:"true"`

	SkyWalkingUrl        string `env:"SKYWALKING_URL" envDefault:"127.0.0.1:11800"`
	SkyWalkingEnable     bool   `env:"SKYWALKING_ENABLE" envDefault:"false"`
	SkyWalkingPartitions uint32 `env:"SKYWALKING_PARTITIONS" envDefault:"1"`
	DBMode               string `env:"DB_MODE" envDefault:"mongo"`
	SqlAddr              string `env:"SQL_ADDR" envDefault:"mongo"`
}
