import styles from "./card.module.css"

export default function Card({ children }) {
    return (
        <div className={`${styles.container}`}>
            <div className={`${styles.card}`}>
                {children}
            </div>
        </div>
    )
}