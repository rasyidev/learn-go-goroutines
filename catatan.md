## Concurency dan Paralel Programming
[TODO]

## Membuat Goroutine
- Menambahkan perintah `go` sebelum memanggil function
- Function yang dijalankan dalam goroutine berjalan secara asynchronous
- Aplikasi akan lanjut ke kode program selanjutnya tanpa harus menunggu goroutine yang kita buat selesai.
- Tidak cocok untuk menjalankan function yang mengembalikan nilai (return value)

## Channel
- Tempat komunikasi antar goroutines secara synchronous
- Alat komunikasi synchronous (blocking)
- Ada pengirim dan penerima goroutine
- Saat melakukan pengiriman data ke channel, goroutine akan terblock sampai ada yang menerima data tersebut
- Alternatif async-await pada JavaScript
- Hanya dapat menampung satu data, jika ingin menambahkan data lagi harus gantian dengan data yang ada di channel saat ini.
- Semua goroutine dapat mengambil data dari channel, tapi hanya 1 goroutine untuk 1 data
- Channel harus di-close jika tida digunakan karena dapat menyebabkan memory leak

## Membuat Channel
- chan: Tipe data Channel
- Dibuat menggunakan function `make()`
- Harus menentukan tipe data yang akan disimpan di dalam channel

## Channel sebagai Parameter
- Channel sering diimplementasikan pada parameter function untuk mengirim data menggunakan pointer (pass by reference)
- Channel secara default akan menggunakan pointer, jadi tidak perlu menambahkan `*` atau `&`

## Channel in dan out
- in: menandai bahwa channel dalam suatu function hanya digunakan untuk **menerima** data
- out: menandai bahwa channel dalam suatu function hanya digunakan untuk **mengirim** data
- in: `chan<-`
- out: `<-chan`

## Buffered Channel
- Untuk menampung data antrian di dalam channel
- Biasa digunakan jika lebih dari 1 data yang akan dikirim ke channel tertentu
- Dapat ditentukan jumlah buffer: make(chan <tipedata>, <jmlbuffer>)

## Range Channel
- Jumlah data di dalam channel tidak menentu, sehingga sulit untuk melihat isi datanya
- Dapat diselesaikan dengan **mengulang isi channel selama channel belum di-close**

## Select Channel
- Mendapatkan data dari beberapa channel
- Dapat memilih data tercepat dari beberapa channel
- Jika data datang bersamaan, akan dipilih secara acak

## Default Select
- Mengatasi select channel dalam loop yang tidak ada datanya

## Race Condition
- Variabel yang sama yang diakses secara concurent (menggunakan goroutine) maupun paralel memiliki potensi terkena race condition
- Race Condition, kondisi saat suatu variabel diakses dan diubah secara bersamaan sehingga terdapat data yang hilang

## Mutex (Mutual Exclusion)
- Digunakan untuk mengatasi race condition. 
- Berupa struct `sync.Mutex`
- Dapat digunakan untuk mekanisme lock-unlock variabel yang diakses oleh goroutine
- Hanya satu goroutine dalam satu waktu yang dapat melakukan lock/unlock
- Analogi: Antrian ke toilet, tiap antrian 
  - Masuk, Kunci pintu (lock)
  - Buang hajat (kerjakan job)
  - Buka pintu, keluar (unlock)

## RWMutex (Read Write Mutex)  
- Memiliki dua lock: lock untuk read dan lock untuk write
- Berupa struct `sync.RWMutex`
- Dapat digunakan untuk mekanisme lock-unlock variabel yang diakses oleh goroutine pada saat read dan write
- Sebenarnya dapat menggunakan Mutex, tapi akan rebutan antara proses membaca dan mengubah data

## Deadlock
- Kondisi saat proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan

## WaitGroup
- Fitur yang digunakan untuk menunggu proses goroutine selesai dilakukan
- Sebelumnya masih menggunakan `time.Sleep()` untuk simulasi menunggu proses, hal ini kurang representatif
- Menggunakan WaitGroup lebih representatif
- Berupa struct `sync.WaitGroup`
- **Functions**
  - `Add(n int)`, untuk menambah n antrian
  - `Done()`, untuk mengurangi antrian `Add(-1)`
  - `Wait()`, untuk semua menunggu semua proses sampai selesai dijalankan

## Once
- Fitur untuk memastikan suatu function dijalankan hanya satu kali
- Berupa struct `sync.Once`

## Pool
- Implementasi design pattern **object pool pattern**
- Tempat menyimpan data, jika data sudah digunakan, tempat penyimpanan ini disimpan kembali
- Sering digunakan untuk koneksi dengan database
- Pool sudah aman dari _race condition_
- Berupa struct `sync.Pool` 

## Map
- Mirip dengan tipe data map, tapi Map yang ini sudah bebas dari _race condition_
- Berupa struct `sync.Map`
- **Functions**
  - `Store(key, value)`
  - Load(key)
  - Delete(key)
  - Range(function(key, value))

## Cond
- Mekanisme locking dengan kondisi tertentu
- Membutuhkan locker (Mutex atau RWMutex)
- Terdapat function `Wait()` untuk menunggu semua proses selesai
  - `Signal()`, memberitahu (satu) goroutine untuk berhenti menunggu
  - `Broadcast()`, memberitahu semua goroutine yang tersisa untuk berhenti menunggu
- Untuk membuat cond dapat menggunakan `sync.NewCond()`

## Atomic
- Digunakan untuk menggunakan tipe data primitive secara aman pada proses concurrent (goroutine)
- Dengan menggunakan Atomic, tidak perlu menggunakan locking (Mutex dan RWMutex)

## time.Timer
- Representasi dari kejadian
- Saat timer sudah expire, event dikirim ke dalam channel
- `time.NewTimer(duration)`

## time.After
- Digunakaan saat hanya membutuhkan channel saja, tidak membutuhkan data Timer

## time.AfterFunc
- Digunakan untuk menjalankan function dengan delay tertentu
- Tidak perlu menggunakan channel, cukup kirim function yang akan dipanggil

## time.Ticker
- Representasi kejadian yang berulang
- Saat ticker sudah expire, event dikirim ke dalam channel
- `time.NewTicker(duration)`

## GOMAXPROCS
- Function yang dapat digunakan untuk mengubah atau mengambbil jumlah thread pada komputer yang digunakan
- Secara default, jumlah thread Go Lang sama dengan jumlah CPU komputer
- `runtime.NumCPU()`, menghitung total CPU pada komputer
- `runtime.GOMAXPROCS(n)`, menghitung jumlah thread pada komputer
  - n>=1: Mengubah thread
  - n==-1: menghitung jumlah thread
- `runtime.NumGoroutine()`, menghitung total goroutine yang sedang berjalan pada komputer