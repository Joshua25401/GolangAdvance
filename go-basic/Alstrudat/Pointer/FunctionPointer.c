#include <stdio.h>

/**
 * Dalam pemrograman bahasa C 
 * ketika mendaklarasikan sebuah variabel ataupun function. hal yang sebenarnya terjadi adalah
 * Kita memesan sejumlah memory pada RAM dan mendapatkan alamat dari variabel ataupun function tersebut
 * Pada file ini akan dibuktikan bahwa function juga memiliki alamat layaknya variabel
 */

 // Disini kita akan deklarasikan function untuk menambah 2 buah integer
 int tambah(int a, int b){
    return a + b;
 }

 int kurang(int a, int b){
    return a - b;
 }

 int operasiRahasia(int a, int (*ptr_to_func)(int,int)){
    return a + ptr_to_func(10,a);
 }

 int operasiRahasia2(int a, int (*ptr_arr_func[])(int,int),int mode){
    switch(mode){
        case 0:
            return a + ptr_arr_func[mode](10,a);
        break;

        case 1:
            return a + ptr_arr_func[mode](10,a);
        break;
    }
    return 0;
 }

int main(){
    
    // Disini kita akan buktikan bahwa function main dan function tambah memiliki alamat
    // Pembuktian ini dilakukan dengan cara melakukan print dengan format %p
    printf("Alamat dari int main() function\t\t: %p\n",main);
    printf("Alamat dari int tambah() function\t: %p\n\n",tambah);

    // Maka dari itu kita juga dapat merujuk kepada alamat dari sebuah function menggunakan pointer
    // Dengan syntax : type (*ptr_name)(params)
    int (*ptr_to_tambah)(int,int);
    ptr_to_tambah = &tambah;
    printf("Hasil dari pemanggilan function melalui pointer\t: %d\n",ptr_to_tambah(10,20));

    // Tidak hanya itu saja, kita juga dapat menjadikan sebuah function menjadi parameter
    // Contohnya seperti pada potongan kode dibawah ini
    printf("Hasil dari operasiRahasia adalah\t: %d\n",operasiRahasia(3,&tambah));
    printf("Hasil dari operasiRahasia adalah\t: %d\n",operasiRahasia(3,&kurang));

    // Kemudian, kita juga dapat menggunakan pointer untuk merujuk kepada banyak function
    // Dengan syntax : type(*ptr_name[])(params)
    // NOTE : Signature dari function juga harus sama mulai dari type dan jumlah parameter
    int (*ptr_arr_func[])(int,int) = {&tambah, &kurang};
    printf("Hasil dari operasiRahasia2 adalah\t: %d\n",operasiRahasia2(3,ptr_arr_func,0));
    printf("Hasil dari operasiRahasia2 adalah\t: %d\n",operasiRahasia2(3,ptr_arr_func,1));
}