#include <stdio.h>
#include <string.h>
#include <stdlib.h>

/*
    Pointer to Multidimensional Array :

        Ketika membuat array multi-dimensi yang sebenarnya terjadi adalah 
        Kita membuat sebuah 'pointer' yang menunjuk atau merujuk ke sekumpulan blok alamat pada memory
        Note : Jumlah blok alamat tersebut tergantung pada angka yang kita buat ketika menginisiasi sebuah array
*/

int main(){

    /**
     * Disini dapat dilihat kita menginisiasi sebuah array 2 dimensi dengan nama array
     * yang merujuk pada 2 blok memory dimana masing - masing memory memiliki ukuran 3 blok memory lainnya
     * Dimana ini yang disebut sebagai array multi dimensi 
     */
    int array[2][3] = {{0,1,2},{10,9,8}};

    /**
     * Disini kita mencoba untuk melihat nilai dan alamat dari variabel array 2 dimensi yang bernama array
     * Kita lakukan dengan 2 cara yaitu :
     *      1. Menggunakan de-referensi menggunakan simbol '*'. karena, ini array 2 dimensi maka simbol '*' ditambahkan 2 kali
     *      2. Menggunakan indeks seperti [0][0], [0][1], dst.
     */
    printf("ADDR a[0]\t: %p, VAL\t: %d\t[ENTRY ADDRESS]\n",array,**array);
    printf("ADDR a[0][0]\t: %p, VAL\t: %d\t[ENTRY ADDRESS]\n",&array[0][0],array[0][0]);
    printf("ADDR a[0][1]\t: %p, VAL\t: %d\t[NEXT ADDRESS]\n",&array[0][1],array[0][1]);
    printf("ADDR a[0][2]\t: %p, VAL\t: %d\t[NEXT ADDRESS]\n\n",&array[0][2],array[0][2]);

    int (*ptr_array)[3];
    ptr_array = array;
    printf("ADDR ptr_array\t\t: %p, VAL\t: %d\t[ENTRY ADDRESS]\n",ptr_array,**ptr_array);
    printf("ADDR ptr_array[0][0]\t: %p, VAL\t: %d\t[ENTRY ADDRESS]\n",&ptr_array[0],*ptr_array[0]);
    printf("ADDR ptr_array[0][1]\t: %p, VAL\t: %d\t[NEXT ADDRESS]\n",&ptr_array[0][1],ptr_array[0][1]);
    printf("ADDR ptr_array[0][2]\t: %p, VAL\t: %d\t[NEXT ADDRESS]\n\n",&ptr_array[0][2],ptr_array[0][2]);

    // Loop Through with pointer
    for(int i = 0 ; i < 2 ; ++i){
        for(int j = 0 ; j < 3 ; ++j){
            printf("ADDR ptr_array[%d][%d]\t: %p, VAL\t: %d\n",i,j,&ptr_array[i][j],ptr_array[i][j]);
        }
    }
}