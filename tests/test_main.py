import pytest
from self_describing_number.model import (
    get_binomials,
    is_binomial_describing,
    is_enough_binomials
)
from self_describing_number.__main__ import is_valid


@pytest.mark.parametrize(
    "value,expected",
    [
        pytest.param("14233221", True, id="14233221"),
        pytest.param("23322114", True, id="23322114: same as 14233221, but different order"),
        pytest.param("22", True, id="22"),
        pytest.param("23", False, id="23, there is 1 3s, not 2"),
        pytest.param("11", False, id="11: there is 2 1s, not 1"),
        pytest.param("123", False, id="123: not even length"),
        pytest.param("123456", False, id="123456: not self-describing"),
        pytest.param("311018", False, id="311018: not self-describing, we don't count 1s"),
        pytest.param("666666", True,
                     id="666666: only 6s, what do we expect here? should we allow repetitions of the same binomials?"),
        pytest.param("", False, id="empty string"),
    ]
)
def test_is_valid(value: str, expected: bool):
    assert is_valid(value) == expected


@pytest.mark.parametrize(
    "value,expected",
    [
        pytest.param("", [], id="empty string"),
        pytest.param("12", ["12"], id="22"),
        pytest.param("1234", ["12", "34"], id="1234"),
        pytest.param("123456", ["12", "34", "56"], id="123456"),
        pytest.param("123", ["12", "3"], id="123")
    ]
)
def test_get_binomials(value: str, expected: list[str]):
    assert get_binomials(value) == expected


@pytest.mark.parametrize(
    "value,count,figure,expected",
    [
        pytest.param("14233221", 1, "4", True, id="one 4 in 14233221"),
        pytest.param("14233221", 2, "3", True, id="two 3s in 14233221"),
        pytest.param("14233221", 3, "2", True, id="three 2s in 14233221"),
        pytest.param("14233221", 2, "1", True, id="two 1s in 14233221"),
        pytest.param("22", 2, "2", True, id="two 2s in 22"),
        pytest.param("123456", 1, "2", True, id="one 2s in 123456"),
        pytest.param("123456", 3, "4", False, id="three 4s in 123456"),
    ]
)
def test_is_binomial_describing(value: str, count: int, figure: str, expected: bool):
    assert is_binomial_describing(value, count, figure) == expected


@pytest.mark.parametrize(
    "value,expected",
    [
        pytest.param("14233221", True, id="14233221"),
        pytest.param("23322114", True, id="23322114: same as 14233221, but different order"),
        pytest.param("22", True, id="22"),
        pytest.param("23", False, id="23, there is 1 3s, not 2"),
        pytest.param("11", True, id="11: there is 2 1s, not 1"),
        pytest.param("123", False, id="123: not even length"),
        pytest.param("123456", False, id="123456: not self-describing"),
        pytest.param("183110", False, id="311018: not self-describing, we don't count 3s"),
        pytest.param("666666", False, id="666666: only 6s, should be only one binomial"),
        pytest.param("", True, id="empty string"),
    ]
)
def test_is_enough_binomials(value: str, expected: bool):
    # This test ensures that all figures in the number are counted correctly.
    binomials = get_binomials(value)
    assert is_enough_binomials(value, binomials) == expected
