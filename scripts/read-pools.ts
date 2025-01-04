import { readFileSync } from 'fs';

interface Pair {
    nameA: string;
    nameB: string;
    tokenA: string;
    tokenB: string;
}

export interface Pools {
    factoryAddress: string;
    pairs: Pair[];
}

export function readPools<Pools>(filePath: string): Pools {
    try {
        const data = readFileSync(filePath, 'utf-8');
        return JSON.parse(data);
    } catch (error) {
        console.error(`Error while reading json ${filePath}:`, error);
        throw error;
    }
}
