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
### Get and set session credentials to environment variable
```bash
caws session --profile <profile for mfa> \
             --serial-number <arn of the mfa device>
```

this command start shell with environment variable for AWS credentials.

## Exit shell
```bash
exit
```