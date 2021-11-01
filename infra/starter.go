package infra

import "github.com/tietang/props/kvs"

const (
	KeyProps = "_conf"
)

// StarterContext 基础资源上下结构体
type StarterContext map[string]interface{}

// Props 这里加载配置文件
func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没被初始化")
	}
	return p.(kvs.ConfigSource)
}

// Starter 基础资源启动器接口
type Starter interface {
	// Init 1.系统启动，初始化一些基础资源
	Init(StarterContext)
	// Setup 2. 系统基础资源的安装
	Setup(StarterContext)
	// Start 3. 启动基础资源
	Start(StarterContext)
	// StartBlocking 启动器是否可阻塞
	StartBlocking() bool
	// Stop 4. 资源停止和销毁
	Stop(StarterContext)
}

var _ Starter = new(BaseStarter)

//基础空启动器实现,为了方便资源启动器的代码实现
type BaseStarter struct {
}



func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

//启动器注册器
type starterRegister struct {
	starters []Starter
}

//启动器注册
func (r *starterRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
}

func (r *starterRegister) AllStarters() []Starter {
	return r.starters
}

var StarterRegister *starterRegister = new(starterRegister)

func Register(s Starter) {
	StarterRegister.Register(s)
}