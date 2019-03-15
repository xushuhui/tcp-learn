import json

class UseJson():
    def __init__(self):
        self.data = self.read_data()
    def read_data(self):
        with open('./static/data.json') as fp:
            data = json.load(fp)
            return data
    def get_data(self,key):
        return self.data[key]

if __name__ == '__main__':
    client = UseJson()
    print(client.get_data('login'))