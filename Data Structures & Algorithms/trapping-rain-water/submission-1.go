func trap(height []int) int {
      l, r := 0, len(height)-1
      leftMax, rightMax := 0, 0
      total := 0
      for l < r {
          if height[l] < height[r] {
              if height[l] >= leftMax {
                  leftMax = height[l]
              } else {
                  total += leftMax - height[l]
              }
              l++
          } else {
              if height[r] >= rightMax {
                  rightMax = height[r]
              } else {
                  total += rightMax - height[r]
              }
              r--
          }
      }
      return total
  }

