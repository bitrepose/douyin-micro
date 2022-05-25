package constants

const (
	CommentTableName             = "comment"
	VideoTableName               = "video"
	UserTableName                = "user"
	FavoriteTableName            = "user_video"
	SecretKey                    = "secret key"
	IdentityKey                  = "id"
	Total                        = "total"
	Notes                        = "notes"
	NoteID                       = "note_id"
	ApiServiceName               = "demoapi"
	VideoServiceName             = "demovideo"
	CommentServiceName           = "democomment"
	UserServiceName              = "demouser"
	MySQLDefaultDSN              = "root:SUDAcs647#SQL@tcp(139.196.145.51:3306)/douyin?parseTime=True&loc=Local"
	EtcdAddress                  = "127.0.0.1:2379"
	CPURateLimit         float64 = 80.0
	DefaultLimit                 = 10
	LogLevel                     = "info"
	LogRootDir                   = "./storage/logs"
	LogFileName                  = "app.log"
	LogFormat                    = ""
	LogShowLine                  = true
	LogMaxBackups                = 3
	LogMaxSize                   = 500
	LogMaxAge                    = 28
	LogCompress                  = true
	AppEnv                       = "local"
	MinioEndpoint                = "139.196.145.51:19000"
	MinioAccessKeyId             = "5yd6ar46sosoNbqR"
	MinioSecretAccessKey         = "pBwgu2sEFfoSh8bHFwTuJFaDCdHceo86"
	MinioUseSSL                  = false
	MinioVideoBucketName         = "douyin-video"
)
