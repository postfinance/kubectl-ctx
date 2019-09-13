[![Release](https://img.shields.io/github/release/postfinance/kubectl-vault_sync.svg?style=for-the-badge)](https://github.com/postfinance/kubectl-ctx/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Go Report Card](https://img.shields.io/badge/GOREPORT-A%2B-brightgreen.svg?style=for-the-badge)](https://goreportcard.com/badge/github.com/postfinance/kubectl-ctx)
# kubectl ctx plugin
Simple Plugin to display/change the current kube context in your KUBECONFIG.

# Installation
Pre-compiled statically linked binaries are available on the [releases page](https://github.com/postfinance/kubectl-ctx/releases).
Binary must be placed anywhere in `$PATH` named `kubectl-ctx` with execute permissions.  
For further information, see the offical documentation on plugins [here](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/).

# Compatibility
Known to work on Windows and Linux. Requires kubectl >= 1.12 (tested with versions >1.12).

# Examples
For all the examples, assume you have the following contexts.
```
foo
bar
baz
localhost
```
## display contexts
Current context is displayed in a different color.
```bash
$ kubectl ctx
bar
baz
foo
localhost
```

Substring matching can be used to display contexts. For example if you are searching a context named `ba` simply type:
```bash
$ kubectl ctx ba
bar
baz
```

## change current context
You can switch the context by providing an exact name:
```bash
$ kubectl ctx foo
current context set to "foo"
```

But it's also possible to switch to the `localhost` context by typing a substring (as long as it is a unique name), for example:
```bash
$ kubectl ctx local
current context set to "localhost"
```
