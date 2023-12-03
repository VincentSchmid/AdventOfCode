package day25

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day25(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(25))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day25(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(25))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
