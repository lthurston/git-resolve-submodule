# git resolve-submodule

A stupid simple git plugin written in Go that helps navigate around submodules
like filesystem directories.

## Usage

$ git resolve-submodule &lt;pathish&gt;

The pathish value can be ".", ".." or a string that identifies the submodule.

## Installation

1. Install Go and make sure that go build path is in $PATH
2. Add a function to your dot pile that changes directory to the returned value:

```
function cm {
  cd `git resolve-submodule $1`
}
```

## How to use it

### Jump into a submodule

```
$ cm build
```

Changes directory to the root of build submodule root.

```
$ cm ..
```

Changes directory to the root parent repository.

```
$ cm .
```

Changes directory to root of the current repository.

## How do I identify a submodule with _pathish_?

Use any string that appears in `git submodule`, including the commit id, the path, or the tag. The first
submodule in which the provided "pathish" is found is the submodule which it will cd to.

## Works on:

OSX 10.9.3 / Go 1.3 / Git 1.9.1

I haven't tried anything else.
