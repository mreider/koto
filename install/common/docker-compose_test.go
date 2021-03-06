package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mreider/koto/install/common"
)

const (
	content = `version: "3.7"

services:
  message-hub:
    image: ghcr.io/kotollc/koto/messagehub:latest
    restart: on-failure
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - 12001:12001
    environment:
      KOTO_LISTEN_ADDRESS: :12001
      KOTO_EXTERNAL_ADDRESS: ${KOTO_EXTERNAL_ADDRESS}
      KOTO_USER_HUB_ADDRESS: ${KOTO_USER_HUB_ADDRESS}
      KOTO_DB_HOST: ${KOTO_DB_HOST}
      KOTO_DB_PORT: ${KOTO_DB_PORT}
      KOTO_DB_SSL_MODE: ${KOTO_DB_SSL_MODE}
      KOTO_DB_USER: ${KOTO_DB_USER}
      KOTO_DB_PASSWORD: ${KOTO_DB_PASSWORD}
      KOTO_DB_NAME: ${KOTO_DB_NAME}
      KOTO_S3_ENDPOINT: ${KOTO_S3_ENDPOINT}
      KOTO_S3_REGION: ${KOTO_S3_REGION}
      KOTO_S3_KEY: ${KOTO_S3_KEY}
      KOTO_S3_SECRET: ${KOTO_S3_SECRET}
      KOTO_S3_BUCKET: ${KOTO_S3_BUCKET}
    depends_on:
      - db
      - s3

  db:
    image: postgres:13
    restart: unless-stopped
    volumes:
      - ${VOLUME_DB}:/var/lib/postgresql/data/
    environment:
      POSTGRES_DB: ${KOTO_DB_NAME}
      POSTGRES_USER: ${KOTO_DB_USER}
      POSTGRES_PASSWORD: ${KOTO_DB_PASSWORD}

  s3:
    image: minio/minio:latest
    restart: on-failure
    volumes:
      - ${VOLUME_MINIO}:/data/
    environment:
      MINIO_ROOT_USER: ${KOTO_S3_KEY}
      MINIO_ROOT_PASSWORD: ${KOTO_S3_SECRET}
`
)

func TestRemoveDB(t *testing.T) {
	result := common.DockerCompose{}.RemoveServiceFromConfig(content, "db")
	assert.Equal(t, `version: "3.7"

services:
  message-hub:
    image: ghcr.io/kotollc/koto/messagehub:latest
    restart: on-failure
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - 12001:12001
    environment:
      KOTO_LISTEN_ADDRESS: :12001
      KOTO_EXTERNAL_ADDRESS: ${KOTO_EXTERNAL_ADDRESS}
      KOTO_USER_HUB_ADDRESS: ${KOTO_USER_HUB_ADDRESS}
      KOTO_DB_HOST: ${KOTO_DB_HOST}
      KOTO_DB_PORT: ${KOTO_DB_PORT}
      KOTO_DB_SSL_MODE: ${KOTO_DB_SSL_MODE}
      KOTO_DB_USER: ${KOTO_DB_USER}
      KOTO_DB_PASSWORD: ${KOTO_DB_PASSWORD}
      KOTO_DB_NAME: ${KOTO_DB_NAME}
      KOTO_S3_ENDPOINT: ${KOTO_S3_ENDPOINT}
      KOTO_S3_REGION: ${KOTO_S3_REGION}
      KOTO_S3_KEY: ${KOTO_S3_KEY}
      KOTO_S3_SECRET: ${KOTO_S3_SECRET}
      KOTO_S3_BUCKET: ${KOTO_S3_BUCKET}
    depends_on:
      - s3

  s3:
    image: minio/minio:latest
    restart: on-failure
    volumes:
      - ${VOLUME_MINIO}:/data/
    environment:
      MINIO_ROOT_USER: ${KOTO_S3_KEY}
      MINIO_ROOT_PASSWORD: ${KOTO_S3_SECRET}
`, result)
}

func TestRemoveS3(t *testing.T) {
	result := common.DockerCompose{}.RemoveServiceFromConfig(content, "s3")
	assert.Equal(t, `version: "3.7"

services:
  message-hub:
    image: ghcr.io/kotollc/koto/messagehub:latest
    restart: on-failure
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - 12001:12001
    environment:
      KOTO_LISTEN_ADDRESS: :12001
      KOTO_EXTERNAL_ADDRESS: ${KOTO_EXTERNAL_ADDRESS}
      KOTO_USER_HUB_ADDRESS: ${KOTO_USER_HUB_ADDRESS}
      KOTO_DB_HOST: ${KOTO_DB_HOST}
      KOTO_DB_PORT: ${KOTO_DB_PORT}
      KOTO_DB_SSL_MODE: ${KOTO_DB_SSL_MODE}
      KOTO_DB_USER: ${KOTO_DB_USER}
      KOTO_DB_PASSWORD: ${KOTO_DB_PASSWORD}
      KOTO_DB_NAME: ${KOTO_DB_NAME}
      KOTO_S3_ENDPOINT: ${KOTO_S3_ENDPOINT}
      KOTO_S3_REGION: ${KOTO_S3_REGION}
      KOTO_S3_KEY: ${KOTO_S3_KEY}
      KOTO_S3_SECRET: ${KOTO_S3_SECRET}
      KOTO_S3_BUCKET: ${KOTO_S3_BUCKET}
    depends_on:
      - db

  db:
    image: postgres:13
    restart: unless-stopped
    volumes:
      - ${VOLUME_DB}:/var/lib/postgresql/data/
    environment:
      POSTGRES_DB: ${KOTO_DB_NAME}
      POSTGRES_USER: ${KOTO_DB_USER}
      POSTGRES_PASSWORD: ${KOTO_DB_PASSWORD}
`, result)
}

func TestRemoveDBAndS3(t *testing.T) {
	result := common.DockerCompose{}.RemoveServiceFromConfig(content, "db")
	result = common.DockerCompose{}.RemoveServiceFromConfig(result, "s3")
	assert.Equal(t, `version: "3.7"

services:
  message-hub:
    image: ghcr.io/kotollc/koto/messagehub:latest
    restart: on-failure
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - 12001:12001
    environment:
      KOTO_LISTEN_ADDRESS: :12001
      KOTO_EXTERNAL_ADDRESS: ${KOTO_EXTERNAL_ADDRESS}
      KOTO_USER_HUB_ADDRESS: ${KOTO_USER_HUB_ADDRESS}
      KOTO_DB_HOST: ${KOTO_DB_HOST}
      KOTO_DB_PORT: ${KOTO_DB_PORT}
      KOTO_DB_SSL_MODE: ${KOTO_DB_SSL_MODE}
      KOTO_DB_USER: ${KOTO_DB_USER}
      KOTO_DB_PASSWORD: ${KOTO_DB_PASSWORD}
      KOTO_DB_NAME: ${KOTO_DB_NAME}
      KOTO_S3_ENDPOINT: ${KOTO_S3_ENDPOINT}
      KOTO_S3_REGION: ${KOTO_S3_REGION}
      KOTO_S3_KEY: ${KOTO_S3_KEY}
      KOTO_S3_SECRET: ${KOTO_S3_SECRET}
      KOTO_S3_BUCKET: ${KOTO_S3_BUCKET}
`, result)
}
