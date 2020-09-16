import java.io.IOException;
import java.util.*;


public class Solution02 {
    public static int requestsServed(List<Integer> timestamp, List<Integer> top) {
        PriorityQueue<Integer> tsQueue = new PriorityQueue<>(timestamp);
        Queue<Integer> topQueue = new ArrayDeque<>(top);
        int cnt = 0;

        while (!topQueue.isEmpty()) {
            Integer requestTop = topQueue.remove();
            Stack<Integer> st = new Stack<>();

            while (!tsQueue.isEmpty() && tsQueue.peek() <= requestTop) {
                Integer ts = tsQueue.remove();
                st.push(ts);
            }

            int removeCnt = 0;

            while (!st.isEmpty() && removeCnt < 5) {
                st.pop();
                removeCnt += 1;
                cnt += 1;
            }

            while (!st.isEmpty()) {
                Integer tmp = st.pop();
                tsQueue.add(tmp);
            }

        }

        return cnt;
    }
}