#coding:utf-8
import requests
import json
class RunMethod:
    def post(self,url=None,data=None,header=None):
        print(url)
    def get(self,url=None,data=None,header=None):
        print("get")
    def put(self,url=None,data=None,header=None):
        print("put")
    def delete(self,url=None,data=None,header=None):
        print("delete")
    def main(self,method):
        method = getattr(self, method)
        return method
		
if __name__ == '__main__':
    client = RunMethod()
    client.main("put")("http://www.baidu.com")