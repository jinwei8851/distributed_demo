package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
	//扩展项目
	RequiredServices []ServiceName //依赖的service
	ServiceUpdateURL string        //暴露URL 服务注册告诉url动态接收更新
}
type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
