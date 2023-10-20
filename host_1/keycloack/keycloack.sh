USERNAME="${USERNAME:=admin}" 
PASSWORD="${PASSWORD:=password}" 
KEYCLOAK_PORT="${KEYCLOAK_PORT:=8443}" 

docker run \
  --name keycloak \
  -e KEYCLOAK_ADMIN=$USERNAME \
  -e KEYCLOAK_ADMIN_PASSWORD=$PASSWORD \
  -e KC_HTTPS_CERTIFICATE_FILE=/opt/keycloak/conf/server.crt.pem \
  -e KC_HTTPS_CERTIFICATE_KEY_FILE=/opt/keycloak/conf/server.key.pem \
  -v $PWD/cert/keycloak-server.crt.pem:/opt/keycloak/conf/server.crt.pem \
  -v $PWD/cert/keycloak-server.key.pem:/opt/keycloak/conf/server.key.pem \
  -e KEYCLOAK_HTTPS_PORT=$KEYCLOAK_PORT \
  -p $KEYCLOAK_PORT:$KEYCLOAK_PORT \
  quay.io/keycloak/keycloak:22.0.3 \
  start-dev