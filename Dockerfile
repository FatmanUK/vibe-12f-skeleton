# Stage 1: Create a minimal runtime image
FROM alpine:latest

# Labels for date, timestamp, and version
ARG BUILD_DATE
ARG BUILD_TIMESTAMP
ARG VERSION
LABEL org.opencontainers.image.created=${BUILD_TIMESTAMP}
LABEL org.opencontainers.image.version=${VERSION}

WORKDIR /root/
COPY main .
ARG METRICS_SERVER_PORT
EXPOSE ${METRICS_SERVER_PORT}
CMD ["./main"]
