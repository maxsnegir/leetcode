class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:

        word1Map = {}
        word2Map = {}

        for word in word1:
            word1Map[word] = word1Map.get(word, 0) + 1
        for word in word2:
            word2Map[word] = word2Map.get(word, 0) + 1

        if len(word1Map) != len(word2Map):
            return False

        return sorted(word1Map.values()) == sorted(word2Map.values())


def main() -> None:
    pass
