import styles from './coincard.module.css'

export default function CoinCard({ coin, className }) {
    return (
        <div className={`${styles.card} ${className}`}>
            <div className={styles.header}>
                <h3 className={styles.title}>Название монеты {coin}</h3>
                <span className={styles.symbol}>Символ</span>
            </div>
            
            <div className={styles.content}>
                <div className={styles.info}>
                    <p className={styles.price}>Текущая цена: $0.00</p>
                    <p className={styles.price}>Изменение за 24ч: 0%</p>
                </div>
                
                <div className={styles.stats}>
                    <p>Объем: $0.00</p>
                    <p>Капитализация: $0.00</p>
                </div>
            </div>
            
            <div className={styles.footer}>
                <button className={styles.button}>
                    Подробнее
                </button>
            </div>
        </div>
    )
}