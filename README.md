# Order & Stock Microservices Repository

## 🚀 Projenin Amacı

Bu repository, çok basit mikroservis mimarisi kullanarak **sipariş yönetimi** ve **stok güncelleme** süreçlerini gerçekçi ve sağlam bir şekilde uygulamayı amaçlar.  
Asenkron mesajlaşma altyapısı (RabbitMQ) ile **eventual consistency** sağlanarak, sipariş oluşturma ile stok güncelleme arasında **güvenilir ve tek seferlik iletişim** sağlanmaktadır.

Projede aşağıdaki anahtar kavramlar ve desenler uygulanmıştır:

- **Inbox-Outbox Pattern:** Mesajların çift işlenmesini veya kaybolmasını engellemek için.  
- **Retry Logic:** Hem veritabanı hem de mesaj kuyruğu bağlantıları için bağlanabilirlik garantisi.  
- **Docker Compose:** RabbitMQ, PostgreSQL, Order-Service ve Stock-Service container’ları aynı ağda çalışır.  
- **Golang & .NET:** Stock-Service Go dili ile, Order-Service ise .NET 9 ile geliştirilmiştir.  

---

## 📦 Projeler ve Teknolojiler

| Proje / Teknoloji     | Detaylar                                 |
|----------------------|-----------------------------------------|
| **Order-Service**     | .NET 9, Entity Framework Core, PostgreSQL, RabbitMQ Publisher, Outbox Pattern. |
| **Stock-Service**     | Golang, GORM (PostgreSQL ORM), RabbitMQ Consumer, Inbox Pattern.  |
| **RabbitMQ**          | Mesaj kuyruğu, kuyruk yönetimi, dayanıklılık. |
| **PostgreSQL**        | Kalıcı veri deposu, sipariş ve mesaj kayıtları için.  |
| **Docker & Docker Compose** | Kolay geliştirme ve test ortamı kurulumu.   |

---

## 🐞 BUG TICKET: Sipariş Oluşturulduğunda Stok Güncellenmiyor

### 🛠️ Hata Senaryosu

Kullanıcıdan gelen hata raporu:

> "🚨 Sipariş oluşturduğumda, stok güncellenmiyor. Sipariş başarılı görünmesine rağmen stok değişmiyor. Bazı siparişler stoktan birden fazla kez düşülmüş gibi görünüyor."

#### 🔄 Sorunun Tekrar Edilme Adımları:

1. Yeni bir sipariş oluşturulur.  
2. Sipariş başarıyla onaylanır.  
3. Stok servisi stok güncellemez veya yanlış günceller.  
4. Sipariş birkaç kez stoktan düşülür veya hiç düşülmez.  
5. Mesajlar çift işlenebilir ya da kaybolabilir.

---

## 🎯 Çözüm & Kabul Kriterleri

- Sipariş oluşturulduğunda stok servisine ürün bilgisi **kesinlikle ve yalnızca bir kez** gönderilecek.  
- Stok servisi gelen mesajları **tam ve tek seferlik** işleyerek stok bilgisini doğru güncelleyecek.  
- Mesajların çift işlenmesi ya da kaybolması engellenecek (Inbox-Outbox pattern ile).  
- Mesaj iletimi asenkron ve dayanıklı olacak; servisler offline olsa bile mesajlar kaybolmayacak, tekrar işlenecek.  
- Sistem **eventual consistency** garantisi verecek.

---

## 🧩 Uygulanan Mimari ve Tasarım

- **Order-Service:**  
  - Sipariş oluşturulduğunda Outbox tablosuna mesaj kaydeder.  
  - Arka planda çalışan `OutboxDispatcherBackgroundService` mesajları RabbitMQ’ya gönderir.  
  - Böylece mesajlar veritabanı ile tutarlı ve güvenilir olarak kuyruğa iletilir.

- **Stock-Service:**  
  - RabbitMQ’dan mesajları tüketir.  
  - Her mesajı **Inbox tablosuna yazar ve kontrol eder** (mesaj daha önce işlenmişse atlar).  
  - İşlenmemiş mesajlar için stok güncellemesini yapar.  

- **Retry Mekanizması:**  
  - Hem veritabanı hem de RabbitMQ bağlantıları için, servis başlatılırken bağlantı denemeleri yapılır.  
  - Servisler, mesaj kuyruğu veya veritabanı hazır olana kadar bekler.

- **Docker Compose:**  
  - Tüm servisler ve bağımlılıklar (RabbitMQ, Postgres) tek komutla ayağa kalkar.  
  - Aynı ağda (bridge network) çalışırlar.

---

## 🛠️ Kurulum & Çalıştırma

1. Docker ve Docker Compose yüklü olmalı.  
2. Repository klonlanır.  
3. `docker-compose up --build` komutu ile tüm servisler çalıştırılır.  
4. Order-Service API’sine sipariş oluşturma istekleri gönderilir.  
5. Stock-Service, mesajları alıp stok günceller.

---

## ⚙️ Geliştirme & İyileştirme Önerileri

- **Event Sourcing** ve **CQRS** ile mimarinin genişletilmesi.  
- Mesaj kuyruğunda daha gelişmiş hata yönetimi ve dead-letter kuyruğu kullanımı.  
- Outbox ve Inbox tabloları için arka plan temizleme (pruning) servislerinin eklenmesi.  
- Daha kapsamlı test otomasyonları ve performans testleri.  
- Merkezi loglama ve monitoring entegrasyonu (Prometheus, Grafana, ELK stack).  
- API Gateway ve servis keşif altyapısının eklenmesi.  
- Görev zamanlayıcı ile Outbox mesajlarının zamanlanmış yeniden denemeleri.  

---
