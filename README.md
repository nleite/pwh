h1. pwh

Password hiding tool

h3. Ever wanted to stop leaking passwords while sharing your scree?

I have, many times.

```
cat some_file_with_passwords  2>&1 |  pwh 
command_that_leasks_passwords 2>&1 | pwh
```

Done.