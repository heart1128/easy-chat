package xerr

import "github.com/zeromicro/x/errors"

// New
//
//	@Description:  创建错误描述的工厂
//	@param code
//	@param msg
//	@return error
func New(code int, msg string) error {
	// 内部就是创建一个结构体吧code,msg包含子在一起，然后New返回对象
	return errors.New(code, msg)
}

func NewMsg(msg string) error {
	return errors.New(SERVER_COMMON_ERROR, msg)
}

// NewDBErr
//
//	@Description: 快速返回错误对象
//	@return error
func NewDBErr() error {
	return errors.New(DB_ERROR, ErrMsg(DB_ERROR))
}

func NewInternalErr() error {
	return errors.New(SERVER_COMMON_ERROR, ErrMsg(SERVER_COMMON_ERROR))
}
