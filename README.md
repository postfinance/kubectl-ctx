# kubectl ctx plugin
Simple Plugin to display/change the current kube context in your KUBECONFIG.

# Build/Installation
## Build from source
go version >= 1.11 with modules support enabled is required to build the plugin from source
```bash
export GO111MODULES=on
go build
```

## Installation
Binary must be placed anywhere in `$PATH` named `kubectl-ctx` with execute permissions.  
For further information, see the offical documentation on plugins [here](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/).

# Compatibility
Known to work on Windows and Linux. Requires kubectl >= 1.12 (tested with 1.12).

# Examples
## display contexts
Current context is displayed in a different color.
```bash
$ kubectl ctx
foo
bar
baz
```

## change current context
```bash
$ kubectl ctx foo
current context set to "foo"
```
