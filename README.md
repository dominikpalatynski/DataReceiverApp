# Industrial Monitoring - Instrukcja uruchomienia lokalnie

Ten plik zawiera kompletną instrukcję, jak uruchomić poszczególne części projektu **Industrial Monitoring** na swoim komputerze lokalnym.

---

## 📦 Uruchamianie mikroserwisów

### 1. Mikroserwis `DataReceiver`

Aby uruchomić mikroserwis `DataReceiver` lokalnie:

1. **Utwórz plik `.env`:**  
   W folderze mikroserwisu `DataReceiver` utwórz plik `.env` i uzupełnij go wymaganymi zmiennymi środowiskowymi.

2. **Ustaw zmienną środowiskową:**  
   W terminalu ustaw zmienną środowiskową:
   ```powershell
   $env:DR_DEPLOYMENT_VARIANT = "local"

3. **Uruchom mikroserwis:**  
  przejdź do directory cmd i uruchom aplikację komendą:
   go run .

### 2. Mikroserwis `DeviceManager`

Aby uruchomić mikroserwis `DeviceManager` lokalnie:

1. **Utwórz plik `.env`:**  
   W folderze mikroserwisu `DeviceManager` utwórz plik `.env` i uzupełnij go wymaganymi zmiennymi środowiskowymi.

2. **Ustaw zmienną środowiskową:**  
   W terminalu ustaw zmienną środowiskową:
   ```powershell
   $env:DM_DEPLOYMENT_VARIANT = "local"

3. **Uruchom mikroserwis:**  
  przejdź do directory mikroserwisu i uruchom aplikację komendą:
   go run .

### 3. Docker Compose `local.compose.yaml`

Aby uruchomić `docker compose` lokalnie:

1. **Utwórz plik `data_receiver.env` i `data_receiver.env`:**  
   W folderze głównym projektu utwórz pliki `.env` i uzupełnij go wymaganymi zmiennymi środowiskowymi.

2. **Ustaw zmienną środowiskową:**  
   W terminalu ustaw zmienną środowiskową:
   ```powershell
   $env:DM_DEPLOYMENT_VARIANT = "local"

3. **Uruchom docker compose:**  
  docker compose -f local.compose.yaml up --build -d

### 4. Frontend Setup `IndustrialMonitoringUI`

Aby uruchomić frontend `IndustrialMonitoringUI` lokalnie:

1. **Uruchom komendy `npm install` i `npm rund dev` w directory IndustrialMonitoringUI** 
