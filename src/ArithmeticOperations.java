package src;

/**
 * Created by Timothy Liu on 6/25/17.
 */
public class ArithmeticOperations {
    private Operation operation;

    public ArithmeticOperations(String input) {
        if (input.equals("+")) {
            this.operation = new AdditionOperation();
        } else if (input.equals("-")) {
            this.operation = new SubtractionOperation();
        } else if (input.equals("*")) {
            this.operation = new MultiplicationOperation();
        } else {
            this.operation = new DivisionOperation();
        }
    }

    public int apply(int x, int y) {
        return operation.apply(x, y);
    }

    public double apply(double x, int y) {
        return operation.apply(x, y);
    }

    public double apply(int x, double y) {
        return operation.apply(x, y);
    }

    public double apply(double x, double y) {
        return operation.apply(x, y);
    }

    private abstract class Operation {
        abstract int apply(int x, int y);
        abstract double apply(int x, double y);
        abstract double apply(double x, int y);
        abstract double apply(double x, double y);
    }

    private class AdditionOperation extends Operation {
        int apply(int x, int y) {
            return x + y;
        }

        double apply(int x, double y) {
            return x + y;
        }

        double apply(double x, double y) {
            return x + y;
        }

        double apply(double x, int y) {
            return x + y;
        }
    }

    private class SubtractionOperation extends Operation {
        int apply(int x, int y) {
            return x - y;
        }

        double apply(int x, double y) {
            return x - y;
        }

        double apply(double x, double y) {
            return x - y;
        }

        double apply(double x, int y) {
            return x - y;
        }
    }

    private class MultiplicationOperation extends Operation {
        int apply(int x, int y) {
            return x * y;
        }

        double apply(int x, double y) {
            return x * y;
        }

        double apply(double x, double y) {
            return x * y;
        }

        double apply(double x, int y) {
            return x * y;
        }
    }

    private class DivisionOperation extends Operation {
        int apply(int x, int y) {
            if (y == 0) {
                throw new IllegalArgumentException("Cannot divide by zero");
            }
            return x / y;
        }

        double apply(int x, double y) {
            if (y == 0) {
                throw new IllegalArgumentException("Cannot divide by zero");
            }
            return x / y;
        }

        double apply(double x, double y) {
            if (y == 0) {
                throw new IllegalArgumentException("Cannot divide by zero");
            }
            return x / y;
        }

        double apply(double x, int y) {
            if (y == 0) {
                throw new IllegalArgumentException("Cannot divide by zero");
            }
            return x / y;
        }
    }
}
