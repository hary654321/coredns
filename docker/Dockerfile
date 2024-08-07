FROM ubuntu:20.04
#
# Setup redishoneypot
#
COPY coredns /opt/coredns/


# Start redishoneypot
WORKDIR /opt/coredns

CMD ./coredns >>/var/log/coredns/coredns.json

# CMD tail -f
# docker build -t hary654321/coredns  .

