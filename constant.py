#coding:utf-8
class global_var:
	#case_id
	Id = '1'
	request_name = '2'
	url = '3'
	run = '4'
	request_way = '5'
	header = '6'
	case_depend = '7'
	data_depend = '8'
	field_depend = '9'
	data = '10'
	expect = '11'
	result = '12'
#获取caseid
def get_id():
	return global_var.Id

#获取url
def get_url():
	return global_var.url
#获取url
def get_request_name():
	return global_var.request_name

def get_run():
	return global_var.run

def get_run_way():
	return global_var.request_way

def get_header():
	return global_var.header

def get_case_depend():
	return global_var.case_depend

def get_data_depend():
	return global_var.data_depend

def get_field_depend():
	return global_var.field_depend

def get_data():
	return global_var.data

def get_expect():
	return global_var.expect

def get_result():
	return global_var.result

def get_header_value():
	return global_var.header
def get_domain():
	return "https://dev.yunsuit.com"