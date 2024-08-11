If you've exited or disconnected from a Docker container and want to reconnect to it, you can do so using Docker's built-in commands. Here’s how you can reconnect to a running or stopped container:

### 1. **Reconnect to a Running Container**

If your container is still running, you can use `docker attach` or `docker exec` to reconnect to it.

#### Using `docker attach`

The `docker attach` command will reattach your terminal to the container's main process. Note that this might not always be ideal if the container is running a service like a web server or application that doesn’t handle terminal input/output well.

```bash
docker attach my_ubuntu_container
```

#### Using `docker exec`

A more flexible approach is using `docker exec` to start a new interactive session. This is especially useful if you want to run a new shell session:

```bash
docker exec -it my_ubuntu_container bash
```

This command starts a new Bash shell in the running container, allowing you to interact with it as if you had just started it.

### 2. **Reconnect to a Stopped Container**

If the container is stopped, you'll need to start it again before you can reconnect.

#### Start the Container

First, start the container:

```bash
docker start my_ubuntu_container
```

#### Reconnect to the Container

After starting it, you can use `docker exec` to open a new session:

```bash
docker exec -it my_ubuntu_container bash
```

### 3. **Check Container Status**

To see if your container is running or stopped, use:

```bash
docker ps -a
```

This command lists all containers, their statuses, and their IDs.

### Summary

- **Running Container**: Use `docker exec -it container_name bash` for a new interactive session.
- **Stopped Container**: Use `docker start container_name` to start it, then `docker exec -it container_name bash` to reconnect.
- **Container Status**: Check with `docker ps -a`.

By following these steps, you should be able to reconnect to your Docker container, whether it's currently running or has been stopped.
