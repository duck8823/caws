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
#### Start shell with environment variable for AWS credentials.
```bash
caws mfa --profile <profile for mfa> \
             --serial-number <arn of the mfa device>
```

This command set environment variables bellow

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`

#### Exit shell
```bash
exit
```