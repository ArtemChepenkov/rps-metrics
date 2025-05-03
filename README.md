# rps-metrics
## Проект, который позволяет запустить клиент и сервер в контейнерах и собирать метрики с них

### Стек:
  * Go
  * Docker
  * Docker compose
  * Grafana
  * Prometheus

### Клиент-сервер
  * В данном проекте схема работы клиента и сервера довольно проста: клиент с определённой частотой (указывается в docker-compose.yaml) посылает запросы, затем сервер делает псевдоработу, и отвечает клиенту

### Контейнеризация
  * Все логические части приложения разбиты в отдельные контейнеры, а именно:
      - server
      - client
      - node-exporter
      - blackbox-exporter
      - prometheus
      - grafana
### Метрики
  * Prometheus собирает метрики с Node-exporter (системные метрики), Blackbox-exporter (метрики HTTP-проверок), а Grafana использует эти данные для построения графиков и дашбордов.

### Имеется возможность посмотреть собранные метрики:
#### Системные метрики (ID 1860, используется Node-exporter)
![image](https://github.com/user-attachments/assets/0b19eaf3-e262-47a6-995c-ea5300093669)
#### HTTP метрики (ID 7587, используется Blackbox-exporter)
![image](https://github.com/user-attachments/assets/1b066b4f-d2c6-4d0d-88a8-ddee9cc60c7d)
