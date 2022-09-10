#include <stdio.h>

struct Operation {
    int data;
    void (*ptr_to_fn_operator[2])(int,int,struct Operation*);
};

enum opr_name {TAMBAH, KURANG};

void tambah(int a, int b, struct Operation* opr){
    opr->data = (a+b);
}

void kurang(int a, int b,struct Operation* opr){
    opr->data = (a-b);
}

struct Operation newOperation(){
    struct Operation temp;

    // Initiate here
    temp.data = 0;
    temp.ptr_to_fn_operator[0] = &tambah;
    temp.ptr_to_fn_operator[1] = &kurang;

    return temp;
}

int main(){
    // Start
    struct Operation operator = newOperation();
    printf("Start\t: %d\n",operator.data);

    // Operate
    operator.ptr_to_fn_operator[TAMBAH](10,20,&operator);

    // Result
    printf("End\t: %d\n",operator.data);
}