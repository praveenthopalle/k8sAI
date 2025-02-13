package config

import (
	rpc "buf.build/gen/go/k8sAI-ai/k8sAI/grpc/go/schema/v1/schemav1grpc"
	schemav1 "buf.build/gen/go/k8sAI-ai/k8sAI/protocolbuffers/go/schema/v1"
	"context"
)

type Handler struct {
	rpc.UnimplementedServerConfigServiceServer
}

func (h *Handler) Shutdown(ctx context.Context, request *schemav1.ShutdownRequest) (*schemav1.ShutdownResponse, error) {
	//TODO implement me
	panic("implement me")
}
