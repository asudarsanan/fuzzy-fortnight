import re

def extract_incoming_ips(log_file):
    log_data = []

    with open(log_file,'r') as lf:
        for line in lf:
            log_match = re.search(r'(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - - \[.*?\] "(.*?) (.*?)" (\d{3})', line)
            print (log_match)
            if log_match:
                ip = log_match.group(1)
                request_method = log_match.group(2).split()[0]  # Extracting the request method (e.g., GET, POST)
                path = log_match.group(3)
                response_code = log_match.group(4)
                log_data.append({'IP': ip, 'Request Method': request_method, 'Path':path, 'Response Code': response_code})
                print(log_data)




def main():
    file_name='resources/sample.log'
    extract_incoming_ips(file_name)

if __name__ == "__main__":
    main()

# log_match = re.search(r'(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) - - ')