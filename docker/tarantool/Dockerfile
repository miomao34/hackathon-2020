# Dockerfile for tarantool wrapper with configs

FROM tarantool/tarantool:1.10.2

WORKDIR /usr/share/tarantool
COPY luascripts/onstartup.lua .
WORKDIR /etc/tarantool/instances.enabled
COPY luascripts/init.lua .
WORKDIR /var/lib/tarantool

EXPOSE 3301

# HEALTHCHECK --interval=10s --retries=3 \
# CMD tarantool_is_up

ENTRYPOINT ["/etc/tarantool/instances.enabled/init.lua"]
# ENTRYPOINT ["tarantool", "/opt/tarantool/init.lua"]
# ENTRYPOINT  tarantool /opt/tarantool/init.lua
