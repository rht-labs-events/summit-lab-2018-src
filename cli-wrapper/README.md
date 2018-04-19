# CLI Labs Wrapper


# Purpose:

To abstract lab init from user and enable ansync scenarios execution.
Each scenario is defined in `config.yml` file and compiled binar =y in server bastion path allows users to call any scenario at any time.

This enables users to start from scenario 3, and skip 1 and 2 if they want.

In addition we keep scripts for scenarios somewhere in the unknown location on the server so "start asses" will not be able to investigate scripts easily :)


# Run:

Create config file:
config.yaml

```
scripts:
  - name: "scenario 1"
    description: "etcd magic magic"
    executor: "/usr/local/bin/test.sh"
    actions: "[init, solve, hint]"
  - name: "scenario 2"
    desicription: "adasd"
    executor: "/usr/local/bin/asd.sh"
    actions: "[init, sovle, hint]"
```

Now running `clu-wrapper` you can get those scenarios:

```
[mjudeiki@redhat cli-wrapper]$ ./cli-wrapper -h
INFO[0000] Starting Wrapper                             
NAME:
   RH Summit Labs - Init Labs Scenarios

USAGE:
   cli-wrapper [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   Dream Team

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --scenario value, -S value, -s value  The Name of the scenario
   --action value, -A value, -a value    action (default: "init")
   --list, -L, -l                        List all scenarios
   --help, -h                            show help
   --version, -v                         print the version


[mjudeiki@redhat cli-wrapper]$ ./cli-wrapper -l
INFO[0000] Starting Wrapper                             
---------------------------------------------------------------------
Scenario 0
Description: etcd magic magic
Actions: [init, solve, hint]
To init this scenario execute:
  cli -s 1 -a init
To get hints for this scenario execute:
  cli -s 1 -a hint

---------------------------------------------------------------------
Scenario 1
Description: 
Actions: [init, sovle, hint]
To init this scenario execute:
  cli -s 1 -a init
To get hints for this scenario execute:
  cli -s 1 -a hint

[mjudeiki@redhat cli-wrapper]$ ./cli-wrapper -s 1 -a init
INFO[0000] Starting Wrapper                             
INFO[0000] Execute scenario 1 and phase init            

INFO[0000] At this place script /usr/local/bin/asd.sh will be executed with args init and output printed to user 
```

# Compile:

```
go build
```