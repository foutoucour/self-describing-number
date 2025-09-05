import os
import time
from concurrent.futures import ThreadPoolExecutor, as_completed

import click
from rich.progress import track
from rich import print
from rich.panel import Panel
from rich.columns import Columns
from .model import (
    is_length_even,
    get_binomials,
    is_binomial_describing,
    is_enough_binomials,
    are_binomials_ordered,
)

timings = {}


def is_valid(number: str) -> bool:
    if not number.isdigit():
        return False

    if not is_length_even(number):
        return False

    binomials = get_binomials(number)
    if len(binomials) != len(set(binomials)):
        return False
    if not is_enough_binomials(number, binomials):
        return False
    if not are_binomials_ordered(binomials):
        return False

    for binomial in binomials:
        if not is_binomial_describing(number, int(binomial[0]), binomial[1]):
            return False

    return True


def process_task(index, task: list[str], progress=False) -> list[str]:
    results = []
    if progress:
        task = track(
            task,
            description=f"Thread {index}...",
            transient=True,
            total=len(task),
            refresh_per_second=10,
            update_period=0.5,
        )
    for i in task:
        if is_valid(str(i)):
            results.append(str(i))
    return results


def round_robin_sublists(l, n=4):
    lists = [[] for _ in range(n)]
    i = 0
    for elem in l:
        lists[i].append(elem)
        i = (i + 1) % n
    return lists


@click.command
@click.argument("number", type=click.INT)
@click.option(
    "number_executor",
    "--number-executor",
    "-n",
    type=click.INT,
    default=os.cpu_count(),
    help="Number of executors to use",
    show_default=True,
)
@click.option(
    "progress",
    "--progress/--no-progress",
    default=False,
    help="Show progress bar",
    show_default=True,
)
def main(number: int, number_executor: int, progress: bool):
    start = time.time()
    if not progress:
        print(f"Using {number_executor} executors... please wait.")
    results = command(number, number_executor, progress)
    end = time.time()

    columns = Columns(results, equal=True, expand=True)
    panel = Panel(
        columns,
        title=f"Self-describing numbers up to {number}",
        subtitle=f"Found {len(results)} self-describing numbers in {end - start:.2f} seconds",
        padding=(1, 2),
        expand=True,
    )
    print(panel)


def command(number: int, number_executor: int, progress_bar: bool = False):
    tasks = round_robin_sublists(range(number), n=number_executor)
    results = []

    with ThreadPoolExecutor(number_executor) as executor:
        # Submit tasks and get Future objects
        futures = [
            executor.submit(process_task, i, n, progress_bar)
            for i, n in enumerate(tasks)
        ]

        for future in as_completed(futures):
            results.extend(future.result())

    results.sort(key=int)
    return results


if __name__ == "__main__":
    main()
