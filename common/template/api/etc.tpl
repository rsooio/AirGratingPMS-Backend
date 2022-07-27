Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
AuthRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: authentication.rpc
Auth:
  AccessSecret: 72ced10c3b65cb77a57de09b62ad15901d6433d47ae95076282dc8fc3224aaaf
  AccessExpire: 7200
