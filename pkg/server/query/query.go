package query

import (
	schemav1 "buf.build/gen/go/k8sAI-ai/k8sAI/protocolbuffers/go/schema/v1"
	"context"
	"github.com/k8sAI-ai/k8sAI/pkg/ai"
)

func (h *Handler) Query(ctx context.Context, i *schemav1.QueryRequest) (
	*schemav1.QueryResponse,
	error,
) {
	aiClient := ai.NewClient(i.Backend)
	defer aiClient.Close()

	resp, err := aiClient.GetCompletion(ctx, i.Query)
	var errMessage string = ""
	if err != nil {
		errMessage = err.Error()
	}
	return &schemav1.QueryResponse{
		Response: resp,
		Error: &schemav1.QueryError{
			Message: errMessage,
		},
	}, nil
}
