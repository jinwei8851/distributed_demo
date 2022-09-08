package main

import (
	"context"
	"fmt"
	"gocode/distributed/log"
	"gocode/distributed/portal"
	"gocode/distributed/registry"
	"gocode/distributed/service"
	stlog "log"
)

func main() {
	err := portal.ImportTemplates() //调用模板
	if err != nil {
		stlog.Fatal(err)
	}
	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.PortalService,
		ServiceURL:  serviceAddress,
		RequiredServices: []registry.ServiceName{ //依赖的服务
			registry.LogService,
			registry.GradingService,
		},
		ServiceUpdateURL: serviceAddress + "/services",
		//HeartbeatURL:     serviceAddress + "/heartbeat",
	}

	ctx, err := service.Start(context.Background(),
		r,
		host,
		port,
		portal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err != nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down portal.")
}
