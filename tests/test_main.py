import os

import pytest
from self_describing_number.__main__ import is_valid, command


@pytest.mark.parametrize(
    "value,expected",
    [
        pytest.param("14233221", True, id="14233221"),
        pytest.param(
            "23322114", False, id="23322114: same as 14233221, but different order"
        ),
        pytest.param("22", True, id="22"),
        pytest.param("23", False, id="23, there is 1 3s, not 2"),
        pytest.param("11", False, id="11: there is 2 1s, not 1"),
        pytest.param("123", False, id="123: not even length"),
        pytest.param("123456", False, id="123456: not self-describing"),
        pytest.param(
            "311018", False, id="311018: not self-describing, we don't count 1s"
        ),
        pytest.param("666666", False, id="666666: not self-describing"),
        pytest.param("", False, id="empty string"),
    ],
)
def test_is_valid(value: str, expected: bool):
    assert is_valid(value) == expected


@pytest.mark.parametrize(
    "value,expected",
    [
        pytest.param(0, [], id="0"),  # 0ms
        pytest.param(10, [], id="10"),  # 0ms
        pytest.param(100, ["22"], id="100"),  # 0ms
        pytest.param(1000, ["22"], id="1000"),  # 0ms
        pytest.param(10000, ["22"], id="10000"),  # 8ms
        pytest.param(100000, ["22"], id="100000"),  # 23ms
        pytest.param(1000000, ["22"], id="1000000"),  # 1s
        pytest.param(10000000, ["22"], id="10000000"),  # 3s
        pytest.param(
            100000000,
            [
                "22",
                "14233221",
                "14331231",
                "14333110",
                "15143331",
                "15233221",
                "15331231",
                "15333110",
                "16143331",
                "16153331",
                "16233221",
                "16331231",
                "16333110",
                "17143331",
                "17153331",
                "17163331",
                "17233221",
                "17331231",
                "17333110",
                "18143331",
                "18153331",
                "18163331",
                "18173331",
                "18233221",
                "18331231",
                "18333110",
                "19143331",
                "19153331",
                "19163331",
                "19173331",
                "19183331",
                "19233221",
                "19331231",
                "19333110",
                "23322110",
                "33123110",
            ],
            id="100000000",
        ),  # 2mins
    ],
)
def test_command(value: int, expected: list[str]):
    assert command(value, os.cpu_count(), False) == expected
