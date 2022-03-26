<h1 align="center">allpasswords</h1>
<p align="center">一个社工密码生成器<br/>喜欢请点赞🌟⭐，不迷路</p>

<p align="center"><img src="https://github.com/zan8in/allpasswords/blob/main/screen.png"/></p>

### 用法

```shell
allpassword.exe
```

按照提示输入项，每一项可填写多个，什么都不填回车跳到下一项输入，示例如下：

```shell
PS C:\Users\zanbi\Desktop> allpassword.exe
FullName[姓名全拼](例如：zhang san)：cheng long
FullName[姓名全拼]：  
ShortName[姓名简拼](例如：zs)：cl
ShortName[姓名简拼]：
BrithDay[生日](例如：19961023)：19901201
BrithDay[生日]：
Phone[个人手机号]：13766666666
Phone[个人手机号]：
CompanyFullName[公司全拼](例如：beijingdaxue)：chengjiaban
CompanyFullName[公司全拼]：
CompanyShortName[公司简拼](例如：bjdx)：cjb
CompanyShortName[公司简拼]：
CompanyTelephone[公司电话]：22020202
CompanyTelephone[公司电话]：
JobNumber[员工号码](例如：9521)：9521
JobNumber[员工号码]：
生成密码文件: password20220326225658.txt
生成账号文件: account20220326225658.txt
```

### [下载地址](https://github.com/zan8in/allpasswords/releases/tag/v0.1.0)

### 概述

工作中发现弱口令 屡试不爽，YYDS，但遇到没有弱口令情况怎么办？不要放弃，再博一下，试试定制社工密码，也许会柳暗花明又一村。

### 人们设置密码习惯

```
姓名全拼 liudehua
姓名简拼 ldh
手机号 13732611616

公司名称全拼 baidu
公司名称简拼 bd
公司电话 2208661
员工工号 9521
```

### 弱密码类型

```
纯数字 123456789 123456 1234567890 123 111111 666666 888888
纯字母 admin test ceshi password p@ssword
字母加数字 a123456 a123456789 woaini1314 abc123456 qq123456 123456789a 123456a admin123 1qaz2wsx aaa111 a1w2e3
键盘 qwertyuiop asdfghjkl zxcvbnm 147258369 147258 741258 123qwe qweasdzxc
中国特色 5201314 iloveyou woaini fuckyou wohenni
手机号密码 13732611616 a13732611616 ldh13732611616
生日 20220326 ldh20220326 l20220326
特殊字符 ! @ # $ % [] / . + ^
```

### 如何定制规则

```
姓名全拼 √
姓名全拼+纯数字 liudehua123 √
姓名全拼+特殊符号+数字 liudehua@123 √
姓名全拼+数字+特殊符号 liudehua123@@ √
姓名全拼+生日 ldh2022 ldh20220326 √
姓名全拼+特殊符号+生日 liudehua@20220326 liudehua@2022 √
姓名全拼+生日+特殊符号 liudehua20220326@@ liudehua2022@@ √
姓名全拼+工号 ldh9527 √
姓名全拼+特殊符号+工号 liudehua@9527 √
姓名全拼+工号+特殊符号 liudehua9527@@ √
姓名全拼+手机号 ldh13732611616 √
姓名全拼+特殊符号+手机号 liudehua@13732611616 √
姓名全拼+手机号+特殊符号 liudehua13732611616@@ √

姓名简拼 规则 同上 √

手机号 √
手机号+姓名全拼 13732611616liudehua √
手机号+特殊字符+姓名全拼 13732611616@liudehua √
手机号+姓名简拼 13732611616ldh √
手机号+特殊字符+姓名简拼 13732611616@ldh √
手机号+特殊字符 13732611616@ √

公司全拼 √
公司名称全拼+公司电话 baidu2208661 √
公司名称全拼+员工工号 baidu9521 √
公司名称全拼+特殊符号+公司电话 baidu@2208661 √
公司名称全拼+特殊符号+员工工号 baidu@9521 √
公司名称全拼+公司电话+特殊符号 baidu2208661!! √
公司名称全拼+员工工号+特殊符号 baidu9521@ √
公司名称全拼+年份 baidu2022 √

公司简拼 规则 同上 √

员工工号 0-9 00-10 0000-0010 0100-0110 1000-1010 2000-2010 3000-3010 00000-00010 00100-00110 10000-10010 20000-20010 30000-30010 000000-000010 000100-000110 100000-100010 200000-200010 300000-300010 11 111 1111 11111 111111 66 666 6666 66666 666666 88 888 8888 88888 888888 99 999 9999 99999 999999 √
```

