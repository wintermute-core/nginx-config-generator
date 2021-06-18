# nginx config generator

Generator of nginx configuration based on declarative definition of services.

Example input and output can be found in `examples` directory.s

Tech stack:
 * go 1.16
 * make
 * docker

Flow:
 * parse input file
 * build internal model
 * render nginx config based on internal model

Implementation details/assumptions:
 * go templates are used to generate nginx configs
 * yaml.v2 is used to parse yaml files
 * In one invocation of application can be processed one file
 * All inputs are passed through environment variables
 * For not passed environment variables, will be used defaults
 * Stdout is used only for logging of application actions
 * make and makefile is used to build, test, package application
 * Output file will be replaced during each execution

# File structure

Main files and directories

| File             | Description                                                              |
|------------------|--------------------------------------------------------------------------|
| model.go         | Define go structures for input and output                                |
| input.go         | Code used to read input file                                             |
| input_test.go    | Tests used to verify input loading                                       |
| renderer.go      | Functions used to render nginx configurations                            |
| renderer_test.go | Tests to validate renderer                                               |
| main.go          | Main application logic to load input, transform in nginx model and render|
| main_test.go     | Tests used to validate main application logic                            | 
| examples         | Example input and output files used to build this tool                   |
| itest            | Directory with files for integration tests                               |

# Usual tasks

Local build:

```
make build
```

Tests execution:
```
make test
```

Note:
 * code coverage report is automatically generated

Integration tests execution:
```
make itest
```
Notes:
 * in integration tests will be used files from `itest` directory
 * integration test validate generated configuration through nginx instance running in docker

Container image building:
```
make container
```
Notes:
 * git hash will be used in image tag

Explicit binary execution:
```
INPUT_FILE=examples/input.yml ./nginx-config-generator  
```

# Future work

 * extended validation: IP ranges validation, ports
 * validate if application paths override
 * add more flexible configuration
 * CI/CD integration for automated building
 * extend test cases to check all generated fields
 * process multiple files in one execution
 * extract more parameters to be configurable through environment variables

# License

Only for reference, distribution and/or commercial usage not allowed
