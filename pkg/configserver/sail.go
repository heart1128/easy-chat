package configserver

import (
	"encoding/json"
	"fmt"
	"github.com/HYY-yu/sail-client"
)

type Config struct {
	ETCDEndpoints  string `toml:"etcd_endpoints"` // 逗号分隔的ETCD地址，0.0.0.0:2379,0.0.0.0:12379,0.0.0.0:22379
	ProjectKey     string `toml:"project_key"`
	Namespace      string `toml:"namespace"`
	Configs        string `toml:"configs"`          // 逗号分隔的 config_name.config_type，如：mysql.toml,cfg.json,redis.yaml，空代表不下载任何配置
	ConfigFilePath string `toml:"config_file_path"` // 本地配置文件存放路径，空代表不存储本都配置文件
	LogLevel       string `toml:"log_level"`        // 日志级别(DEBUG\INFO\WARN\ERROR)，默认 WARN
}

type Sail struct {
	*sail.Sail
	sail.OnConfigChange
	cfg *Config
}

func NewSail(cfg *Config) *Sail {
	//s := sail.New(&sail.MetaConfig{
	//	ETCDEndpoints:  cfg.ETCDEndpoints,
	//	ProjectKey:     cfg.ProjectKey, // sail项目管理的key
	//	Namespace:      cfg.Namespace,
	//	Configs:        cfg.Configs,
	//	ConfigFilePath: cfg.ConfigFilePath,
	//	LogLevel:       cfg.LogLevel,
	//})
	return &Sail{cfg: cfg}
}

// FromJsonBytes
//
//	@Description: 冲配置中心读取配置文件，转为json
//	@receiver s
//	@return []byte
//	@return error
func (s *Sail) FromJsonBytes() ([]byte, error) {
	if err := s.Pull(); err != nil {
		return nil, err
	}

	return s.fromJsonBytes(s.Sail)
}

// fromJsonBytes
//
//	@Description: 重新加载配置
//	@receiver s
//	@return []byte
//	@return error
func (s *Sail) fromJsonBytes(sail *sail.Sail) ([]byte, error) {
	v, err := s.MergeVipers()
	if err != nil {
		return nil, err
	}

	data := v.AllSettings()

	return json.Marshal(data)
}

func (s *Sail) Build() error {
	var opts []sail.Option

	if s.OnConfigChange != nil {
		opts = append(opts, sail.WithOnConfigChange(s.OnConfigChange))
	}

	//s := sail.New(&sail.MetaConfig{
	//	ETCDEndpoints:  s.cfg.ETCDEndpoints,
	//	ProjectKey:     s.cfg.ProjectKey, // sail项目管理的key
	//	Namespace:      s.cfg.Namespace,
	//	Configs:        s.cfg.Configs,
	//	ConfigFilePath: s.cfg.ConfigFilePath,
	//	LogLevel:       s.cfg.LogLevel,
	//})
	return nil
}

func (s *Sail) setOnChange(change OnChange) {
	s.OnConfigChange = func(configFileKey string, sail *sail.Sail) {
		data, err := s.fromJsonBytes(sail)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err = change(data); err != nil {
			fmt.Println("OnChange err ", err)
		}
	}
}
