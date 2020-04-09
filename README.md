# Container from scratch

This demo is inspired by the talk by Liz Rice at Container Camp 2016.

To run the demo, start Ubuntu in a Virtual Machine and build the package:

`go build main.go`

And run with privileges:

`sudo ./main run /bin/sh`

You can check that the container has its own namespace:

`ls /`
`hostname container-host`
`hostname`

And its own processes:

`ps`

The filesystem of this example uses the Alpine Docker image:

`docker image inspect alpine:latest -f \ 
‘{{.GraphDriver.Data.UpperDir}}’`

`cp -r \ 
/var/lib/docker/overlay2/4cb707a20edd03410a4190e1d0a8402655b985a152239afefe7c8e2ed055e994/diff \
/home/fer/DataScience/rootfs`



