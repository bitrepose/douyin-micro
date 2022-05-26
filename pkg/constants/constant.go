package constants

const (
	VideoCommentName  			 = "video_comment"
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
	EtcdAddressByComment                  = "127.0.0.1:2380"
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

type PostCommentMessage int8
const (
	PostCommentSuccess PostCommentMessage = 0 
	PostCommentFailure PostCommentMessage = 1
)

func (c PostCommentMessage)String()string{
	switch (c){
	case PostCommentSuccess: return "Post comment successfully"
	case PostCommentFailure: return "Fail to post comment"
	default: return "Other Errors"
	}
}
type DeleteCommentMessage int8
const (
	DeleteCommentSuccess DeleteCommentMessage = 0 
	DeleteCommentFailure DeleteCommentMessage = 1 
)
func (c DeleteCommentMessage)String()string {
	switch(c){
	case DeleteCommentSuccess: return "Delete comment successfully"
	case DeleteCommentFailure: return "Fail to delete comment"
	default: return "Other Errors"
	}
}
type GetCommentListMessage int32
const (
	GetCommentListSuccess GetCommentListMessage = 0
	GetCommentListFailure GetCommentListMessage =1
)
func (c GetCommentListMessage)String()string{
	switch(c){
	case GetCommentListSuccess: return "Get Comment List Successfully"
	case GetCommentListFailure: return "Fail to get Comment List"
	default : return "Other errors"
	}
}