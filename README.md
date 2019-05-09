# What is codegen?

It is a database first code generator for an ORM Layer.

# How does it work?

It uses 2 steps.

1. Inspect MySQL and create or update a JSON configuration file, that represents the database.
   - This file may be changed to use some advanced features, which are not fully automagically possible with MySQL inspection.
2. Read the config file and generate the code.

# Example for project know-api

platform

```
cd cmd/codegen
go build
 ./codegen -config ~/gospace/know-api/codegenplatform.json update

 # generation will not work at the moment (see Attention below)
 # ./codegen -config ~/gospace/know-api/codegenplatform.json gen
```

instance / tenant

```
 ./codegen -config ~/gospace/know-api/codegen.json update
```

## Attention

The two config files are not up to date. Maybe even some codegen changes have to be made. Currently paths are absolute and refer to an old location on my machine.
We need to switch to relative paths and update the files respectively.
