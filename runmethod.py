#coding:utf-8
import requests
import json
class RunMethod:
    def post(self,url=None,data=None,header=None):
        res = requests.post(url=url,data=data,headers=header,verify=False)
        return res.json()
    def get(self,url=None,data=None,header=None):
        res = requests.get(url=url,data=data,headers=header,verify=False)
        return res.json()
    def put(self,url=None,data=None,header=None):
        res = requests.put(url=url,data=data,headers=header,verify=False)
        return res.json()
    def delete(self,url=None,data=None,header=None):
        res = requests.delete(url=url,data=data,headers=header,verify=False)
        return res.json()
    def main(self,method):
        method = getattr(self, method)
        return method
		
if __name__ == '__main__':
    client = RunMethod()
    client.main("put")("http://www.baidu.com")