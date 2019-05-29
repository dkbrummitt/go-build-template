# https://opencensus.io/service/components/collector/install/

docker run \
    -d \
    --rm \
    --interactive \
    -- tty \
    --publish 55678:55678 --publish 8888:8888 \
    --volume $(pwd)/occollector-config.yaml:/conf/occollector-config.yaml \
    occollector \
    --config=/conf/occollector-config.yaml
