# nginx config generator

Generator of nginx configuration based on declarative definition of services.

Example input and output can be found in `examples` directory.s

Tech stack:
 * go 1.16
 * make

Flow:
 * parse input file
 * build internal model
 * render nginx config based on internal model



# Building

```
go build
```

Local execution:
```
INPUT_FILE=examples/input.yml ./nginx-config-generator  
```

# Future work

 * extended validation: IP ranges validation, ports
 * validate if application paths override
 * add more flexible configuration

# License

Only for reference, distribution and/or commercial usage not allowed
