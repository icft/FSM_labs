package com.mephi;

import java.io.IOException;
import java.util.*;

public class Main {

    public static MyList merge(MyList l1, MyList l2) {
        MyList mergeList;
        if (l1 == null && l2 == null) {
            mergeList = new MyList();
            return mergeList;
        } else if (l1 == null && l2 != null) {
            mergeList = new MyList(l2);
            return mergeList;
        } else if (l1 != null && l2 == null) {
            mergeList = new MyList(l1);
            return mergeList;
        }
        mergeList = new MyList();
        int s1 = l1.size();
        int s2 = l2.size();
        int i = 0;
        int j = 0;
        while (i < s1 && j < s2) {
            int val1 = (int) l1.get(i);
            int val2 = (int) l2.get(j);
//            System.out.printf("%d %d %d\n", i, val1, val2);
            if (val1 < val2) {
//                System.out.printf("i=%d\n", i);
                mergeList.add(val1);
                i++;
            } else {
//                System.out.printf("j=%d\n", j);
                mergeList.add(val2);
                j++;
            }
        }
//        System.out.printf("%d %d %d\n", i, j, mergeList.size());
        if (i < s1) {
//            System.out.println(1);
            for (int k = i; k < s1; k++) {
                mergeList.add((Integer) l1.get(k));
            }
        }
        if (j < s2) {
//            System.out.println(2);
            for (int k = j; k < s2; k++) {
//                System.out.println(l2.get(k));
                mergeList.add((Integer) l2.get(k));
            }
        }
        return mergeList;
    }

    public static void main(String[] args) {
        int[] a = new int[3];
        a[0] = 1; a[1] = 6; a[2] = 15;
        int[] b = new int[4];
        b[0] = 2; b[1] = 4; b[2] = 8; b[3] = 67;
        MyList l1 = new MyList(a);
        MyList l2 = new MyList(b);
        MyList res = merge(l1, l2);
        for (int i = 0; i < res.size(); i++) {
            System.out.println(res.get(i));
        }
    }
}

