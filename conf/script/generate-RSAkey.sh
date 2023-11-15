openssl req \
    -x509 \
    -nodes \
    -newkey rsa:4096 \
    -keyout manager.key \
    -out manager.crt \
    -days 3650 \
    -subj "/C=KR/ST=Seoul/L=Seoul/O=nexclipper.io/OU=Dev/CN=*"