# import os
#
# import yaml
# import json
# # from main import Environ
#
#
# class ReadData:
#     def __init__(self, filename, data_dir='/'):
#         self.filepath = os.path.abspath(os.path.dirname(os.path.dirname(__file__)) + f"/data/{data_dir}")+"/"+filename
#
#     def read_yaml_data(self, input_path=None):
#         filepath = input_path if input_path else self.filepath
#         with open(filepath, "r", encoding="utf-8")as f:
#             yaml_data = yaml.load(f, Loader=yaml.FullLoader)
#         f.close()
#
#         return yaml_data
#
#     def read_json_data(self):
#         with open(self.filepath, "r", encoding="utf-8")as f:
#             json_data = json.loads(f.read())
#         f.close()
#
#         return json_data
#
#
# if __name__ == '__main__':
#     data = ReadData("test_yml_data.yml").read_yaml_data()
#     data1 = ReadData("test_yml_data.json").read_json_data()
#     print(data)
#     print(data1)
