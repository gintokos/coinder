import styles from './number.module.css'
import { Statistic } from 'antd'

export default function NNumber({count}) {
    var suf = ""
    if (count > 1000000000) {
        suf = "B"
        count = Math.round(count/1000000000 * 100) / 100
    } else if (count > 1000000) {
        suf = "M"
        count = Math.round(count/1000000 * 100) / 100
    } else if (count > 1000) {
        suf = "K"
        count = Math.round(count/1000 * 100) / 100
    }
    
    const needsPrecision = !Number.isInteger(count)
    
    return (
        <>
            <Statistic 
                value={count}
                className={styles.number}
                suffix={<span className={styles.suffix}>{suf}</span>}
                precision={needsPrecision ? 2 : 0} 
                valueStyle={{ fontSize: "1.2rem", fontWeight: "bold", color: "var(--accent-purple)" }}
            />
        </>
    )
}