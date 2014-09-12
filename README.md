dockron
=======

dockron allows you to schedule your 'docker run's. 

Example
-------

```
# Create a wrapper script for convenience
$ cat > /usr/bin/dockron
docker run --rm -v $(which docker):/usr/bin/docker:ro -v /var/run/docker.sock:/var/run/docker.sock activestate/dockron $*
^Z
$

# Invoke your favourite container periodically (here, every minute):
$ dockron "0 * * * *" docker run ubuntu /bin/bash -c "echo Hello world"
...
```



