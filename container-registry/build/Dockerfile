FROM golang:1.19 as build_context

ENV OUTDIR=/out
RUN mkdir -p ${OUTDIR}/usr/local/bin/

WORKDIR /workspace
ADD . /workspace/

RUN make build
RUN mv /workspace/dist/cr-server ${OUTDIR}/usr/local/bin/

FROM alpine:3.11

COPY --from=build_context /out/ /

WORKDIR /

EXPOSE 5000
CMD ["sh"]