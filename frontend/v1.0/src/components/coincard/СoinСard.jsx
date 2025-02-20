import Card from "../../components/card/card"
import styles from "./coinCard.module.css"
import { useState } from 'react'
import Price from "../price/Price"
import NNumber from "../number/number"
import Modal from "../modal/modal"
import { useModal } from "../../hooks/modal"

import { HeartIcon, CommentsIcon } from "../icons/icons"
import { useDoubleTap } from "../../hooks/double"

export default function CoinCard({ coin }) {
    console.log("coincard render")

    const [liked, setLiked] = useState(coin.isLiked)
    const [likeamount, setLikeAmount] = useState(coin.likes_count)
    const [isAnimating, setIsAnimating] = useState(false)
    const handleLike = () => {
        setLikeAmount(prev => liked ? prev - 1 : prev + 1)
        setLiked(prev => !prev)
        setIsAnimating(true)
        setTimeout(() => setIsAnimating(false), 500)
    }
    const doubleTapBind = useDoubleTap(handleLike)

    const [isModalOpen, openModal, closeModal] = useModal()

    return (
        <Card>
            <div className={styles.clickable} {...doubleTapBind}
             >
                <div className={styles.header}>
                    <div className={styles.titles_container}>
                        <h3 className={styles.slug}>{coin.slug.toUpperCase()}</h3>
                        <h4 className={styles.symbol}>{coin.symbol}</h4>
                    </div>
                </div>

                <div className={styles.price_container}>
                    <Price coin={coin} />
                </div>

                <div className={styles.websites_container}>
                    <h5>Want to know more? check related Websites</h5>
                    <button className={styles.button} onClick={openModal}>
                        Websites
                    </button>
                    <Modal isOpen={isModalOpen} onClose={closeModal}>
                        <h2>Заголовок</h2>
                        <p>Контент модального окна</p>
                    </Modal>
                </div>

                <div className={styles.footer}>
                        <button 
                                className={`${styles.icon_button} ${isAnimating ? styles.animate_like : ''}`} 
                                onClick={handleLike}
                            >
                        <HeartIcon className={`${styles.icon} ${styles.heart}`} filled={liked} />
                        <NNumber count={likeamount} />
                    </button>
                    <button className={styles.icon_button}>
                        <CommentsIcon className={styles.icon} />
                        <NNumber count={coin.comments_count} />
                    </button>
                </div>
            </div>
        </Card>
    )
}