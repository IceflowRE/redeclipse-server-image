# Red Eclipse Server Image

![maintained](https://img.shields.io/badge/maintained-yes-brightgreen.svg)
[![github actions](https://img.shields.io/github/actions/workflow/status/IceflowRE/redeclipse-server-image/build.yml)](https://github.com/IceflowRE/redeclipse-server-image/actions/workflows/build.yaml)  
[![GHCR](https://img.shields.io/badge/GHCR-GHCR?style=social&logo=github)](https://github.com/IceflowRE/redeclipse-server-image/pkgs/container/redeclipse-server) 
[![DockerHub](https://img.shields.io/badge/DockerHub-DockerHub?style=social&logo=docker)](https://hub.docker.com/r/iceflower/redeclipse-server)
[![GitHub](https://img.shields.io/badge/GitHub-GitHub?style=social&logo=github)](https://github.com/IceflowRE/redeclipse-server-image)

---

An easy-to-use container image for a [Red Eclipse](https://www.redeclipse.net) Server.

## Supported tags and respective `Dockerfile` Links

* [`dev`, `master`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_2_0_0) - [Source Commit](https://github.com/redeclipse/base/tree/master)  
  The development branch.
* [`latest`, `2.0.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_2_0_0) - [Source Commit](https://github.com/redeclipse/base/tree/v2.0.0)
* [`stable`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/stable)  
  The latest legacy version (stable branch).
* [`1.6-impulse`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/ae3c6fa1585fc7c0375e74ed6f618b5fc791a109)  
  The latest version with the 1.6. impulse system.
* [`1.6.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/v1.6.0)
* [`1.5.8`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.8)
* [`1.5.6`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.6)
* [`1.5.5`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.5)
* [`1.5.3`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_3) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.3)
* [`1.5.2`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_0) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.2)  
  Bundles 1.5.3 maps, as the 1.5.2 maps are no longer available.
* [`1.5.1`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_0) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.1)  
  Bundles 1.5.3 maps, as the 1.5.1 maps are no longer available.
* [`1.5.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_5_0) - [Source Commit](https://github.com/redeclipse/base/tree/v1.5.0)  
  Bundles 1.5.3 maps, as the 1.5.0 maps are no longer available.
* [`1.4.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_4_0) - [Source Commit](https://github.com/redeclipse/oldsvnre/tree/v1.4.0)
* [`1.3.1`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_0_0) - [Source Commit](https://github.com/redeclipse/oldsvnre/tree/v1.3.0)
* [`1.2.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_0_0) - [Source Commit](https://github.com/redeclipse/oldsvnre/tree/v1.2.0)
* [`1.1.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_0_0) - [Source Commit](https://github.com/redeclipse/oldsvnre/tree/v1.1.0)
* [`1.0.0`](https://github.com/IceflowRE/redeclipse-server-image/blob/main/Dockerfile_1_0_0) - [Source Commit](https://github.com/redeclipse/oldsvnre/tree/v1.0.0)

## Quick reference

* **Registries**:  
  GitHub Container Registry (recommended): [`ghcr.io/iceflower/redeclipse-server`](ghcr.io/iceflower/redeclipse-server)  
  Docker Hub: [`docker.io/iceflower/redeclipse-server`](https://hub.docker.com/r/iceflower/redeclipse-server)

* **Built architectures**:  
  `amd64`, `arm64v8`

* **Where to file issues**:  
  [https://github.com/IceflowRE/redeclipse-server-image/issues](https://github.com/IceflowRE/redeclipse-server-image/issues)

## Usage

Using [Docker Compose](https://docs.docker.com/compose/) is the easiest way to manage your Red Eclipse server container.

1. Create a file named `compose.yaml`. You can copy the template below or use the provided [compose-template.yaml](https://github.com/IceflowRE/redeclipse-server-image/blob/main/compose-template.yaml).
2. Replace all `<variable_placeholders>` (including the angle brackets) with your actual settings.

```yml
services:
  # <service_name>: The identifier used to manage this service (e.g., myserver, master, re_2_0)
  <service_name>:
    # <tag>: The image tag you want to use (e.g., master, latest, 2.0.0)
    image: ghcr.io/iceflower/redeclipse-server:<tag>
    ports:
      # <serverport>: Must match the port defined in Red Eclipse's servinit.cfg
      - "28801:28801/udp"
      # <serverport + 1>: The server port plus one (e.g., if serverport is 28801, use 28802)
      - "28802:28802/udp"
    restart: unless-stopped
    volumes:
      # <home_dir>: Path to the RE home/config directory on your host system (e.g., /home/iceflower/re-master/home)
      - <home_dir>:/re-server-config/home:ro
      # <package_dir>: Optional path to your local package directory for custom maps. Remove this section if not needed.
      - <package_dir>:/re-server-config/package:ro
      # <sauerbraten_dir>: Only required for legacy versions (< 2.0.0). Remove this block if using v2.0.0 or master.
      - <sauerbraten_dir>:/re-server-config/sauer:ro
```

To pull/start/stop a specific service add the service name to the end (`docker compose pull myserver`) otherwise all services will be affected.

* Pull the latest docker image from Docker Hub for all defined services  
  `docker compose pull`

* Start/ Restart container  
  `docker compose up -d`

* Shutdown and wait a maximum of 10 minutes before forcing it  
  `docker compose stop --time=600`

#### Multiple Server

Copy and paste the whole service section above and change the values (service name, port, home directory, etc.)

### An image is not available for my architecture

If an image is not natively available for your architecture, first verify if Red Eclipse can be compiled on your target platform. If it can, follow these steps:

1. Copy the appropriate `Dockerfile_X_X_X` file corresponding to your version from this repository and place it next to your `compose.yaml`:

|   Version Range    |     Dockerfile     |
| :----------------: | :----------------: |
| `2.0.0` - `master` | `Dockerfile_2_0_0` |
| `1.5.3` - `stable` | `Dockerfile_1_5_3` |
| `1.5.0` - `1.5.3`  | `Dockerfile_1_5_0` |
|      `1.4.0`       | `Dockerfile_1_4_0` |
| `1.0.0` - `1.3.1`  | `Dockerfile_1_0_0` |

2. Create a file named `.dockerignore` next to the `Dockerfile` with the content `**`.

3. Edit your `compose.yaml` and replace the `image: ...` part in the service you want to build with

```yml
build:
  # <dockerfile>: name of the dockerfile, usable for the chosen reference below
  dockerfile: <dockerfile>
  args:
    # <reference>: can be any git reference (branch name, tag, SHA) (e.g. master, 2.0.0, ...)
    RE_REF: <reference>
```

4. Use `docker compose build` to build the custom image.

> [!NOTE]
> Depending on your specific host environment, minor modifications to the `Dockerfile` instructions may be necessary to complete the build successfully.

## Update server automatically

### Linux

1. Create a script named `update-server.sh` next to your `compose.yaml`

```bash
#!/bin/bash -e

# Adjust this path to match your deployment directory
cd /home/iceflower/re/
docker compose stop -t 600
# Uncomment next line if using a custom local architecture build
# docker compose build
docker compose pull
docker compose up -d
```

2. Make the script executable:

```shell
chmod +x update-server.sh
```

3. Configure a cron job (`crontab -e`) to execute the updates on a schedule. For example, to run every day at 03:00:

```cron
0 3 * * * /home/iceflower/re/update-server.sh > /home/iceflower/re/cron.log 2>&1
```

## Image updater

The repository include a small tool to check for updates and get the image digest and support the image creation in the CI.

Building:

```shell
cd go-image-updater
go build -x -o updater ./cmd/updater/
```

For usage options see `--help`.


## License

Copyright 2017-present Iceflower S

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

*Red Eclipse is licensed under [THE RED ECLIPSE LICENSE](https://github.com/redeclipse/base/blob/master/doc/license.txt). Red Eclipse game binary and map assets are bundled directly during container provisioning and are independent properties of their respective authors.*
