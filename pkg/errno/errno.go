package errno

import (
	"github.com/Ocyss/Qblog/kitex_gen/common"
)

const (
	ParamCode = 401 + iota
)

const (
	UserNotExistCode = 1001 + iota
)

const (
	TokenExpiredCode = 1101 + iota
)

const (
	ArticleNotExistCode = 2001 + iota
	ArticleUriRepeatedCode
)

var (
	Param              = &common.RequestException{Code: ParamCode, Msg: "parameter validation failed"}
	UserNotExist       = &common.RequestException{Code: UserNotExistCode, Msg: "this user does not exist"}
	ArticleNotExist    = &common.RequestException{Code: ArticleNotExistCode, Msg: "this article does not exist"}
	ArticleUriRepeated = &common.RequestException{Code: ArticleUriRepeatedCode, Msg: "Article uri is repeated a lot"}
	TokenExpired       = &common.RequestException{Code: TokenExpiredCode, Msg: "your token has expired"}
)
