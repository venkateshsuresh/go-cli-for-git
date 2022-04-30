# go-cli-for-git

This Project helps to get the Repository Information of the gitHub Account.

This Project uses [Cobra](https://github.com/spf13/cobra) and [gitHub Rest API](https://docs.github.com/en/rest)

```shell
go build -o <binary-name>
```

```shell
./<binary-name> get -u <gitHub username> -p <gitHub_Access_Token>
```

or 

```shell
./<binary-name> get --username <gitHub username> --password <gitHub_Access_Token>
```
