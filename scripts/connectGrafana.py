import requests
import json

# Adres URL do Grafany i API
grafana_url = "http://localhost:3000/api/datasources"
api_key = "glsa_YWOsdfHYsYmt8JucU3jhYJcAwmeWQR8h_62c0a583"

# Nagłówki uwierzytelniające i Content-Type
headers = {
    "Authorization": f"Bearer {api_key}",
    "Content-Type": "application/json"
}

# Konfiguracja źródła danych dla InfluxDB 2.0 (Flux)
data = {
    "name": "TestDB",
    "type": "influxdb",
    "access": "proxy",
    "url": "http://host.docker.internal:8086",  # Adres do InfluxDB
    "user": "",  # Nie wymagane przy używaniu tokenów
    "password": "",  # Nie wymagane przy używaniu tokenów
    "database": "",  # Puste, bo w Flux nie definiujemy bazy danych na tym etapie
    "jsonData": {
        "httpMode": "GET",  
        "version": "Flux",
        "organization": "myorg",  # Twoja organizacja w InfluxDB
        "defaultBucket": "bucket"  # Domyślny bucket (baza danych)
    },
    "secureJsonData": {
        "token": "mytoken"  # Twój token do InfluxDB
    }
}

# Wysłanie zapytania do API Grafany w celu utworzenia źródła danych
response = requests.post(grafana_url, headers=headers, data=json.dumps(data))

# Sprawdzenie odpowiedzi
if response.status_code == 200:
    print("Źródło danych zostało pomyślnie dodane.")
else:
    print(f"Błąd: {response.status_code}, {response.text}")
