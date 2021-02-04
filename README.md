# OmnIS Client

[![License](https://img.shields.io/badge/license-Apache%20license%202.0-blue.svg)](https://github.com/omnis-org/omnis-client/blob/main/LICENSE)

OmnIS Client is part of the OmnIS project. It allows the transfer of information from a host to the server.

![omnis_logo](./omnis_logo.png)



## How to build ?


```bash
cd build
make build
# That will generate a binary omnis-client
```



## Create configuration file

You have examples of configuration file in build/testdata/example.json :

```
{
    "server" : {
        "timeout" : 5,                  # Max time before stop send informations
        "serverIp" : "127.0.0.1",      # The IP address of the omnis server service
        "serverPort" : 4320,           # The port of the omnis server service
        "tls": true,                    # Is TLS activated ?
        "insecureSkipVerify": false   # Check if certificate is valid
    },
    "client" : {
        "location" : "Paris",           # The location of the client (physical place)
        "perimeter" : "Network1",       # The perimeter of the client (virtual place)
        "sendTime" : 60                # The time between each sending of information
    }

}
```

## How to launch ?

Lauch the client with the created config file :

```bash
./omnis-client testdata/example.json
```


## Licensing

OmnIS Client is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/omnis-org/omnis-client/blob/main/LICENSE) for the full license text.
