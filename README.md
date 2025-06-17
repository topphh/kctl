Usage:

install:
go build && go install

to get all available commands:
kctl --help

sample command:
kctl top service --sort-by memory --sort-by memory 
sortened version:
kctl top service --s memory -H
