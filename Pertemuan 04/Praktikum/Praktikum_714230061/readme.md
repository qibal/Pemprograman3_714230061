# Menjalankan pengujian spesifik
go test -v ./test -run TestUpdateMahasiswa
go test -v ./test -run TestGetMahasiswaByNPM
go test -v ./test -run TestInsertMahasiswa
go test -v ./test -run TestDeleteMahasiswa
go test -v ./test -run TestGetAllMahasiswa

# Menjalankan semua pengujian
go test -v ./...

database in the cloud for free access