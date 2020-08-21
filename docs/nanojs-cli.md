# Nanojs CLI Tool

Nanojs is designed as an embedding script language for Go, but, it can also be
compiled and executed as native binary using `nanojs` CLI tool.

## Installing Nanojs CLI

To install `nanojs` tool, run:

```bash
go get github.com/zeaphoo/nanojs/cmd/nanojs
```

Or, you can download the precompiled binaries from
[here](https://github.com/zeaphoo/nanojs/releases/latest).

## Compiling and Executing Nanojs Code

You can directly execute the Nanojs source code by running `nanojs` tool with
your Nanojs source file (`*.js`).

```bash
nanojs myapp.js
```

Or, you can compile the code into a binary file and execute it later.

```bash
nanojs -o myapp myapp.js   # compile 'myapp.js' into binary file 'myapp'
nanojs myapp                  # execute the compiled binary `myapp`
```

Or, you can make nanojs source file executable

```bash
# copy nanojs executable to a dir where PATH environment variable includes
cp nanojs /usr/local/bin/

# add shebang line to source file
cat > myapp.js << EOF
#!/usr/local/bin/nanojs
fmt := import("fmt")
fmt.println("Hello World!")
EOF

# make myapp.js file executable
chmod +x myapp.js

# run your script
./myapp.js
```

**Note: Your source file must have `.js` extension.**

## Resolving Relative Import Paths

If there are nanojs source module files which are imported with relative import
paths, CLI has `-resolve` flag. Flag enables to import a module relative to
importing file. This behavior will be default at version 3.

## Nanojs REPL

You can run Nanojs [REPL](https://en.wikipedia.org/wiki/Read–eval–print_loop)
if you run `nanojs` with no arguments.

```bash
nanojs
```
