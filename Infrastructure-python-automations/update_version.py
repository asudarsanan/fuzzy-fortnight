import yaml
import semver


def read_yaml(file_name):
    with open(file_name,'r') as yf:
        yaml_data = yaml.safe_load(yf)
        # print(yaml_data)
        return yaml_data
    
def get_dependencies(yaml_data):
    dependencies = yaml_data.get('dependencies',[])
    for dependencie in dependencies:
        print (dependencie)
    return dependencies


def update_dependency_version(dependencies,name,version):
    for dependency in dependencies:
        if dependency['name'] == name:
            print("Ongoing version: "+ dependency['version'])
            print("Incoming version: "+ version)
            outcome = semver.compare(dependency['version'],version)
            if outcome == -1:
                print("Accepting Incoming version post semver comparison:= Incoming  {} >> Ongiong {}".format(version,dependency['version']))
                dependency['version'] = version
                print(dependency)
            elif outcome == 1:
                print("Rejecting Incoming version post semver comparison:= Ongoing {} >> Incoming {}".format(dependency['version'],version) )
            else:
                print("No change in dependency versions:= Incoming {} == Ongoing {}".format(version, dependency['version']))
    return dependencies


                
def rewrite_dependencies_into_yaml(yaml_data,file_name,dependencies):
    yaml_data['dependencies'] = dependencies
    with open (file_name, 'w') as yf:
        yaml.dump(yaml_data,yf,default_flow_style=False,sort_keys=False)
        yf.close()


# def semver_outcome(incoming,ongoing):



def main():
    file_name = 'charts.yaml'
    print("Reading {} file".format(file_name))
    yaml_data = read_yaml(file_name)
    print("Checking for dependencies..")
    dependencies = get_dependencies(yaml_data)
    print("Updating version for redis..")
    dependencies = update_dependency_version(dependencies,'redis','17.9.5')
    print("Rewrite yaml file {}".format(file_name))
    rewrite_dependencies_into_yaml(yaml_data,file_name,dependencies)


if __name__ == "__main__":
    main()