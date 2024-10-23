# from grafana_api.grafana_face import GrafanaFace

# api_key = "glsa_KDi41RbDhde3KM00UYAEjWQgw3R4Xt63_3de141db"

# grafana = GrafanaFace(auth=api_key, host="localhost", port=3000)

# orgs = grafana.organization.create_organization({
#     "name": "New Organization Name"
# })
# print(orgs)

# import requests

# # URL do Grafany (zmień na swój)
# grafana_url = "http://localhost:3000"

# # Twój token API do autoryzacji (uzyskaj przez UI Grafany)
# headers = {
#     "Authorization": "Bearer glsa_KDi41RbDhde3KM00UYAEjWQgw3R4Xt63_3de141db",
#     "Content-Type": "application/json"
# }

# # Dane nowej organizacji
# payload = {
#     "name": "CompanyB"
# }

# # Wysyłanie żądania do API Grafany w celu utworzenia organizacji
# response = requests.post(f"{grafana_url}/api/orgs", headers=headers, json=payload)

# # Sprawdzenie odpowiedzi
# if response.status_code == 200:
#     print("Organization created successfully")
# else:
#     print(f"Failed to create organization: {response.status_code}, {response.text}")

import requests
from requests.auth import HTTPBasicAuth

# URL do Grafany (zmień na swój)
grafana_url = "http://localhost:3000"

# Dane logowania
username = "admin"  # np. "admin"
password = "password"    # hasło do konta admina

# Dane nowej organizacji
payload = {
    "name": "CompanyB"
}

# Wysyłanie żądania do API Grafany w celu utworzenia organizacji
response = requests.post(f"{grafana_url}/api/orgs", auth=HTTPBasicAuth(username, password), json=payload)

# Sprawdzenie odpowiedzi
if response.status_code == 200:
    print("Organization created successfully")
else:
    print(f"Failed to create organization: {response.status_code}, {response.text}")