# Zvon Backend
Zvon is a software that automatically rings a bell when the issue gets closed.
Zvon backend acts as a web hook proxy. You set the backend as the issues' web hook on gitlab. Once the gitlab
notifies backend that the issue was closed, the server remembers that the bell should be banged. The client then 
periodically aks the server weather the bell should be banged.

## API
All API URLs are prefixed with `/api/v1`. The possible paths are
- `/webhook/issues`
    - `POST` The url to which the gitlab makes issues web hook request with data. Read
    more in [gitlab issues web hook documentation](https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#issues-events).
- `/bang`
    - `GET` The URL that the client visits when it wants to know if the bell should be rang. The client gets JSON with 
    property `needs_banging` that is `true` if the bell should be banged. This path also accepts a `GET` parameter 
    `will_bang`. If this parameter is `true` then the server will set `needs_banging` to `false`, otherwise the server
    won't change the state.

## Building
You can run `go build` to build the project. 

For production building run `build.sh` which will prune the dependencies, build the project 
and move the executable and all the files needed to run the server
(like config file) to `build/` directory. 

## Config
Config of the backend is specified inside `config.json` file. JSON should be structured as
```json
{
  "conf_type": "name of config type",
  "name of config type 1": {
    "parameter": "value"  
  },
  "name of config type 2": {
    "parameter": "value"
  }
}
```

Parameters are:
- `min_log_level (int)` minimum level to log,
- `listen_address (string)` address (with port) on which the server should run. 
