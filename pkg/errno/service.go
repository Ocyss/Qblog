package errno

import (
	"fmt"

	"github.com/Ocyss/Qblog/pkg/constants"

	"github.com/Ocyss/Qblog/kitex_gen/common"
)

const (
	ServiceCode = 500 + iota
	ServiceCpuLimitCode
)

var ServiceCpuLimitErr = &common.RequestException{Code: ServiceCpuLimitCode, Msg: fmt.Sprintf("the current server cpu exceeds the limit(%.2g)", constants.CPURateLimit)}
