# M3DB

M3DB is an open source project recently released by Uber that stores metrics over a simple
HTTP API. You can write to it with a simple CURL command using JSON or hookup Prometheus to
it. This will take a look at a general M3DB setup that allows you to hook up prometheus to it
when your storage backend sits on a different server than Prometheus.

This was derived from two posts on the m3db github site that you can read (here)[1] and (here)[2].

## Setting up M3DB storage backend and ETCD

First we will create what the M3DB documentation calls a "Seed Node". A seed node is a superset of
the storage node that in addition to storging your data runs ETCD which is responsible for keeping
track of all storage nodes in your cluster. In this case ETCD will keep track of the node it's running
on since in this setup we are running them side by side.

Another thing to note is that the software with it's default config needs at least 4GB to run even
at a very small scale which in my case is just scraping a few hundred metrics.

OK now to actually get started you will want to pull down the storage node config from (here)[3] and change
a few things.

```
  config:
    service:
      env: default_env
      zone: embedded
      service: m3db
      cacheDir: /var/lib/m3kv
      etcdClusters:
        - zone: embedded
          endpoints:
            - <reachable_ip>:2379
    seedNodes:
      initialCluster:
        - hostID: admin
          endpoint: http://<reachable_ip>:2380
```

Where reachable_ip is an IP that your prometheus instance on another box will still be able to talk to. In
my case this was just a public IP of my VM running on digital ocean. Additionally the hostID needs to match
my systems hostname this can be changed but I used the default config below which tells it to pull the hostID
of a node from the hostname.

```
  hostID:
    resolver: hostname
```

After you have setup your config and edited to your liking run the following command to startup your instance.
Make sure to change the config path for your setup.

```
#!/bin/bash

docker run \
        --net host \
        --name m3db \
        -v $(pwd)/config/m3dbnode.yml:/etc/m3dbnode/m3dbnode.yml \
        -v /var/lib/m3db:/var/lib/m3db \
        -d \
        quay.io/m3/m3dbnode:latest
```

After this has started up you will need to initialize the database as well. Again make sure to
change out the id and hostname where it says admin for the hostID of your instance.

```
#!/bin/bash

curl -X POST localhost:7201/api/v1/placement/init -d '{
    "num_shards": 1024,
    "replication_factor": 1,
    "instances": [
        {
            "id": "admin",
            "isolation_group": "nyc",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "142.93.19.169:9000",
            "hostname": "admin",
            "port": 9000
        }
    ]
}'

sleep 10

curl -X POST localhost:7201/api/v1/namespace -d '{
  "name": "default",
  "options": {
    "bootstrapEnabled": true,
    "flushEnabled": true,
    "writesToCommitLog": true,
    "cleanupEnabled": true,
    "snapshotEnabled": true,
    "repairEnabled": false,
    "retentionOptions": {
      "retentionPeriodDuration": "720h",
      "blockSizeDuration": "12h",
      "bufferFutureDuration": "1h",
      "bufferPastDuration": "1h",
      "blockDataExpiry": true,
      "blockDataExpiryAfterNotAccessPeriodDuration": "5m"
    },
    "indexOptions": {
      "enabled": true,
      "blockSizeDuration": "12h"
    }
  }
}'
```

After running this you you can tail the logs and you should see everything running smoothly. For a much more indepth
version of the above look at this (post)[1] which I highly recommend.

## Setup the M3 Coordinator

Now that we have the backend running and happy we need to setup the M3 Coordinator. This is a sidecar for prometheus
that will manage making the actual API calls needed to get data from the storage node.

To get started we just need to fill in a few config values again. Here is the config that I used you will only need
to change the reachable_ip provided you followed everything above.

```
listenAddress:
  type: "config"
  value: "0.0.0.0:7201"

metrics:
  scope:
    prefix: "coordinator"
  prometheus:
    handlerPath: /metrics
    listenAddress: 0.0.0.0:7203 # until https://github.com/m3db/m3/issues/682 is resolved
  sanitization: prometheus
  samplingRate: 1.0
  extended: none

clusters:
   - namespaces:
       - namespace: default
         retention: 48h
         type: unaggregated
     client:
       config:
         service:
           env: default_env
           zone: embedded
           service: m3db
           cacheDir: /var/lib/m3kv
           etcdClusters:
             - zone: embedded
               endpoints:
                 - <reachable_ip>:2379
       writeConsistencyLevel: majority
       readConsistencyLevel: one
       writeTimeout: 10s
       fetchTimeout: 15s
       connectTimeout: 20s
       writeRetry:
         initialBackoff: 500ms
         backoffFactor: 3
         maxRetries: 2
         jitter: true
       fetchRetry:
         initialBackoff: 500ms
         backoffFactor: 2
         maxRetries: 3
         jitter: true
       backgroundHealthCheckFailLimit: 4
       backgroundHealthCheckFailThrottleFactor: 0.5
```

Once this is written to disk you can run the following command to get it started.

```
#!/bin/bash

docker run \
        --net host \
        --name m3coordinator \
        -v $(pwd)/m3coordinator.yml:/etc/m3coordinator/m3coordinator.yml \
        -d \
        quay.io/m3/m3coordinator:latest
```

## Setup Prometheus

Now that you have all the M3 pieces setup you just need to standup Prometheus and let it start scraping
some targets.

Here is the config that I used which will write to M3DB metrics that it scrapes from itself for easy testing.

```
global:
  scrape_interval:     15s
  evaluation_interval: 30s

remote_read:
  - url: "http://0.0.0.0:7201/api/v1/prom/remote/read"
    # To test reading even when local Prometheus has the data
    read_recent: true

remote_write:
  - url: "http://0.0.0.0:7201/api/v1/prom/remote/write"

scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'prometheus'

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ['0.0.0.0:9090']
```

You can start this up with the following script

```
#!/bin/bash

docker run \
        --net host \
        -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
        -d \
        --name prometheus \
        prom/prometheus
```

Now you should be able to go to <your_ip>:9090 and see metrics that are being ingested. What is super cool if you can
wait a few minutes destroy your container and start it again and all your metrics still exist! 

## Notes

I used host networking to make things easier but you can easily containerize all the networking if you wish. The official
M3DB and Prometheus docs provide all the info you need around port mappings. This guide mostly came about because I started
with the single node install on the docs and tried to run prometheus against it not knowing that the node the simple single
install sets up uses localhost for the instance which causes all sorts of run problems and strange errors.

[1]: https://m3db.github.io/m3/how_to/cluster_hard_way/
[2]: https://m3db.github.io/m3/integrations/prometheus/
[3]: https://github.com/m3db/m3/blob/master/src/dbnode/config/m3dbnode-cluster-template.yml
