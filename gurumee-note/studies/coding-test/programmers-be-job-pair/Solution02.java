import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;


public class Solution02 {
    public String solution(String p, int n) {
        String [] times = p.split(" ");
        String amOrPm = times[0];
        String time = times[1];
        LocalDateTime dateTime = LocalDateTime.parse(
                "2021/01/01 " + time,
                DateTimeFormatter.ofPattern("yyyy/MM/dd HH:mm:ss")
        );

        if (amOrPm.equals("AM") && dateTime.getHour() == 12) {
            dateTime = dateTime.minusHours(12);
        } else if (amOrPm.equals("PM") && dateTime.getHour() < 12) {
            dateTime = dateTime.plusHours(12);
        }

        n %= (3600 * 24);
        dateTime = dateTime.plusSeconds(n);
        String answer = dateTime.format(DateTimeFormatter.ofPattern("HH:mm:ss"));
        return answer;
    }
}