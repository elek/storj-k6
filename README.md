## storj-k6

This repsitory contains module and example scripts for the [k6.io](https://k6.io/) load testing tool.

### Quick start

```
go install go.k6.io/xk6/cmd/xk6@latest
xk6 build --with github.com/elek/storj-k6=.

# change database definition in simple.js

./k6 run examples/simple.js
```

### Create image

```
docker buildx bake image --push
```