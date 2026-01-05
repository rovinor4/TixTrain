# Contoh Response API Find Tickets

## Endpoint
```
GET /tickets/find
```

## Query Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| departure_station | string | Yes | Kode stasiun keberangkatan (misal: GMR, SBY) |
| arrival_station | string | Yes | Kode stasiun tujuan |
| departure_date | string | Yes | Tanggal keberangkatan (format: 2006-01-02) |
| count_passenger | integer | Yes | Jumlah penumpang (minimal 1) |
| filter_class | string | No | Filter berdasarkan class (Ekonomi/Bisnis/Eksekutif) |

---

## Success Response - Dengan Data

**Request:**
```
GET /tickets/find?departure_station=GMR&arrival_station=SBY&departure_date=2026-01-05&count_passenger=2
```

**Response:** `200 OK`
```json
{
  "data": [
    {
      "schedule_group_id": 1,
      "schedule_group_name": "Argo Bromo Anggrek - Route 1",
      "train_name": "Argo Bromo Anggrek",
      "train_code": "ABA",
      "departure_station": "Gambir",
      "arrival_station": "Surabaya Pasar Turi",
      "departure_time": "08:00",
      "arrival_time": "16:30",
      "classes": [
        {
          "class": "Ekonomi",
          "price": 50000,
          "available_seats": 45
        },
        {
          "class": "Bisnis",
          "price": 100000,
          "available_seats": 30
        },
        {
          "class": "Eksekutif",
          "price": 150000,
          "available_seats": 20
        }
      ]
    },
    {
      "schedule_group_id": 3,
      "schedule_group_name": "Bima - Route 2",
      "train_name": "Bima",
      "train_code": "BIM",
      "departure_station": "Gambir",
      "arrival_station": "Surabaya Pasar Turi",
      "departure_time": "14:00",
      "arrival_time": "22:45",
      "classes": [
        {
          "class": "Ekonomi",
          "price": 50000,
          "available_seats": 38
        },
        {
          "class": "Bisnis",
          "price": 100000,
          "available_seats": 25
        }
      ]
    }
  ]
}
```

---

## Success Response - Dengan Filter Class

**Request:**
```
GET /tickets/find?departure_station=GMR&arrival_station=SBY&departure_date=2026-01-05&count_passenger=2&filter_class=Eksekutif
```

**Response:** `200 OK`
```json
{
  "data": [
    {
      "schedule_group_id": 1,
      "schedule_group_name": "Argo Bromo Anggrek - Route 1",
      "train_name": "Argo Bromo Anggrek",
      "train_code": "ABA",
      "departure_station": "Gambir",
      "arrival_station": "Surabaya Pasar Turi",
      "departure_time": "08:00",
      "arrival_time": "16:30",
      "classes": [
        {
          "class": "Eksekutif",
          "price": 150000,
          "available_seats": 20
        }
      ]
    }
  ]
}
```

---

## Success Response - Tidak Ada Data

**Request:**
```
GET /tickets/find?departure_station=GMR&arrival_station=BDG&departure_date=2026-01-05&count_passenger=2
```

**Response:** `200 OK`
```json
{
  "data": []
}
```

---

## Error Response - Validation Error

**Request:**
```
GET /tickets/find?departure_station=GMR&count_passenger=2
```

**Response:** `400 Bad Request`
```json
{
  "errors": {
    "arrival_station": "ArrivalStation wajib di isi",
    "departure_date": "DepartureDate wajib di isi"
  }
}
```

---

## Error Response - Invalid Date Format

**Request:**
```
GET /tickets/find?departure_station=GMR&arrival_station=SBY&departure_date=05-01-2026&count_passenger=2
```

**Response:** `400 Bad Request`
```json
{
  "errors": {
    "departure_date": "DepartureDate harus berupa tanggal yang valid dengan format 2006-01-02"
  }
}
```

---

## Error Response - Invalid Count Passenger

**Request:**
```
GET /tickets/find?departure_station=GMR&arrival_station=SBY&departure_date=2026-01-05&count_passenger=0
```

**Response:** `400 Bad Request`
```json
{
  "errors": {
    "count_passenger": "CountPassenger minimal 1"
  }
}
```

---

## Penjelasan Response Field

### TicketResponse
- `schedule_group_id`: ID dari schedule group
- `schedule_group_name`: Nama rute kereta
- `train_name`: Nama kereta
- `train_code`: Kode kereta
- `departure_station`: Nama stasiun keberangkatan
- `arrival_station`: Nama stasiun tujuan
- `departure_time`: Waktu keberangkatan (format: HH:mm)
- `arrival_time`: Waktu tiba (format: HH:mm)
- `classes`: Array dari class yang tersedia

### TicketClassInfo
- `class`: Nama class (Ekonomi/Bisnis/Eksekutif)
- `price`: Harga per penumpang dalam Rupiah
- `available_seats`: Jumlah kursi yang masih tersedia

---

## Catatan
- Response hanya menampilkan jadwal yang jadwal keberangkatannya belum lewat
- Response hanya menampilkan class yang available seats-nya >= count_passenger
- Jika menggunakan `filter_class`, hanya class yang dipilih yang akan ditampilkan
- Price diambil dari database (field `price` di tabel `coaches`)
