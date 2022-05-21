package constants

const (
	NoteTableName              = "note"
	UserTableName              = "user"
	SecretKey                  = "secret key"
	IdentityKey                = "id"
	Total                      = "total"
	Notes                      = "notes"
	NoteID                     = "note_id"
	ApiServiceName             = "demoapi"
	VideoServiceName           = "demovideo"
	CommentServiceName         = "democomment"
	UserServiceName            = "demouser"
	MySQLDefaultDSN            = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                = "127.0.0.1:2379"
	CPURateLimit       float64 = 80.0
	DefaultLimit               = 10
	LogLevel                   = "info"
	LogRootDir                 = "./storage/logs"
	LogFileName                = "app.log"
	LogFormat                  = ""
	LogShowLine                = true
	LogMaxBackups              = 3
	LogMaxSize                 = 500
	LogMaxAge                  = 28
	LogCompress                = true
	AppEnv                     = "local"
)
