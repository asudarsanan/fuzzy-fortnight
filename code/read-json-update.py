import json
import yaml



def update_json(d):
    for i in d:
        print (i['Name'])
        if i['Name'] == "sudarnet-outline":
            print(i['Status'])
            if i['Status'] != '1':
                print(i)
                i['Status'] = "Not Working"
                print(i)
    with open('updated_stacks.json',"x") as f:
        json.dump(d,f,ensure_ascii=False,indent=4,)
                

def create_yaml_file(d):
    print(yaml.dump(d))
    with open ('stacks.yaml',"x") as yaml_file:
        yaml.dump(d,yaml_file,default_flow_style=False)

with open('stacks.json') as json_data:
    d =json.load(json_data)
    json_data.close()
    print(d)

# update_json(d)
create_yaml_file(d)