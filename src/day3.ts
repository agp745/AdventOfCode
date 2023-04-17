import fs from 'fs'
import path from 'path'

const filePath = path.join(__dirname, '/inputs/day3.txt')
export const data: string[] = fs.readFileSync(filePath, 'utf8').split('\n')


const findPair = (sack: string): string => {
    const mid: number = (sack.length / 2)
    const side1: string = sack.slice(0, mid)
    const side2: string = sack.slice(mid , sack.length)
    return `${side1} - ${side2}`
}

export const test = findPair(data[0])

const compareSides = (side1: string, side2: string): string => {
    for (let i = 0; i < side1.length; i++) {
        for(let j = 0; j < side2.length; j++) {
            if (side1[i] === side2[j]) {
                return side1[i]
            }
        }
    }
}

export const test2 = compareSides()

