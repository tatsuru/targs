# targs

construct argument list(s) and execute in tmux window like `xargs`

## SYNOPSIS

```
find /var/log -type f | targs -r ls -l
echo 'host1\nhost2' | targs mysql -uroot -h
```
