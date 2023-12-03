package day14

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day14(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(14))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day14(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(14))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
