# Copyright 2023 The k8sAI Authors.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.23-alpine3.19 AS builder

ENV CGO_ENABLED=0
ARG VERSION
ARG COMMIT
ARG DATE
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /workspace/k8sAI -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" ./ 

FROM gcr.io/distroless/static AS production

LABEL org.opencontainers.image.source="https://github.com/k8sAI-ai/k8sAI" \
    org.opencontainers.image.url="https://k8sAI.ai" \
    org.opencontainers.image.title="k8sAI" \
    org.opencontainers.image.vendor='The k8sAI Authors' \
    org.opencontainers.image.licenses='Apache-2.0'

WORKDIR /
COPY --from=builder /workspace/k8sAI .
USER 65532:65532

ENTRYPOINT ["/k8sAI"]
