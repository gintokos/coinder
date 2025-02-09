import styles from './likes.module.css'
import { HeartOutlined, HeartFilled } from '@ant-design/icons'
import { useCallback, useState } from 'react'
import NNumber from '../number/number'


export default function Likes({coin, initLiked}) {
    const [liked, setLiked] = useState(false)
    const [count, setCount] = useState(12000)

    const onClick = useCallback(() => {
        setLiked(!liked)
        setCount((prev)=> {
            if (liked) return prev - 1
            return prev + 1
        })
    }, [liked])
    
    return (
        <div className={styles.container}>
            <button className={styles.button} onClick={onClick}>
                {liked ? 
                        <HeartFilled className={styles.heart}  /> : 
                        <HeartOutlined className={styles.heart} />
                }
                <NNumber count={count} />
            </button>
        </div>
    )
}