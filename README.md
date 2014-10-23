dockron
=======

dockron allows you to schedule your 'docker run's. 

Example
-------

```
# Create a convenient alias:
$ alias dockron="docker run --rm \
       -v $(which docker):/usr/bin/docker:ro \
       -v /var/run/docker.sock:/var/run/docker.sock \
       activestate/dockron"

# Invoke your favourite container periodically (here, every minute):
$ dockron "0 * * * *" docker run ubuntu /bin/bash -c "echo Hello world"
...
```

The first argument is the crontab-formatted repeat schedule. Rest of
the arguments should specify the entire 'docker run' command-line.

Logging
-------

The dockron container will log appropriately such that you may setup
log triggers to get notified if a command fails to run. This is
especially useful when you are managing containers using
[Papertrail](https://papertrailapp.com) and
[logspout](https://github.com/progrium/logspout), and as such demonstrates
the reason for using dockron instead of crontab on the docker host; the
scheduler is no different from that which it schedules, as all of them
are docker containers, and managed in the same way.
