package com.mephi;

import java.util.Arrays;

class MyList {
    private int[] arr;
    MyList() {
        this.arr = null;
    }
    MyList(int[] array) {
        int size = array.length;
        this.arr = new int[size];
        System.arraycopy(array, 0, this.arr, 0, size);
    }
    MyList(MyList lt) {
        int size = lt.arr.length;
        this.arr = new int[size];
        System.arraycopy(lt.arr, 0, this.arr, 0, size);
    }
    public void add(int val) {
        int size;
        if (this.arr == null) {
            size = 0;
            this.arr = new int[1];
            this.arr[0] = val;
        } else {
            size =  this.arr.length;
            int[] tmp = new int[size+1];
            System.arraycopy(this.arr, 0, tmp, 0, size);
            tmp[size] = val;
            this.arr = new int[size+1];
            System.arraycopy(tmp, 0, this.arr, 0, size+1);
        }
    }
    public Object remove(int index) {
        int size = this.arr.length;
        int[] tmp = new int[size-1];
        int j = 0;
        Object removed = null;
        if (index < 1) {
            removed = this.arr[0];
            System.arraycopy(this.arr, 1, tmp, 0, size-1);
        } else if (index > size) {
            removed = this.arr[size-1];
            System.arraycopy(this.arr, 0, tmp, 0, size-1);
        } else {
            for (int i = 0; i < size; i++) {
                if (i != index-1) {
                    tmp[j] = this.arr[i];
                    j++;
                } else {
                    removed = this.arr[i];
                }
            }
        }
        this.arr = new int[size-1];
        System.arraycopy(tmp, 0, this.arr, 0, size-1);
        return removed;
    }
    //вставка идет после элемента на который указывает index
    public void add(Object val, int index) {
        int size = this.arr.length;
        Object[] tmp = new Object[size + 1];
        if (index < 0) {
            tmp[0] = val;
            System.arraycopy(this.arr, 0, tmp, 1, size);
        } else if (index > this.arr.length-1) {
            System.arraycopy(this.arr, 0, tmp, 0, size);
            tmp[size] = val;
        } else {
            System.arraycopy(this.arr, 0, tmp, 0, index+1);
            tmp[index] = val;
            System.arraycopy(this.arr, index, tmp, index+2, size-index-1);
        }
        this.arr = new int[size];
        System.arraycopy(tmp, 0, this.arr, 0, size+1);
    }
    public Object get(int index) {
        if (index > this.arr.length-1 || index < 0) {
            return null;
        } else {
            return this.arr[index];
        }
    }
    public int indexOf(int obj) {
        if (this.arr == null) {
            return -1;
        }
        for (int i = 0; i < this.arr.length; i++) {
            if (this.arr[i] == obj) {
                return i;
            }
        }
        return -1;
    }
    public boolean contains(int val) {
        boolean flag = false;
        int t = indexOf(val);
        if (t > -1) {
            flag = true;
        }
        return flag;
    }
    public void set(int val, int index) {
        if (index > this.arr.length-1) {
            this.arr[this.arr.length-1] = val;
        } else if (index < 0) {
            this.arr[0] = val;
        } else {
            this.arr[index] = val;
        }
    }
    public int size() {
        if (this.arr == null) {
            return 0;
        }
        return this.arr.length;
    }
    public boolean isEmpty() {
        return this.arr == null;
    }
}