# What is codegen?

It is a database first code generator for an ORM Layer.

# How does it work?

It uses 2 steps.

1. Inspect MySQL and create or update a JSON configuration file, that represents the database.
   - This file may be changed to use some advanced features, which are not fully automagically possible with MySQL inspection.
2. Read the config file and generate the code.

# Usage for project know-api

platform and tenant settings have been merged into one config file.
tool has been refactored to use relative paths.

## Install codegen

```
cd cmd/codegen
go install
```

## Update from Database

```
# goto know-api folder
codegen update
```

## Generate code

```
# goto know-api folder
codegen gen
```
