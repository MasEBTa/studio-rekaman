# challenge-godb

## Daftar Isi
1. [Judul: studio-music](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/tree/master?ref_type=heads#challenge-godb-studio-music)
2. [Petunjuk Instalasi](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/tree/master?ref_type=heads#installation)
3. [Petunjuk Penggunaan](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/tree/master?ref_type=heads#usage)

    a. [Interaksi Dengan Band](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/tree/master?ref_type=heads#interaksi-dengan-band)

        - Menambah Band

        - Melihat Daftar Band

        - Menghapus band terdaftar

        - Update data band

    b. [Interaksi Dengan Alat Musik](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/tree/master?ref_type=heads#interaksi-dengan-instrumentalat-musik)

        - Menambah Instrument

        - Mengedit Instrument

        - Menghapus Instrument

    c. [Booking Jadwal](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/tree/master?ref_type=heads#booking-jadwal)

        - Booking Jadwal

        - Menampilkan *Booked* Jadwal
4. [Tutorial Penggunaan](https://www.youtube.com/embed/KHfZ8EWLaNs)

## challenge-godb : Studio Music
Projek ini merupakan sebuah aplikasi sederhana dengan golang. aplikasi golang ini digunakan untuk memanipulasi database yang digunakan untuk sebuah studio music yang disewakan. jadi akan ada 2 buah ruangan studio yang di sewakan selama 1 jam setip hari.

## Installation
Untuk menjalankan aplikasi ini :
- [ ] [Buka file connectDb.go pada directory database lalu isikan data database yang digunakan.](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/blob/531673039c6a76a5851b14d97d2647c2ce9ad78b/database/connectDb.go)
    ```
      const (
        host     = "localhost" // default
        port     = 5432        // default
        user     = "postgres"  // drfault
        password = "12345678"
        dbname   = "studio" // sesuaikan
      )
    ```
- [ ] [Jalankan Query yang ada pada file DDL.sql](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/blob/master/DDL.sql)
- [ ] [Jalankan Query yang ada pada file DML-default.sql untuk mengisi data default di database.](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/blob/master/DML-default.sql)
- [ ] Jalankan `go run main.go makeSchedule` untuk membuat jadwal kosong dari studio.

## Usage
Setelah melakukan Instalasi ada beberapa hal yang dapat dilakukan. Berbagai perintah yang ada di bawah ini ada pada file [main](https://git.enigmacamp.com/enigma-20/muhammad-sholikhul-ebta/challenge-godb/-/blob/master/main.go).

### Interaksi dengan Band
Ada beberapa hal yang bisa dilakukan untuk berinteraksi dengan tabel band dan instrument.
1. Menambah Band
    
    Untuk menambah band baru anda bisa menjalankan `go run main.go newBand`. Lalu anda tinggal mengisi data yang dibutuhkan.
2. Melihat daftar Band
    
    Untuk <span id="daftar-band">melihat daftar band</span> yang telah terdaftar di database anda bisa menjalankan `go run main.go daftarBand` lalu anda akan ditampilkan daftar band yang telah terdaftar.
    ```
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================
    Masukkan Id Band untuk melihat daftar Instrument yang digunakan :
    2
    ```
    ketikkan Id dari band untuk melihat daftar instrument(alat musik) yang mereka gunakan.
3. Menghapus band terdaftar

    Untuk menghapus band yang ada di database anda bisa menjalankan `go run main.go deleteBand`
    ```
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================

    Silahkan masukkan Id band yang ingin dihapus :
    4
    ```
    Setelah itu tinggal masukkan Id band yang ingin dihapus.
4. Update data band

    Untu mengupdate data band anda bisa menjalankan `go run main.go updateBand`
    ```
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================

    Silahkan masukkan Id band yang ingin diupdate :
    1
    ```
    Setelah itu masukkan Id band yang ingin di update. Lalu anda tinggal mengisi data baru dari band tersebut. (nb : kosongkan jika tidak ingin dirubah)
### Interaksi dengan instrument/Alat musik
Sebuah band pasti memiliki alat musik. untuk melihat daftar alat musik dari sebuah band anda bisa mengikuti langkah [Melihat daftar Band](#daftar-band). Beberapa hal yang anda bisa lakukan adalah :
1. Menambah Instrument

    Anda bisa menjalankan `go run main.go updateInstruments` lalu anda akan disuguhkan daftar band.
    ```
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================

    Silahkan masukkan Id band Untuk melihat detail Instrument :
    1
    ```
    Masukkan Id band yang anda ingin tambahan instrumennya.

    Setelah itu anda akan dapat melihat daftar instrument yang digunakan oleh band tersebut.
    ```
    Noah, Instrument Yang Digunakan :
    - 1 drum kit
    - 1 guitar
    - 1 bass guitar

    Silahkan masukkan update yang diperlukan (ketik kata kuncinya) :
    - tambah : menambah instrument baru
    - edit : mengubah jumlah instrument yang digunakan
    - delete : menghapus instrument yang telah dipilih

    (tambah/edit/delete) :
    tambah
    ```
    Silahkan input `tambah`, karena anda ingin menambahkan instrument. Lalu akan ada daftar instrument yang bisa digunakan,
    ```
    Instrument yang bisa digunakan :
    key:4. keyboard
    Tuliskan instrumen yang digunakan :
    4
    ```
    Anda bisa masukkan instrument yang ingin di tambahkan. Lalu masukkan jumlah instrument yang di butuhkan :
    ```
    Masukkan jumlah instrumen yang dibutuhkan.
    Jumlah keyboard:
    1
    ```
2. Mengedit Instrument

    Anda bisa menjalankan `go run main.go updateInstruments` lalu anda akan disuguhkan daftar band.
    ```
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================

    Silahkan masukkan Id band Untuk melihat detail Instrument :
    1
    ```
    Masukkan Id band yang anda ingin edit instrumennya.

    Setelah itu anda akan dapat melihat daftar instrument yang digunakan oleh band tersebut.
    ```
    Noah, Instrument Yang Digunakan :
    - 1 drum kit
    - 1 guitar
    - 1 bass guitar
    - 1 keyboard

    Silahkan masukkan update yang diperlukan (ketik kata kuncinya) :
    - tambah : menambah instrument baru
    - edit : mengubah jumlah instrument yang digunakan
    - delete : menghapus instrument yang telah dipilih

    (tambah/edit/delete) :
    edit
    ```
    Silahkan input `edit`, karena anda ingin mengedit instrument. Lalu anda tinggal memasukkan jumlah instrument yang akan digunakan.
    ```
    edit
    [{drum kit 1} {guitar 1} {bass guitar 1} {keyboard 1}]

    Masukkan jumlah instrumen yang dibutuhkan.
    Jumlah drum kit:
    1
    Jumlah guitar:
    2
    Jumlah bass guitar:
    2
    Jumlah keyboard:
    1

    Noah, Instrument Yang Digunakan :
    - 1 drum kit
    - 2 guitar
    - 2 bass guitar
    - 1 keyboard
    ```
3. Menghapus Instrument

    Anda bisa menjalankan `go run main.go updateInstruments` lalu anda akan disuguhkan daftar band.
    ```
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================

    Silahkan masukkan Id band Untuk melihat detail Instrument :
    1
    ```
    Masukkan Id band yang anda ingin hapus instrumennya.

    Setelah itu anda akan dapat melihat daftar instrument yang digunakan oleh band tersebut.
    ```
    Noah, Instrument Yang Digunakan :
    - 1 drum kit
    - 1 guitar
    - 1 bass guitar
    - 1 keyboard

    Silahkan masukkan update yang diperlukan (ketik kata kuncinya) :
    - tambah : menambah instrument baru
    - edit : mengubah jumlah instrument yang digunakan
    - delete : menghapus instrument yang telah dipilih

    (tambah/edit/delete) :
    delete
    ```
    Silahkan input `delete`, karena anda ingin menghapus instrument.
    ```
    Instrument yang digunakan :
    Id:1. Name:drum kit
    Id:2. Name:guitar
    Id:3. Name:bass guitar
    Id:4. Name:keyboard
    Tuliskan instrumen yang ingin duhapus :
    4
    ```
    Masukkan Id instrument yang ingin dihapus.
    ```
    Success Delete Instrument
    =================

    Noah, Instrument Yang Digunakan :
    - 1 drum kit
    - 2 guitar
    - 2 bass guitar
    ```
### Booking Jadwal
Sebuah band yang sudah terdaftar bisa melakukan booking jadwal yang masih kosong.
1. Booking Jadwal

    Untuk booking jadwal anda bisa menjalankan `go run main.go jadwal`, lalu akan ditampilkan hari yang tersedia. anda hanya perlu memasukkan hari yang anda inginkan. (anda tidak perlu mengisi nama hari secara penuh)
    ```
    Hari yang Tersedia.
    - Hari : sabtu
    - Hari : minggu

    Tulis hari yang dinginkan :
    mi
    ```
    Setelah itu akan muncul jadwal yang tersedia 
    ```
    Jadwal yang tersedia :
    2023-08-19 23:55:54.1645984 +0700 +07 m=+5.655055301
    ID | ruangan | Hari | Jam Mulai - Jam selesai
    ==============================================
    79 | room A | minggu | 8:0:0 - 9:0:0
    80 | room A | minggu | 9:5:0 - 10:5:0
    81 | room A | minggu | 10:10:0 - 11:10:0
    82 | room A | minggu | 11:15:0 - 12:15:0
    83 | room A | minggu | 13:0:0 - 14:0:0
    84 | room A | minggu | 14:5:0 - 15:5:0
    85 | room A | minggu | 15:10:0 - 16:10:0
    86 | room A | minggu | 16:15:0 - 17:15:0
    87 | room A | minggu | 17:20:0 - 18:20:0
    88 | room A | minggu | 19:5:0 - 20:5:0
    89 | room A | minggu | 20:10:0 - 21:10:0
    90 | room A | minggu | 21:15:0 - 22:15:0
    91 | room A | minggu | 22:20:0 - 23:20:0
    170 | room B | minggu | 8:0:0 - 9:0:0
    171 | room B | minggu | 9:5:0 - 10:5:0
    172 | room B | minggu | 10:10:0 - 11:10:0
    173 | room B | minggu | 11:15:0 - 12:15:0
    174 | room B | minggu | 13:0:0 - 14:0:0
    175 | room B | minggu | 14:5:0 - 15:5:0
    176 | room B | minggu | 15:10:0 - 16:10:0
    177 | room B | minggu | 16:15:0 - 17:15:0
    178 | room B | minggu | 17:20:0 - 18:20:0
    179 | room B | minggu | 19:5:0 - 20:5:0
    180 | room B | minggu | 20:10:0 - 21:10:0
    181 | room B | minggu | 21:15:0 - 22:15:0
    ==============================================


    Pilih ID jadwal yang diinginkan:
    181
    ```
    Masukkan id jadwal yang diinginkan. Lalu masukkan id band yang akan mem-*booking* jadwal tersebut
    ```
    Pilih ID jadwal yang diinginkan: 181
    Daftar Band yang terdaftar 
    ===========================
    No | Id | nama | email | penanggung jawab
    1. | 1 | Noah | noah@mail.m | Ariel
    2. | 2 | Peterpant | pterw@yj.b | Boril
    3. | 4 | Bagus | erty | qwerty
    ===========================

    Masukkan Id band yang sesuai :
    1
    ```
    Lalu data akan tersimpan di database.
    ```
    Succesfully Choose Schadule data!
    Berhasil claim jadwal.
    ID | ruangan | Hari | Jam Mulai - Jam selesai
    ==============================================
    182 | room B | minggu | 22:20:0 - 23:20:0
    76 | room A | sabtu | 20:10:0 - 21:10:0
    181 | room B | minggu | 21:15:0 - 22:15:0
    ==============================================
    ```
2. Menampilkan *Booked* Jadwal

    Anda dapat melihat jadwal yang telah terisi dengan menjalankan `go run main.go jadwalTerisi`.
    ```
    ID | ruangan | Hari | Jam Mulai - Jam selesai
    ==============================================
    182 | room B | minggu | 22:20:0 - 23:20:0
    76 | room A | sabtu | 20:10:0 - 21:10:0
    181 | room B | minggu | 21:15:0 - 22:15:0
    ==============================================
    ```

### Video Contoh Penggunaan

  Tutorial Penggunaan

  ![Video YouTube](https://www.youtube.com/embed/KHfZ8EWLaNs)
