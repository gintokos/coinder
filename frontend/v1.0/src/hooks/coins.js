import { useCallback, useRef } from "react";
import { useLoaderData } from "react-router";
import { coinderApi } from "../api/api";

export default function useCoins({limit}) {
    let [coins, params] = useLoaderData()
    if (coins === null) return [undefined, []]
    console.log('coins: ', coins)
    console.log('params: ', params)

    const index = useRef(0)
    const isAll = useRef(false)

    const nextCoins = useCallback(async () => {
        if (isAll.current) return []
        
        if (index.current >= coins.length) {
            index.current = 0

            params = {
                ...params,
                page: params.page + 1,
            }
            
            try {
                console.log("fetching")
                const response = await coinderApi.coins(params)  
                const fetchedCoins = response.data
                console.log("fetchedcoins: ", fetchedCoins)
                if (!fetchedCoins?.length) {
                    console.log("no more coins")
                    isAll.current = true
                    return []
                }
                coins = fetchedCoins
            } catch (error) {
                console.error('Error fetching coins:', error)
                return []
            }
        }

        const newCoins = []
        while (index.current < coins.length && newCoins.length < limit) {
            newCoins.push(coins[index.current])
            index.current++
        }
        
        return newCoins
    }, [])

    const firstCoins = coins.slice(0, limit)
    index.current = firstCoins.length 
    
    if (firstCoins.length < limit) {
        isAll.current = true
    }
    
    return [nextCoins, firstCoins]
}