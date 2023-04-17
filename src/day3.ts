import fs from 'fs'

const path = __dirname + '/inputs/day3.txt'

const items = ['a', 'b' , 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']

const rucksacks = (path: string): string[] => {
    try{
        const input = fs.readFileSync(path, 'utf-8')
        return input.split('\n')
    } catch(e) {
        console.log(e)
        return []
    }
}

export const list = rucksacks(path)

const getValue = (rucksack: string): number => {
    const mid = rucksack.length / 2
    const side2 = rucksack.slice(mid, rucksack.length)
    for(let i = 0; i < mid; i++) {
        let letter = rucksack[i]
        for(let j = 0; j < mid; j++) {
            if(side2[j] === letter) {
                const value = items.indexOf(side2[j])
                return value + 1
            }
        }
    }
    return -1
}

export const values: number[] = list.map((sack) => {
    return getValue(sack)
})


export const total = values.reduce((acc, curr) => acc + curr)



const getNewValue = (sack1: string, sack2: string, sack3: string): string => {
    for(let i = 0; i < sack1.length; i++) {
        let letter = sack1[i]
        for(let j = 0; j < sack2.length; j++) {
            let letter2 = sack2[j]
            for(let k = 0; k < sack3.length; k++) {
                if(letter === sack2[j] && letter2 === sack3[k]) {
                    return letter
                }
            }
        }
    }
    return ''
}


const getTotal = (list: string[]) => {

    let values = []

    for(let i = 0; i < list.length; i += 3) {
        const letter = getNewValue(list[i], list[i + 1], list[i + 2])
        const value = items.indexOf(letter)
        values.push(value + 1)
    }

    const newTotal = values.reduce((acc, curr) => acc + curr)
    return newTotal
}

export const result = getTotal(list)



