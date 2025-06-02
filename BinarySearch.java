import java.nio.charset.Charset;
import java.util.Arrays;
import java.util.Scanner;

public class BinarySearch {

    public static int rank(int key, int[] a) {
        int lo = 0;
        int hi = a.length - 1;

        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (key < a[mid]) {
                hi = mid - 1;
            } else if (key > a[mid]) {
                lo = mid + 1;
            } else {
                return mid;
            }
        }

        return -1;
    }

    public static void main(String[] args) {
        int[] whitelist = NumList.nums();

        Arrays.sort(whitelist);

        Scanner sc = new Scanner(System.in, Charset.defaultCharset());

        while (!sc.hasNext()) {
            int key = sc.nextInt();

            if (rank(key, whitelist) == -1) {
                System.out.println(key);
            }
        }

        sc.close();
    }
}
