import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class main {
    public static void main(String[] args)
    {
        solution sol = new solution();
        int input[] = {1,2,3};
        System.out.println(Arrays.toString(sol.twoSum(input, 4)));

        int input2[] = {1, 2, 3, 4, 2, 3, 4};
        System.out.println(sol.kuaizi(input2));

        System.out.println(sol.powerfulIntegers(2, 3, 10));

        int input3[] = {-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0};
        System.out.println(sol.threeSum(input3));

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
}
