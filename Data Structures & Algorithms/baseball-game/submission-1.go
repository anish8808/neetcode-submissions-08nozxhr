func calPoints(operations []string) int {
    st := []int{}

    for i :=0; i<len(operations); i++{
        if operations[i] == "+"{
             n := len(st)
             sum := st[n-1]+st[n-2]
            st = append(st , sum)
        }else if operations[i] == "D"{
            st = append(st , 2 * st[len(st)-1])
        }else if operations[i] == "C"{
            st = st[:len(st)-1]
        }else {
            num , _ := strconv.Atoi(operations[i])
            st = append(st , num)
        }
    }

    total := 0

    for _, val := range st{
        fmt.Println(val)
        total +=val
    }

    return total
}
