[![Build Status](https://travis-ci.org/cloudfoundry-incubator/routing-api.svg)](https://travis-ci.org/cloudfoundry-incubator/routing-api)

# CF Routing API Server

## Installing this Repo

To clone this repo you will need to have godeps installed. You can install godeps
by running the command `go get github.com/tools/godep`.

To then install this repo you can run the following commands.

```sh
go get github.com/cloudfoundry-incubator/routing-api
cd $GOPATH/src/github.com/cloudfoundry-incubator/routing-api
godep restore
```

To install the server binary you can do

```sh
cd $GOPATH/src/github.com/cloudfoundry-incubator/routing-api
go install ./cmd/routing-api

# OR
go get github.com/cloudfoundry-incubator/routing-api/cmd/routing-api
```

## Development

### etcd

To run the tests you need a running etcd cluster on version 2.0.1. To get that do:

```sh
go get github.com/coreos/etcd
cd $GOPATH/src/github.com/coreos/etcd
git fetch --tags
git checkout v2.0.1
go install .
```

Once installed, you can run etcd with the command `etcd` and you should see the
output contain the following lines:
```
   | etcd: listening for peers on http://localhost:2380
   | etcd: listening for peers on http://localhost:7001
   | etcd: listening for client requests on http://localhost:2379
   | etcd: listening for client requests on http://localhost:4001
```

Note that this will run an etcd server and create a new directory at that location 
where it stores all of the records. This directory can be removed afterwards, or 
you can simply run etcd in a temporary directory.

## Running the API Server

### Server Configuration

#### jwt token

To run the routing-api server, a configuration file with the public uaa jwt token must be provided.
This configuration file can then be passed in with the flag `-config [path_to_config]`.
An example of the configuration file can be found under `example_config/example.yml` for bosh-lite.

To generate your own config file, you must provide a `uaa_verification_key` in pem format, such as the following:

```
uaa_verification_key: "-----BEGIN PUBLIC KEY-----

      MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDHFr+KICms+tuT1OXJwhCUtR2d

      KVy7psa8xzElSyzqx7oJyfJ1JZyOzToj9T5SfTIq396agbHJWVfYphNahvZ/7uMX

      qHxf+ZH9BL1gk9Y6kCnbM5R60gfwjyW1/dQPjOzn9N394zd2FJoFHwdq9Qs0wBug

      spULZVNRxq7veq/fzwIDAQAB

      -----END PUBLIC KEY-----"
```

This can be found in your Cloud Foundry manifest under `uaa.jwt.verification_key`

#### Oauth Clients

The Routing API uses OAuth tokens to authenticate clients. To obtain a token from UAA that grants the API client permission to register routes, an OAuth client must first be created for the API client in UAA. An API client can then authenticate with UAA using the registered OAuth client credentials, request a token, then provide this token with requests to the Routing API. 

Registering OAuth clients can be done using the cf-release BOSH deployment manifest, or manually using the `uaac` CLI for UAA. 

- For API clients that wish to register routes with the Routing API, the OAuth client in UAA must be configured with the `route.advertise` authority.
- For API clients that require admin permissions with the Routing API, the OAuth client in UAA must be configured with the `route.admin` authority.

For instructions on fetching a token, see the section "Using the API" below.

##### Configure OAuth clients in the cf-release BOSH Manifest

E.g:
```
uaa:
   clients:
      routing_api_client:
         authorities: route.advertise
         authorized_grant_type: client_credentials
         secret: route_secret
```

##### Configure OAuth clients manually using `uaac` CLI for UAA

1. Install the `uaac` CLI

   ```
   gem install cf-uaac
   ```
   
2. Get the admin client token

   ```bash
   uaac target uaa.10.244.0.34.xip.io
   uaac token client get admin # You will need to provide the client_secret, found in your CF manifest.
   ```

3. Create the OAuth client. 

   ```bash
   uaac client add routing_api_client --authorities "route.advertise" --authorized_grant_type "client_credentials"
   ```

### Starting the Server

To run the API server you need to provide all the urls for the etcd cluster, a configuration file containg the public uaa jwt key, plus some optional flags.

Example 1:

```sh
routing-api -config example_config/example.yml -port 3000 -maxTTL 60 http://etcd.127.0.0.1.xip.io:4001
```

Where `http://etcd.127.0.0.1.xip.io:4001` is the single etcd member.

Example 2:

```sh
routing-api http://etcd.127.0.0.1.xip.io:4001 http://etcd.127.0.0.1.xip.io:4002
```

Where `http://etcd.127.0.0.1.xip.io:4001` is one member of the cluster and `http://etcd.127.0.0.1.xip.io:4002` is another.

Note that flags have to come before the etcd addresses.

## Using the API

The Routing API uses OAuth tokens to authenticate clients. To obtain a token from UAA an OAuth client must first be created for the API client in UAA. For instructions on registering OAuth clients, see "Server Configuration" above.

### Using the API with the `rtr` CLI

A CLI client called `rtr` has been created for the Routing API that simplifies interactions by abstracting authentication.

- [Documentation](https://github.com/cloudfoundry-incubator/routing-api-cli)
- [Downloads](https://github.com/cloudfoundry-incubator/routing-api-cli/releases)

### Using the API manually

#### Authorization Token

To obtain an token from UAA, use the `uaac` CLI for UAA.

1. Install the `uaac` CLI

   ```
   gem install cf-uaac
   ```

2. Retrieve the OAuth token using credentials for registered OAuth client

   ```bash
   uaac token client get routing_api_client
   ```

3. Display the `access_token`, which can be used as the Authorization header to `curl` the Routing API.

   ```
   uaac context
   ```

#### `curl` Examples 

To add a route to the API server:

```sh
curl -vvv -H "Authorization: bearer [token with uaa route.advertise scope]" -X POST http://127.0.0.1:8080/v1/routes -d '[{"ip":"1.2.3.4", "route":"a_route", "port":8089, "ttl":45}]'
```

To delete a route:

```sh
curl -vvv -H "Authorization: bearer [token with uaa route.advertise scope]" -X DELETE http://127.0.0.1:8080/v1/routes -d '[{"ip":"1.2.3.4", "route":"a_route", "port":8089, "ttl":45}]'
```

To list advertised routes:
```sh
curl -vvv -H "Authorization: bearer [token with uaa route.admin scope]" http://127.0.0.1:8080/v1/routes
```

To subscribe to route changes:
```sh
curl -vvv -H "Authorization: bearer [token with uaa route.admin scope]" http://127.0.0.1:8080/v1/events
```

## Known issues

+ The routing-api will return a 404 if you attempt to hit the endpoint `http://[router host]/v1/routes/` as opposed to `http://[router host]/v1/routes`
+ The routing-api currently logs everything to the ctl log.
