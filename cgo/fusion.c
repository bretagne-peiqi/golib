void memeryarray(int arraya[], int a, int arrayb[], int b, int arrayc[]) {

  int i, j, n;
  i=j=n =0;
  
  while(i < a && j < b) {
  if arraya[i] < arrayb[j] {
     arrayc[n] = arraya[i];
       i++;
  }
  else {
     arrayc[n] = arrayb[j];
       j ++;
  }
  n++;
  }


  while (i < a)
    arrayc[n++] =arraya[i++];
  while (j < b)
    arrayc[n++] =arrayb[j++];
}


void mergearray(int a[], int first, int mid, int last, int temp[])
{
    int i = first, j = mid + 1;
    int m = mid,   n = last;
    int k = 0;

    while (i <= m && j <= n)
    {
        if (a[i] <= a[j])
            temp[k++] = a[i++];
        else
            temp[k++] = a[j++];
    }

    while (i <= m)
        temp[k++] = a[i++];

    while (j <= n)
        temp[k++] = a[j++];

    for (i = 0; i < k; i++)
        a[first + i] = temp[i];
}

void mergesort(int a[], int first, int last, int temp[])
{
    if (first < last)
    {
        int mid = (first + last) / 2;
        mergesort(a, first, mid, temp);    //左边有序
        mergesort(a, mid + 1, last, temp); //右边有序
        mergearray(a, first, mid, last, temp); //再将二个有序数列合并
    }
}

