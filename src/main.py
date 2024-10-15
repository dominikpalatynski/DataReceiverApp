from influxdb_client import InfluxDBClient, Point, WritePrecision
from datetime import datetime

url = "http://localhost:8086"  
token = "mytoken"               
org = "myorg"                   
bucket = "mybucket"             

client = InfluxDBClient(url=url, token=token, org=org)

with client.write_api() as write_api:
    point = Point("measurement_name") \
        .tag("companyId", "company_1") \
        .tag("machineId", "machine_1") \
        .field("temperature", 25.0) \
        .time(datetime.utcnow(), WritePrecision.NS)
    
    write_api.write(bucket=bucket, record=point)


query = f'from(bucket: "{bucket}") |> range(start: -1h) |> filter(fn: (r) => r._measurement == "measurement_name")'
tables = client.query_api().query(query=query, org=org)

for table in tables:
    for record in table.records:
        print(f'Time: {record.get_time()}, Value: {record.get_value()}')
