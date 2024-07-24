package wuid

import (
	"database/sql"
	"fmt"
	"github.com/edwingeng/wuid/mysql/wuid"
	"sort"
	"strconv"
)

var w *wuid.WUID

func Init(dsn string) {
	newDB := func() (*sql.DB, bool, error) {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, false, err
		}
		return db, true, nil
	}
	w = wuid.NewWUID("default", nil)
	_ = w.LoadH28FromMysql(newDB, "wuid")
}

// GenUid
//
//	@Description: 获取一个不重复的uuid，用wuid的包
//	@param dsn
//	@return string
func GenUid(dsn string) string {
	if w == nil {
		Init(dsn)
	}
	return fmt.Sprintf("%#016x", w.Next())
}

// CombineId
//
//	@Description: A和B用户聊天，需要生成聊天消息会话，会话的id生成的方法就是排序组合两者的id
//	@param aid
//	@param bid
//	@return string  会话id
func CombineId(aid, bid string) string {
	ids := []string{aid, bid}

	sort.Slice(ids, func(i, j int) bool {
		a, _ := strconv.ParseUint(ids[i], 0, 64)
		b, _ := strconv.ParseUint(ids[j], 0, 64)
		return a < b
	})

	return fmt.Sprintf("%s_%s", ids[0], ids[1])
}
