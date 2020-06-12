# drone-plugin-starter

Starter project for creating Drone plugins. Plugins are packaged in docker images and are applications built to read 
environment variables passed to them by a Drone executor. This project serves as a skeleton to writing a drone plugin in
go using (urfave/cli)[https://github.com/urfave/cli]] and configuring the flags to accept the DRONE_ and PLUGIN_ 
prefixed environment parameters. A full list of the DRONE_ params can be found in the (documentation)
[https://docker-runner.docs.drone.io/configuration/environment/variables/] and the PLUGIN_ params are passed in via 
the `settings:` YAML field on your step calling the plugin.

### Metadata via Environment variables

Build and Repository metadata are prefixed with `DRONE_` and sent to the plugin at runtime. The full list of available 
parameters are already included in the `flags.go` file as command line flags. You should remove un-used parameters from 
the list so that one can easily see which parameters are used by which plugins.

Example parameters:

```
cli.IntFlag{
    Name:   "build.number",
    Usage:  "build number",
    EnvVars: []string{"DRONE_BUILD_NUMBER"},
},
cli.StringFlag{
    Name:   "build.status",
    Usage:  "build status",
    Value:  "success",
    EnvVars: []string{"DRONE_BUILD_STATUS"},
},
```

### Parameters

Plugin settings are defined in the yaml file:

```
settings:
  username: drone
```

They are prefixed with `PLUGIN_` and sent to the plugin at runtime:

```
PLUGIN_CHANNEL=dev
PLUGIN_USERNAME=drone
```

Plugin settings can be retrieved using `cli.Flag` as seen below:

```
cli.StringFlag{
    Usage:  "slack channel",
    EnvVars: []{"PLUGIN_CHANNEL"},
},
cli.StringFlag{
    Usage:  []string{"slack username"},
    EnvVars: []string{"PLUGIN_USERNAME"},
},
```

### Secrets

Sensitive fields should not be specified in the yaml file. Instead they are passed to your plugin as environment variable. Secrets should use a prefix that corresponds to the plugin name. For example, the Slack plugin prefixes secrets with `SLACK_`:

```
cli.StringFlag{
    Usage:  "slack api token",
    EnvVars: []string{"SLACK_TOKEN"},
},
```

### Logos

Please replace the logo.svg file with a meaningful svg logo for your plugin. If you are you building a Slack plugin, for example, please provide the Slack logo. This icon is displayed when your plugin is listed in the official plugin index.

### Docs

Please provide a DOCS.md file in the root of your repository that documents plugin usage. This documentation is displayed when your plugin is listed in the official plugin index. Use the README.md file to describe building and testing the plugin locally.

### Images

Images are distributed as Docker images in the public Docker registry. Please use a minimalist alpine image when possible to limit the image download size. We are also working on supporting multiple plugin architectures, with compatibility guidelines coming soon

### Testing

Please create plugins that are easily runnable from the command line. This makes it much easier to debug and test plugins locally without having to launch actual builds.

### Vendoring

Please vendor dependencies in a manner compatible with go modules. All official drone plugins should use [go mod vendor](https://blog.golang.org/using-go-modules).
