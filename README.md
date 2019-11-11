# caws
A command line tool which manage aws credentials

```
　　　　　 ,,.　､､
　　　_,ノ´.: : : : 丶-‐ｧ…　　ｰ- ､ ＿＿＿＿,　 
　　 ＼ __,ノ: : : : : : /　　　　　／: : : : 
　　　/ ':: : : : : : : :〈　　　　　/ : : : :
　　ノ ﾉ: :ヽ: : : : : : :}　　　　 {: : : : 
　〔_　 ＼:_:_>､: :l: : :ヽ、　　　1: : : : : 
　　 ｀￣　　　 ｀｛: : : : ＼　、 }: : : : : :
　　　 　 　 　 　 ヽ: : : :！ ヽ ￣{ : : : 
```

## Installation
```bash
go get github.com/duck8823/caws
```

## Usage
#### Show profile names
```bash
caws ls [--file <path to credentials file>]
```

#### Use specific profile
```bash
caws use --profile <profile to use>
```

Set specific credentials to environment variables bellow

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`

and start new shell.

#### Login and set credentials with MFA
```bash
caws mfa --profile <profile for mfa> \
         --serial-number <arn of the mfa device> \
         [--output <output profile name>] \
         [--file <file set new profile>]
```

If you set a `--output` flag, this command set new profile to credentials file.
Without `--output`, set environment variables bellow

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`

and new shell.

#### Exit shell
```bash
exit
```