package immodels

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ChatLogModel = (*customChatLogModel)(nil)

type (
	// ChatLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatLogModel.
	ChatLogModel interface {
		chatLogModel
	}

	customChatLogModel struct {
		*defaultChatLogModel
	}
)

// NewChatLogModel returns a model for the mongo.
func NewChatLogModel(url, db, collection string) ChatLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customChatLogModel{
		defaultChatLogModel: newDefaultChatLogModel(conn),
	}
}

// MustChatLogModel
//
//	@Description: goctl生成的创建模型，每次都要传入集合名（类似sql的表），这里固定写入比较方便
//	@param url
//	@param db
//	@return ChatLogModel
func MustChatLogModel(url, db string) ChatLogModel {
	return NewChatLogModel(url, db, "chat_log")
}
