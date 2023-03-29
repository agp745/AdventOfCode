import fs from 'fs'

const path = __dirname + '/inputs/day1.txt'

const numbersArr = (path: string) =>{
    try {
        const numbers = fs.readFileSync(path, 'utf8')
        return numbers.split('\n').map((num) => parseInt(num))
    } catch (e) {
        console.log(e)
        return []
    }
}

export const data = numbersArr(path)

let sumsArr: number[] = []
let tempArr: number[] = []

const biggestSum = (data: number[], arr: number[], temp: number[]): number => {
    for(let i = 0; i < data.length; i++) {
    temp.push(data[i])
        if(Number.isNaN(data[i])) {
            temp.pop()
            const sum = temp.reduce((acc, curr) => acc + curr)
            arr.push(sum)
            temp = []
        }
    }
    return sumsArr.reduce((acc, curr) => Math.max(acc, curr))
}

export const biggest = biggestSum(data, sumsArr, tempArr)

const nextBiggestSum = (arr: number[], biggest: number): number => {
    const idx = arr.indexOf(biggest)
    arr.splice(idx, 1)
    return sumsArr.reduce((acc, curr) => Math.max(acc, curr))
}

const second = nextBiggestSum(sumsArr, biggest)
const third = nextBiggestSum(sumsArr, second)

export const top3 = biggest + second + third
