# 生成服务器端私钥
openssl genrsa -out server.key 2048

# 生成服务器端证书
openssl req -new -x509 -key server.key -out server.pem -days 3650

# 生成客户端的私钥
openssl genrsa -out client.key 2048

# 生成客户端证书
openssl req -new -x509 -key client.key -out client.pem -days 3650

# 输入项
# Country Name (2 letter code) []:CN
# State or Province Name (full name) []:Shanghai
# Locality Name (eg, city) []:Shanghai
# Organization Name (eg, company) []:JY
# Organizational Unit Name (eg, section) []:JY
# Common Name (eg, fully qualified host name) []:*.yzhdomain.com
# 其余步骤回车跳过

# ------------------------------------------------------------------------

# 生成CA证书

# 生成CA私钥
openssl genrsa -out ca.key 2048

# 创建我们自己CA的CSR，并且用自己的私钥自签署之，得到CA的身份证

openssl req -x509 -new -nodes -key ca.key -days 3650 -out ca.crt -subj "/CN=we-as-ca"

# 创建server的私钥，CSR，并且用CA的私钥自签署server的身份证

openssl genrsa -out ca_server.key 2048
openssl req -new -key ca_server.key -out ca_server.csr -subj "/CN=localhost"
openssl x509 -req -in ca_server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out ca_server.crt -days 365

# 创建client的私钥，CSR，以及用ca.key签署client的身份证

openssl genrsa -out ca_client.key 2048
openssl req -new -key ca_client.key -out ca_client.csr -subj "/CN=localhost"
openssl x509 -req -in ca_client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out ca_client.crt -days 365
