/* <p>
  Write a function that takes in an array of integers and returns the length of
  the longest peak in the array.
</p>
<p>
  A peak is defined as adjacent integers in the array that are <b>strictly</b>
  increasing until they reach a tip (the highest value in the peak), at which
  point they become <b>strictly</b> decreasing. At least three integers are required to
  form a peak.
</p><p>
  For example, the integers <span>1, 4, 10, 2</span> form a peak, but the
  integers <span>4, 0, 10</span> don't and neither do the integers
  <span>1, 2, 2, 0</span>. Similarly, the integers <span>1, 2, 3</span> don't
  form a peak because there aren't any strictly decreasing integers after the
  <span>3</span>.
</p>
*/

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}

	fmt.Println(LongestPeak(arr))
}

func LongestPeak(array []int) int {
currentPeak, longestPeak:=1,0
    var isPeak bool

    for i:=1; i < len(array);i++ {

        if isPeak {
            if array[i] < array[i-1]{
            // a) pico valido y continua caída
                currentPeak+=1
            } else if array[i]> array[i-1] {
             // era pico válido pero se pierde y llevamos 2 de alza   
                currentPeak = 2
                isPeak = false 
            } else {
                // c) cada número debe contar para la suma del peak
                currentPeak = 1
                isPeak= false
            }
        } else {
			if array[i] < array[i-1] && currentPeak > 1 {
				// d) comienza caida
				currentPeak+=1
				isPeak = true
			} else if array[i] > array[i-1] {
                //sigue sin caida, aumentamos valor
				currentPeak +=1
			} else {
                // de ser igual se regresa el score acumulado a 1
				currentPeak = 1
			}
		}	
	
		
		// queremos asegurarnos de guardar en cada iteracion no solo al romper
		if isPeak && currentPeak > longestPeak {
			longestPeak = currentPeak
		}
	}
    
    
    return longestPeak
    
}
