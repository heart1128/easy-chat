package configserver

import (
	"errors"
	"github.com/zeromicro/go-zero/core/conf"
)

// 配置服务中心

var ErrNotSetConfig = errors.New("未设置配置信息")

type OnChange func([]byte) error

type ConfigServer interface {
	Build() error
	SetOnChange(OnChange) // 配置改变的方法

	FromJsonBytes() ([]byte, error)
	//Error() error
}

type configServer struct {
	ConfigServer
	configFile string
}

func NewConfigServer(configFile string, s ConfigServer) *configServer {
	return &configServer{
		ConfigServer: s,
		configFile:   configFile,
	}
}

func (s *configServer) MustLoad(v any, onChange OnChange) error {
	//if s.ConfigServer.Error() != nil {
	//	return s.ConfigServer.Error()
	//}

	// 没有设置配置信息，文件名或者对象为空
	if s.configFile == "" && s.ConfigServer == nil {
		return ErrNotSetConfig
	}

	// 公共配置对象为空
	if s.ConfigServer == nil {
		// 使用go-zero的默认方式
		conf.MustLoad(s.configFile, v)
		return nil
	}

	// 设置配置变更之后执行的方法
	if onChange != nil {
		s.SetOnChange(onChange)
	}
	if err := s.ConfigServer.Build(); err != nil {
		return err
	}

	// 有配置
	data, err := s.ConfigServer.FromJsonBytes()
	if err != nil {
		return err
	}
	// data是json的，用go-zero方法直接转化为yaml
	return conf.LoadFromYamlBytes(data, v)
}

//func (s *configServer) Error() error {
//	return s.ConfigServer.Error()
//}
