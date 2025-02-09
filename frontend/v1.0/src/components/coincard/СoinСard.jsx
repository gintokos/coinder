import { useState } from 'react';
import styles from './coincard.module.css';
import Price from '../price/Price';
import Likes from '../likes/likes';
import CommentsLogo from '../commentslogo/commentslogo';
import FavoriteIcon from '../favorite/favorite';
import ShareIcon from '../share/share';
import Urls from '../urls/urls';
import { Statistic } from 'antd';


export default function CoinCard({ coin, className }) {
    return (
        <div className={`${styles.container} ${className}`}>
            {coin}
            <div className={`${styles.card}`}>
                <div className={styles.header}>
                    <h2 className={styles.slug}>Bitcoin</h2>
                    <h3 className={styles.symbol}>BTC</h3>
                </div>
                <div className={styles.content}>
                    <Price coin={coin} />
                    <div className={styles.container2}>
                        <div className={styles.metacontainer}>
                            <span className={styles.metatitle}>
                                Volume24h
                            </span>
                            <div className={styles.value}>
                                <Statistic
                                    value={123518713251}
                                    suffix= {<span className={styles.usd}>$</span>}
                                    valueStyle={{
                                        color: "var(--accent-primary)"
                                    }}
                                    />
                            </div>
                            <span className={styles.metatitle}>
                                Volumechange24h
                            </span>
                            <div className={styles.value}>
                                <Statistic
                                    value={31243125}
                                    suffix= {<span className={styles.usd}>$</span>}
                                    valueStyle={{
                                        color: "var(--accent-primary)"
                                    }}
                                    />
                            </div>
                            <span className={styles.metatitle}>
                                Marcetcap
                            </span>
                            <div className={`${styles.value} ${styles.last}`}>
                                <Statistic
                                    value={94871239867}
                                    suffix= {<span className={styles.usd}>$</span>}
                                    valueStyle={{
                                        color: "var(--accent-primary)"
                                    }}
                                    />
                            </div>
                            <div className={styles.urls}></div>
                            <Urls />
                        </div>
                    </div>
                    <div className={styles.sidemenu}>
                        <Likes id={coin} initLiked={false} />
                        <CommentsLogo coin={coin} />
                        <FavoriteIcon style={{ fontSize: "2rem" }} />
                        <ShareIcon style={{ fontSize: "2rem" }} />
                    </div>
                </div>
            </div>
        </div>
    );
}