package infra

import "github.com/tietang/props/kvs"

// 应用程序启动器管理器
type BootApplication struct {
	conf kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	return &BootApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}
}

func (b *BootApplication) Start() {
	//	初始化starter
	//	安装所有的starter
	// 	启动starter
}

func (b *BootApplication)init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}
func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {

		if starter.StartBlocking() {
			//如果是最后一个可阻塞的，直接启动并阻塞
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				//如果不是，使用goroutine来异步启动，
				// 防止阻塞后面starter
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}

