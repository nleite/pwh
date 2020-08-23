# pwh

Password hiding tool

## Usage

*Ever wanted to stop leaking passwords while sharing your screen?*

I have, many times.

```
cat data.txt | pwh 
command_that_leasks_passwords 2>&1 | pwh
```

Done.


## Installation

```
go get -v github.com/nleite/pwh
```
