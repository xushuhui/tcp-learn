import json

class UseJson():
    def __init__(self):
        pass
    def read_data(self):
        with open('./static/data.json') as fp:
            data = json.load(fp)
            return data
    def get_data(self,key):
        self.data = self.read_data()
        return self.data[key]
    def get_json_data(self,key):
        return json.dumps(self.get_data(key))
    def write_data(self,key,value=None):
        data = self.read_data()
        with open('./static/data.json',"w+") as f:
            data[key] = value
            return json.dump(data,f)
            
if __name__ == '__main__':
    client = UseJson()
    print(client.get_data('createRoom','111'))