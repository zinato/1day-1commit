class Solution {
    public String solution(String new_id) {
        StringBuilder sb = new StringBuilder(new_id);


        for(int i = 0 ; i < sb.length() ; i++) {
            if(sb.charAt(i) >= 65 && sb.charAt(i) <= 90) {
                int c = (int)sb.charAt(i);
                c = c + 32;
                sb.setCharAt(i, (char)c);
            }


            if(sb.charAt(i) != 46 && sb.charAt(i) != 45 && sb.charAt(i) != 95) {
                if(sb.charAt(i) >= 97 && sb.charAt(i) <= 122) {
                    continue;
                }
                if(sb.charAt(i) >= 48 && sb.charAt(i) <= 57) {
                    continue;
                }
                sb.deleteCharAt(i--);
            }            
        }

        for(int i = 0 ; i < sb.length() ; i++) {
            if(sb.charAt(i) == 46) {
                int index = i;
                while(true) {
                    if(++i >= sb.length()) break;
                    if(sb.charAt(i) == 46) {
                        sb.deleteCharAt(i--);
                    } else {
                        break;
                    }
                }
                i = index--;
            }
        }

        if(sb.length() >= 1 && sb.charAt(0) == 46) {
            sb.deleteCharAt(0);
        }
        if(sb.length() >= 1 && sb.charAt(sb.length() - 1) == 46) {
            sb.deleteCharAt(sb.length() - 1);
        }
        
        if(sb.length() == 0) sb = new StringBuilder("a");

        if(sb.length() >= 16) {
            sb = sb.delete(15, sb.length());
            if(sb.charAt(14) == 46) {
                sb = sb.deleteCharAt(14);
            }
        }

        if(sb.length() <= 2) {
            while(true) {
                if(sb.length() == 3) break;
                sb.append(sb.charAt(sb.length() - 1));
            }
        }

        new_id = sb.toString();
        return new_id;
    }
}