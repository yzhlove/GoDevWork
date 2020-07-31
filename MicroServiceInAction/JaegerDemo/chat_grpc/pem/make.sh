# 生成服务器私钥
openssl genrsa -out server.key 2048

# 生成服务器证书
openssl req -new -x509 -key server.key -out server.pem -days 3650

# 生成客户端密钥
openssl genrsa -out client.key 2048

# 生成客户端证书
openssl req -new -x509 -key client.key -out client.pem -days 3650
