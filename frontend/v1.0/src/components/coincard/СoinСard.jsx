import Card from "../../components/card/card"
import styles from "./coinCard.module.css"
import { useState, useRef } from 'react'
import Price from "../price/Price"
import NNumber from "../number/number"
import Modal from "../modal/modal"
import { useModal } from "../../hooks/modal"
import { HeartIcon, CommentsIcon } from "../icons/icons"
import { useDoubleTap } from "../../hooks/double"
import Links from "../links/links"
import { coinderApi } from "../../api/api"

export default function CoinCard({ coin }) {
    console.log("coincard render")

    const urls = {
        ...coin.urls,
        CoinMarcetCap: `https://coinmarketcap.com/currencies/${coin.slug}/`
    }

    const [liked, setLiked] = useState(coin.is_liked)
    const [likeamount, setLikeAmount] = useState(coin.likes_count)
    const heartButtonRef = useRef(null)

    const handleLike = (isDouble) => {
        return () => {
            const button = heartButtonRef.current
            button.classList.add(styles.animate_like) 
            setTimeout(() => button.classList.remove(styles.animate_like), 500)

            if (isDouble) {
                if (!liked) {
                    setLikeAmount(prev => prev + 1) 
                    setLiked(true)
                    coinderApi.like(coin.id)
                }
                return
            } 
            setLikeAmount(prev => liked ? prev - 1 : prev + 1)
            setLiked(prev => !prev)

            if (!liked) {
                coinderApi.like(coin.id)
            } else {
                coinderApi.dislike(coin.id)
            }
        }
    }
    const doubleTapBind = useDoubleTap(handleLike(true))

    const [isModalOpen, openModal, closeModal] = useModal()

    return (
        <Card>
            <div className={styles.clickable} {...doubleTapBind}>
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
                        <Links urls={urls} />
                    </Modal>
                </div>

                <div className={styles.footer}>
                    <button 
                        ref={heartButtonRef}
                        className={styles.icon_button} 
                        onClick={handleLike(false)}
                    >
                        <HeartIcon className={`${styles.icon} ${styles.heart}`} filled={liked} />
                        <NNumber count={likeamount} />
                    </button>
                    {/* <button className={styles.icon_button}>
                        <CommentsIcon className={styles.icon} />
                        <NNumber count={coin.comments_count} />
                    </button> */}
                </div>
            </div>
        </Card>
    )
}