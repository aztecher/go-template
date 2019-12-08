# README

A template that is useful for implementing a command line tool that takes only one subcommand or a subcommand that expresses resources and execution contents with a single word connected by dots (like  host.ls, license.assigned.ls)

# Example Command

```
# exc [subcmd]

$ exc host.ls
$ exc port.ls

# output option
$ -json/--json -- json format
$ -xml/--xml   -- xml format

```


