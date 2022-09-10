#include <stdio.h>
#include <stdlib.h>
#include <string.h>


/*
    Pointer to Pointer :
        (RULE 1)> Sebuah pointer dapat menunjukkan sebuah pointer lainnya. Dimana pointer tersebut merujuk kepada sebuah variable.
*/

int main(){
    // (RULE 1) Demonstration
    int variable = 10; // Ini adalah sebuah variabel biasa yang memiliki alamat contoh : 0x33 dan nilai / value = 10
    int *ptr_to_variable = &variable; // Dan ini adalah sebuah pointer ke variable biasa yang sebelumnya di deklarasikan

    // Disini kita akan membuktikan beberapa hal yakni :
    //      1. Alamat dari variable sama dengan nilai yang disimpan oleh pointer_to_variabel
    //      2. Alamat asli dari ptr_to_varible berbeda dengan alamat dari variable untuk membuktikan bahwa mereka adalah 2 individu yang berbeda
    printf("Alamat dari variable\t\t\t: %p\n",&variable); // Disini kita gunakan %p dan menambahkan simbol '&' pada kata variable untuk melihat alamat dari variable
    printf("Alamat yg ada di ptr_to_variable\t: %p\n",ptr_to_variable); // Disini kita akan menampilkan alamat yg disimpan oleh ptr_to_variable
    printf("Alamat asli ptr_to_variable\t\t: %p\n",&ptr_to_variable); // Potongan kode berikut digunakan untuk melihat alamat asli yg dimiliki oleh ptr_to_variable

    // Untuk membuktikan RULE 1 pada kasus ini
    // Kita harus membuat sebuah pointer baru yang akan merujuk pada alamat dari pointer sebelumnya
    // Contohnya dapat dilihat pada potongan kode dibawah ini :
    int **ptr_to_ptr = &ptr_to_variable; // Disini kita membuat sebuah pointer baru yang merujuk pada alamat dari pointer ptr_to_variable
    // Note : Semakin banyak alamat yg dirujuk oleh sebuah pointer maka semakin banyak simbol '*' yg ditambahkan di depan nama variable

    // Kemudian, kita akan buktikan bahwa alamat yang ada pada ptr_to_ptr adalah alamat yang sama dengan alamat asli dari ptr_to_variable
    printf("\n\nPEMBUKTIAN POINTER TO POINTER\n");
    printf("Alamat ptr_to_ptr\t\t: %p\n",ptr_to_ptr);
    printf("Alamat asli ptr_to_variable\t: %p\n",&ptr_to_variable); // Note : Jika ingin mengetahui alamat asli dari sebuah variable. kita cukup menambahkan simbol '&' di depan nama variable

    // Terakhir, untuk membuktikan bahwa praktikum pointer to pointer terpenuhi seutuhnya
    // Kita dapat melakukan pembuktian dengan cara menampilkan "real value" yg dimiliki oleh 3 variable diatas adalah sama
    // Note : value != alamat/reference/pointer
    printf("\n\nPEMBUKTIAN KESAMAAN NILAI DARI 3 VARIABLE\n");
    printf("Value yang dimiliki oleh variable\t: %d\n",variable);
    printf("Value yang dirujuk oleh ptr_to_variable\t: %d\n",*ptr_to_variable);
    printf("Value yang dirujuk oleh ptr_to_ptr\t: %d\n",**ptr_to_ptr);
}