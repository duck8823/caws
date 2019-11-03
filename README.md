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
### Get and set session credentials
```bash
caws session --profile <profile for mfa> \
             --serial-number <arn of the mfa device> \
             --session-profile <profile for session>
```