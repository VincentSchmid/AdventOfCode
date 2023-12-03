package day24

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day24(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(24))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day24(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(24))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
