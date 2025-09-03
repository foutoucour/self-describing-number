def is_length_even(number: str) -> bool:
    return len(number) % 2 == 0


def get_binomials(number: str) -> list[str]:
    return [number[i:i + 2] for i in range(0, len(number), 2)]


def is_binomial_describing(number: str, count: int, figure: str) -> True:
    return number.count(figure) == count


def is_enough_binomials(number: str, binomials: list[str]) -> bool:
    unique_figures = set(c for c in number)
    return len(unique_figures) == len(binomials)
