Project golang my-restaurant-app Project my-restaurant-app adalah aplikasi sederhana yang berfungsi sebagai server backend untuk sebuah restoran. Aplikasi ini ditulis dalam bahasa pemrograman Go dan menggunakan beberapa pustaka, seperti Gorilla Mux, untuk mengelola routing HTTP.

Aplikasi ini memiliki beberapa fitur utama:

    Menu Makanan: Aplikasi menyediakan daftar menu makanan. Setiap item dalam menu memiliki nama, harga dasar, serta topping dan filling (pengisi) yang dapat ditambahkan. Ini memungkinkan pelanggan untuk memilih makanan mereka dan menambahkan variasi, seperti topping dan filling, ke dalam pesanan mereka.

    Pemesanan: Pelanggan dapat membuat pemesanan melalui API POST dengan spesifikasi nama item makanan yang mereka inginkan, serta topping dan filling yang dipilih. Aplikasi akan menghitung harga total pesanan berdasarkan item yang dipilih beserta topping dan filling yang ditambahkan.

    Lihat Menu: Pelanggan juga dapat melihat menu makanan yang tersedia dengan mengakses API GET. Ini memungkinkan mereka untuk melihat daftar menu dan harga dasar untuk setiap item.

    Penghitungan Harga: Aplikasi ini memiliki logika penghitungan harga yang mengambil item makanan yang dipilih, topping, dan filling yang ditambahkan, dan menghitung total harganya.

    Server HTTP: Aplikasi ini berfungsi sebagai server HTTP yang mendengarkan permintaan dari pelanggan. Ini berarti pelanggan dapat terhubung ke aplikasi ini melalui jaringan dan mengirimkan permintaan HTTP, seperti permintaan pemesanan atau melihat menu.

Saat menjalankan aplikasi ini dengan Docker, akan dapat mengaksesnya di dalam wadah Docker, yang berarti dapat menjalankannya di lingkungan yang terisolasi tanpa harus menginstal semua dependensinya secara lokal. Aplikasi ini adalah contoh sederhana tentang bagaimana membuat layanan backend dengan Go dan menggunakan Docker untuk mengemas dan menjalankannya secara konsisten di berbagai lingkungan.

Namun, ada beberapa kendala saat mengembangkan aplikasi ini, terutama pada penggunaan database disistem laptop saya dan juga docker yang mendapat perbedaan user saat melakukan push akibat desktop docker lambat.