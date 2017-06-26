package tests;
import src.ArithmeticOperations;
import org.junit.Test;
import static org.junit.Assert.*;

/**
 * Created by Timothy Liu on 6/25/17.
 */
public class ArithmeticOperationsTests {
    @Test
    public void testAddition() {
        ArithmeticOperations a = new ArithmeticOperations("+");
        assertEquals(7, a.apply(3, 4));
        assertEquals(3.5, a.apply(3, 0.5), 0.1);
        assertEquals(10.0, a.apply(6.0, 4), 0.1);
        assertEquals(10, a.apply(6.4, 3.6), 0.1);
    }

    @Test
    public void testSubtraction() {
        ArithmeticOperations a = new ArithmeticOperations("-");
        assertEquals(12, a.apply(24, 12));
        assertEquals(3, a.apply(3.0, 0), 0.1);
        assertEquals(5, a.apply(7.3, 2.3), 0.1);
        assertEquals(10.5, a.apply(13.0, 2.5), 0.1);
    }

    @Test
    public void testMultiplication() {
        ArithmeticOperations a = new ArithmeticOperations("*");
        assertEquals(25, a.apply(5, 5));
        assertEquals(1.5, a.apply(0.5, 3), 0.1);
        assertEquals(1.5, a.apply(3, 0.5), 0.1);
        assertEquals(0, a.apply(40, 0));
    }

    @Test
    public void testDivision() {
        ArithmeticOperations a = new ArithmeticOperations("/");
        assertEquals(1, a.apply(5, 5));
        assertEquals(0.5, a.apply(2.5, 5), 0.1);
        assertEquals(2, a.apply(5, 2.5), 0.1);
        try {
            a.apply(40 ,0);
            assertTrue(false);
        } catch (IllegalArgumentException e) {
            assertTrue(true);
        }
    }
}
