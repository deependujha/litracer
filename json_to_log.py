import json

file_name = "meow.log"

# read the json file
with open('multiprocess_trace.json', 'r') as f:
    data = json.load(f)

# print the data
trace_events = data['traceEvents']

for event in trace_events:
    for key, value in event.items():
        with open(file_name, 'a') as f:
            f.write(f"{key}: {value}; ")
    with open(file_name, 'a') as f:
        f.write("\n")

