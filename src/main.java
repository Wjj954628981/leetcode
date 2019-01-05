import java.util.Arrays;

public class main {
    public static void main(String[] args)
    {
        solution sol = new solution();
        int input[] = {1,2,3};
        int result[] = sol.twoSum(input, 4);
        System.out.println(Arrays.toString(result));
        int input2[] = {1, 2, 3, 4, 2, 3, 4};
        int result2 = sol.kuaizi(input2);
        System.out.println(result2);
    }
}
