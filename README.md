# kubectl ctx plugin
Simple Plugin to display/change the current kube context in your KUBECONFIG.

# Requirements
* kubectl >= 1.12
* go >= 1.11 (with GO111MODULES=on)

# Build/Installation
## Buildfrom source
```bash
go build
```

## Installation
Must be placed anywhere in `$PATH` named `kubectl-ctx`.

# Examples
## display contexts
Current context is printed in a different color.
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
