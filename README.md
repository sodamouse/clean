# Clean
Clean is a simple command-line utility that clean's the user's home directory by removing file/directory names specified in a simple list from $HOME.

## Current implementations
Implementations are current available in the following languages:
- Go

## Usage
All implementations use the same configuration file located at "$XDG_CONFIG_DIR/clean/files.list". Populate this file with a **line-break** ('\n') separated list of files to be removed. For example using the follwing list...

*~/config/clean/files.list*
``` text
foo
.bar
-baz
```

... result in the following files or directories (recursive) being deleted from "/home/user/": ~/foo, ~/.bar, ~/-baz
