FROM cgr.dev/chainguard/go:1.19 AS build

WORKDIR /code
COPY . .
RUN make build

FROM cgr.dev/chainguard/glibc-dynamic AS chnsd
COPY --from=build /code/chnsd /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/chnsd"]
