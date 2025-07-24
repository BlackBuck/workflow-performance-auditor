FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .

# Build the Go binary
RUN go build -o /workflow-auditor .

# Final step: use minimal image for final action
FROM alpine:latest

COPY --from=builder /workflow-auditor /workflow-auditor

# Set the entrypoint for the container
ENTRYPOINT [ "/workflow-auditor" ]

