# go-exporter

Ein einfaches Beispiel zum erfassen von Metriken aus go heraus.

## Erstellen des binary

### Windows
Unter Windows 10 die Powershell starten und die Umgebungsvariablen setzen und erstellen:

```Powershell
$env:GOOS="linux"
$env:GOARCH="amd64"

go build  -o bin/go_exporter
```

### MacOS & Linux

```bash
env GOOS=linux GOARCH=amd64 go build  -o bin/go_exporter
```
## Setup

Go exporter auf einer Linux VM starten

```bash
# Benutzer `go_exporter` erstellen:
useradd --no-create-home --shell /bin/false go_exporter

# binary kopieren
scp bin/go_exporter azureuser@1.2.3.4:/tmp/

# service file kopieren
scp bin/go_exporter.service azureuser@1.2.3.4:/tmp/

# auf dem server einloggen
ssh azureuser@1.2.3.4

# Berechtigungen setzen
sudo chown go_exporter:go_exporter

# service file verschieben
sudo mv /tmp/go_exporter.service /etc/systemd/system/go_exporter.service

# daemon neuladen
sudo systemctl daemon-reload

# go_exporter starten
sudo systemctl start go_exporter.service
```

## Prometheus Setup

 Den Exporter in der Konfiguration hinzufügen:

```bash
# login auf prometheus server
ssh azureuser@5.6.7.8

# prometheus config bearbeiten
sudo nano /etc/prometheus/prometheus.yml
```

und den `scrape_configs` erweitern:

```yaml

scrape_configs:
  - job_name: "go_exporter"
    static_configs:
      - targets: ["10.55.0.5:9123"]
```

Hinweis: die IP muss entsprechend der VM auf der der Exporter gestartet wurden angepasst werden.

```bash
# Prometheus neustarten:
sudo systemctl restart prometheus
```

Auf der Prometheus Oberfläche sollte dann unter _targets_ ein weiterer Exporter gelistet sein.