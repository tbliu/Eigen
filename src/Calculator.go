package src

type Calculator struct{}

func NewCalculator() *Calculator {
    return new(Calculator);
}

func Read(query string) string {
    return Transact(query);
}
