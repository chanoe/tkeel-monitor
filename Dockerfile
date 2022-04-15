FROM alpine:3.13
ENV PLUGIN_ID=tkeel-monitor
ADD monitor /
CMD ["/monitor"]