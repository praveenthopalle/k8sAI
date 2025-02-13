package query

import rpc "buf.build/gen/go/k8sAI-ai/k8sAI/grpc/go/schema/v1/schemav1grpc"

type Handler struct {
	rpc.UnimplementedServerQueryServiceServer
}
