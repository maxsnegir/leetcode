def equalPairs(grid: list[list[int]]) -> int:
    rows_set = {}
    for row in grid:
        rows_set[tuple(row)] = rows_set.get(tuple(row), 0) + 1

    res = 0
    for column in range(len(grid)):
        columns_set = []

        for row in range(len(grid)):
            columns_set.append(grid[row][column])

        columns_set = tuple(columns_set)
        res += rows_set.get(columns_set, 0)

    return res


if __name__ == "__main__":
    print(equalPairs([[3, 2, 1], [1, 7, 6], [2, 7, 7]]))
    print(equalPairs([[3, 1, 2, 2], [1, 4, 4, 5], [2, 4, 2, 2], [2, 4, 2, 2]]))
