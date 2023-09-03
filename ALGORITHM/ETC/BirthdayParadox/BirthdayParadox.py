from typing import List
from random import SystemRandom
from collections import Counter


def get_random(day_count, people_count):
    random_class = SystemRandom()
    return [random_class.randint(0, day_count) for _ in range(people_count)]


def find_unique_count(randoms: List[int]):
    pair_counts = 0
    cnt = Counter(randoms)
    for counts in cnt.values():
        if counts > 1:
            pair_counts += 1
    return pair_counts


if __name__ == '__main__':
    has_same = 0
    no_same = 0
    for _ in range(1000000):
        _count = find_unique_count(get_random(365, 23))
        if _count == 0:
            no_same += 1
        else:
            has_same += _count
    print("non count is ", no_same, "and same counts are", has_same, "and probability of same are", (has_same/no_same))

