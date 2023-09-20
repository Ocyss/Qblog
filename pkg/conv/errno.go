package conv

import (
	"errors"

	"github.com/Ocyss/Qblog/kitex_gen/common"
)

type Response struct {
	Code int32
	Msg  string
}

func ToResponse(err error) Response {
	v := new(common.RequestException)
	if errors.As(err, &v) {
		return Response{v.GetCode(), v.GetMsg()}
	}
	return Response{400, err.Error()}
}
