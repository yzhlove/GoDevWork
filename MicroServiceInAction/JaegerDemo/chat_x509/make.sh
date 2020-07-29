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
