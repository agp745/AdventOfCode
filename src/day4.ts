import fs from 'fs'

const path = __dirname + '/inputs/day4.txt'


const input = (file:string) => {
    try{
        const input = fs.readFileSync(file, 'utf-8')
        return input.split('\n')
    } catch(e) {
        console.log(e)
        return []
    }
}

export const data = input(path)

const compare = (input: string[]) => {
    const removedDashes = input.map((item) => {
        const index = item.indexOf(',')
        const item1 = item.slice(0, index)
        const item2 = item.slice(index + 1, item.length)
        
    })
    return removedDashes
}

export const test = compare(data)


