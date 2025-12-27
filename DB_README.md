Singkat dan to the point:
Kamu bisa bikin **user** dan **database** PostgreSQL lewat CLI `psql` atau GUI (misal pgAdmin). Cara paling standar langsung dari terminal.

---

### 1. Masuk ke PostgreSQL

```bash
psql postgres
```

(Ubah `postgres` sesuai superuser/Postgres admin kamu)

---

### 2. Bikin user baru

```sql
CREATE USER namamu WITH PASSWORD 'passwordmu';
```

Contoh:

```sql
CREATE USER rovino WITH PASSWORD 'secret123';
```

---

### 3. Bikin database baru

```sql
CREATE DATABASE namadb OWNER namamu;
```

Contoh:

```sql
CREATE DATABASE tixtrain OWNER rovino;
```

---

### 4. Kasih akses ke user (optional, kalau mau full akses ke DB)

```sql
GRANT ALL PRIVILEGES ON DATABASE tixtrain TO rovino;
```

---

### 5. (Opsional) Cek hasil

* Cek user:

  ```sql
  \du
  ```
* Cek database:

  ```sql
  \l
  ```

---

### 6. **Setting di Go (GORM):**

Isi dsn:

```go
dsn := "host=localhost user=rovino password=secret123 dbname=tixtrain port=5432 sslmode=disable"
```

Pastikan port dan host sesuai (bisa juga remote jika server).

---

**Kesimpulan:**
a. Login psql,
b. CREATE USER,
c. CREATE DATABASE,
d. Grant priv,
e. Pakai kredensial itu di config aplikasi Go/GORM kamu.

Kalau pakai Docker, bisa sekalian pakai environment variable di `docker-compose.yml` untuk setup otomatis.

Butuh script bash otomatis juga? Atau docker-compose? Bisa juga!
