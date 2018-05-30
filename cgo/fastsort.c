/***************************************************** 
File name：Quicksort 
Author：Peiqi SHI   Version:1.0    Date: 2016/11/04 
Description: 对数组进行快速排序 
Funcion List: 实现快速排序算法 
*****************************************************/  
  
#include <stdio.h>  
#include <stdlib.h>  
  
#define BUF_SIZE 10  
  
/************************************************** 
 *函数名：display 
 *作用：打印数组元素 
 *参数：array - 打印的数组，maxlen - 数组元素个数 
 *返回值：无 
 **************************************************/  
void display(int array[], int maxlen)  
{  
    int i;  
  
    for(i = 0; i < maxlen; i++)  
    {  
        printf("%-3d", array[i]);  
    }  
    printf("\n");  
  
    return ;  
}


void swap(int *a, int *b) {
   int temp;

   temp = *a;
   *a = *b;
   *b = temp;
}


void quitsort(int array[], begin, last) {

    int i, j;

    if (begin < last){
     i = begin+1;
       j= last;

       while (i < j){
         if (array[i] > array[begin]){
         swap(&array[i], &array[j]);
         j--;
         } else {
         i++;
         }
       }
       
       if (array[i] >= array[begin])
       {i--;}

       swap(&array[i], &array[begin]);

       quitsort(array, begin, i);
       quitsort(array, j, last);

    }
}

int main()
{
    int n;
    int array[BUF_SIZE] = {12,85,25,16,34,23,49,95,17,61};
    int maxlen = BUF_SIZE;

    printf("排序前的数组\n");
    display(array, maxlen);

    quicksort(array, maxlen, 0, maxlen-1);  // 快速排序

    printf("排序后的数组\n");
    display(array, maxlen);

    return 0;
}
