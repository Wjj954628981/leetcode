import java.util.*;

import util.*;

public class main {
    public static void main(String[] args)
    {
        solution sol = new solution();

//        twoSumTest(sol);

//        kuaiziTest(sol);

//        powerfulIntegersTest(sol);

//        threeSumTest(sol);

//        maximumSwapTest(sol);

        addTwoNumbersTest(sol);
    }

    private static void twoSumTest(solution sol){
        int input[] = {1,2,3};
        System.out.println(Arrays.toString(sol.twoSum(input, 4)));
    }

    private static void kuaiziTest(solution sol){
        int input2[] = {1, 2, 3, 4, 2, 3, 4};
        System.out.println(sol.kuaizi(input2));
    }

    private static void powerfulIntegersTest(solution sol){
        System.out.println(sol.powerfulIntegers(2, 3, 10));
    }

    private static void threeSumTest(solution sol){
        int input3[] = {-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0};
        System.out.println(sol.threeSum(input3));
    }

    private static void maximumSwapTest(solution sol){
        Interval tmp1 = new Interval(1, 3);
        Interval tmp2 = new Interval(2, 6);
        Interval tmp3 = new Interval(8, 10);
        Interval tmp4 = new Interval(15, 18);
        List<Interval> input4 = new ArrayList<>();
        input4.add(tmp1);
        input4.add(tmp2);
        input4.add(tmp3);
        input4.add(tmp4);
        sol.merge(input4);

        System.out.println(sol.maximumSwap(9973));
    }

    private static void addTwoNumbersTest(solution sol){
//        ListNode l1_leaf = new ListNode(3);
//        ListNode l1_node = new ListNode(4);
//        ListNode l1_root = new ListNode(2);
//
//        ListNode l2_leaf = new ListNode(4);
//        ListNode l2_node = new ListNode(6);
//        ListNode l2_root = new ListNode(5);
//
//        l1_root.next = l1_node;
//        l1_node.next = l1_leaf;
//
//        l2_root.next = l2_node;
//        l2_node.next = l2_leaf;

        ListNode l1_leaf = new ListNode(9);
        ListNode l1_node = new ListNode(9);
        ListNode l1_root = new ListNode(8);
        l1_root.next = l1_node;
        l1_node.next = l1_leaf;

        ListNode l2_root = new ListNode(2);

        ListNode result = sol.addTwoNumbers(l1_root, l2_root);
        System.out.println(result.toString());
    }
}
