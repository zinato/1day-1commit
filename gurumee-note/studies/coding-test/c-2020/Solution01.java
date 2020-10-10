import java.util.Arrays;

public class Solution01 {

    public String get_numbers_kth(int n, int k) {
        String T = "0123456789";
        int div = n / k;
        int mod = n % k;

        if (div == 0) {
            return "" + T.charAt(mod);
        } else {
            return get_numbers_kth(div, k) + T.charAt(mod);
        }
    }
    public int[] solution(int N) {
        int maxK = 0;
        int maxN = 0;

        for (int k=2; k<10; k++) {
            String nth = get_numbers_kth(N, k);
            int result = Arrays.stream(nth.split(""))
                    .map(Integer::parseInt)
                    .filter(i -> i != 0)
                    .reduce(1, (a, b) -> a*b);

            if (maxN <= result) {
                maxN = result;
                maxK = k;
            }
        }


        int[] answer = {maxK, maxN};
        return answer;
    }
}
