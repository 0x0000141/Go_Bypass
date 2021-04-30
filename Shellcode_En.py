# -*- coding: UTF-8 -*-
import requests
import base64
import sys
import json

shellcode = b"shellcode"
ShellTo_base64 = base64.b64encode(shellcode).decode()
password = sys.argv[1]
url = "http://tool.chacuo.net/cryptaes"

header = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4469.4 Safari/537.36",
    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
}

payload = {
    "data": ShellTo_base64,
    "type": "aes",
    "arg": "m=ecb_pad=zero_block=128_p="+password+"_i=10_o=0_s=gb2312_t=0"

}
res = requests.post(url, data=payload, headers=header)
shellcode_base64 = res.text
json = json.loads(shellcode_base64)
Exp = json['data']
print(Exp)
with open('Dir.txt', 'a', encoding='utf-8') as f:
    f.write(str(Exp))
