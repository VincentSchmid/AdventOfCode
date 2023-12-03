package day20

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day20(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(20))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day20(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(20))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
