
from get_data import GetData
from runmethod import RunMethod
from utils.use_json import UseJson
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
class main():
    def __init__(self):
        self.run_method = RunMethod()
        self.data = GetData()
        self.json = UseJson()
    def run(self):
        res = header =method = None
        pass_count = []
        fail_count = []
        rows_count = self.data.get_case_lines()
        
        for i in range(1,rows_count):
            is_run = self.data.get_is_run(i)
            if is_run:
                url = self.data.get_request_url(i)
                method = self.data.get_request_method(i)
                request_data = self.data.get_data_for_json(i)
                # expect = self.data.get_expcet_data_for_mysql(i)
                header = self.data.is_header(i)

                request_name =  self.data.get_request_name(i)
                
                # depend_case = self.data.is_depend(i)
            # if depend_case != None:
            #     self.depend_data = DependdentData(depend_case)
            #     #获取的依赖响应数据
            #     depend_response_data = self.depend_data.get_data_for_key(i)
            #     #获取依赖的key
            #     depend_key = self.data.get_depend_field(i)
                #request_data[depend_key] = depend_response_data
                if header == 'yes':
                    
                    token = self.json.get_data("token")
                    
                    head = {
                        "token": token
                    }
                    res = self.run_method.main(method)(url,request_data,head)
                    
                else:
                    res = self.run_method.main(method)(url,request_data)
                    
                if res['error_code'] == 0:
                    print("测试："+request_name+",","结果：success")
                    #存token到文件
                    # if res['data']['token']:
                    #     self.json.write_data("token",res['data']['token'])
                else:
                    print("测试："+request_name+",","结果：fail"+",","原因："+res["msg"])
                   
                

            # if self.com_util.is_equal_dict(expect,res) == 0:
            #     self.data.write_result(i,'pass')
            #     pass_count.append(i)
            # else:
            #     self.data.write_result(i,res)
            #     fail_count.append(i)

if __name__ == '__main__':
    client = main()
    client.run()