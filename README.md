# ocp-gen

`ocp-gen` is a code transformation tool to use together with `go generate`.

## Commands

`ocp-gen` defines a `Command` interface to transform go source files line by line.

`Commands` are executed by a `Runner` which processes go source files in single pass fashion.

This requires `Commands` to be stateful in regards of their activation state and arguments.

## Runner

A `Runner` takes an abitrary number of `Commands` to transform individual go source files.

When nesting `Commands` in source files, be aware of the order that `Commands` are defined in the `Runner`.

## Usage

`ocp-gen` is expected to be used with `go generate`, see the following example source code file:

```go
//go:generate ocp-gen
package test

// ocp-gen:replace TestReconciler=RECONCILER_NAME
type TestReconciler struct {
	namespace string
	//ocp-gen:if CONDITIONAL=include
	conditionalName string
	// test nested command
	//ocp-gen:replace Age=GOLINE
	conditionalAge int
	//ocp-gen:fi
}

// ocp-gen:replace TestReconciler=RECONCILER_NAME
func (r *TestReconciler) Reconciler() {
	//ocp-gen:replace name=RECONCILER_NAME age=GOLINE
	var _age, name string
	//ocp-gen:replace name=RECONCILER_NAME
	_ = name
	//ocp-gen:replace age=GOLINE
	_ = _age
}
```

`ocp-gen` allows in memory transformation via the `DRY_RUN` environment variable:

```shell
DRY_RUN=true go generate ./...
```

Pass values by defining the environment variables used in the command directives:

```shell
DRY_RUN=true RECONCILER_NAME=MyReconciler WITH_FIELDS=true go generate ./...
```

Save changes to disk:

```shell
go generate ./...
```

## Execution Options

Environment Variables:

- DRY_RUN (bool): 1 or true result in the in-memory source code transformation being printed to stdout.
- DEBUG (bool): 1 or true result in debug output being printed to stdout.
