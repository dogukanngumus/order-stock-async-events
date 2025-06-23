# 🐞 BUG TICKET: Sipariş Oluşturulduğunda Stok Güncellenmiyor

## 🛠️ Repro Steps: Hata Durumu

Kullanıcıdan gelen hata raporu:

> "🚨 Sipariş oluşturduğumda, stok güncellenmiyor. Sipariş ekranında işlem başarılı görünüyor ama stok miktarı değişmiyor. Ayrıca bazı siparişler birden fazla kez stoktan düşülmüş gibi görünüyor."

### 🔄 Adımlar:

1. ➡️ Yeni bir sipariş oluştur (ürün ve adet bilgisi gir).  
2. ✅ Sipariş başarılı şekilde onaylanıyor.  
3. ❌ Stok servisi tarafında ürün stoğunun güncellenmediği gözlemleniyor.  
4. 🔁 Aynı siparişi tekrar oluştur veya siparişten kaynaklı stok güncellemesi birkaç kez işlenmiş oluyor.  
5. ⚠️ Sistem bazen mesajları çift işleyebiliyor veya hiç işlemeyebiliyor.

---

## 🎯 Kabul Kriteri:  

- ✅ Sipariş oluşturulduğunda stok servisine ürün bilgisi kesinlikle eksiksiz ve **tek seferlik** iletilecek.  
- ✔️ Stok servisi gelen mesajları **tam ve tek sefer** işleyerek stok bilgisini doğru güncelleyecek.  
- 🔒 Mesajların çift işlenmesi veya kaybolması engellenecek.  
- 🌐 Mesaj iletimi asenkron olacak ve servislerden biri geçici olarak offline olsa bile mesajlar kaybolmayacak, tekrar işlenecek.  
- ⏳ Sistem, mesaj tabanlı iletişimde eventual consistency sağlayacak.

---
