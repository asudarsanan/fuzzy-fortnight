import requests
import json




portainer_host = 'busybox.sudarnet.in'
base_path = '/api'

endpoint ="https://"+portainer_host+base_path

def get_all_stacks(username,password):

    auth_data = {}
    auth_data["username"] = username
    auth_data["password"] = password
    json_data = json.dumps(auth_data)
    response = requests.post(endpoint+'/auth',data=json_data)
    response_json = response.json()
    jwt = response_json["jwt"]
    response_stacks = actual_get_stacks(jwt)
    wite_json_to_file(response_stacks)

    print(response,response_json,username,password,json_data,endpoint)
    return response_json


def actual_get_stacks(jwt):
    header = {"Authorization":"Bearer "+jwt}
    response = requests.get(endpoint+'/stacks',headers=header)
    response_json = response.json()
    print(response,header,response_json)
    print(response_json)

    return response_json


def wite_json_to_file(response_json):
    f = open('stacks.json',"x")
    json.dump(response_json,f,ensure_ascii=False, indent=4)
    f.close()

    

access_token='ptr_fvjxvSBI2/0CkMi48Un+3MnffQDQ/489GPvkN5JqDxo='
username='asudarsanan-admin'
password='@Autowired123'
get_all_stacks(username,password)

