import { Link } from "react-router"
import Card from "../../components/card/card"
import styles from "./browsing.module.css"
import { useState } from 'react'

export default function Browsing() {
    console.log("Browsing render")
    const [sortOption, setSortOption] = useState('') 
    const [showLiked, setShowLiked] = useState(false)


    return (
        <Card>
            <div className={styles.container}>
                <h2>Search settings</h2>
                
                <div className={styles.sortingSection}>
                    <h3>Sort by:</h3>
                    <label>
                        <input
                            type="radio"
                            name="sortOption" 
                            value="price"
                            checked={sortOption === 'price'}
                            onChange={(e) => setSortOption(e.target.value)}
                        />
                        <span>Price</span>
                    </label>
                    <label>
                        <input
                            type="radio"
                            name="sortOption"
                            value="marketCap"
                            checked={sortOption === 'marketCap'}
                            onChange={(e) => setSortOption(e.target.value)}
                        />
                        <span>Market Cap</span>
                    </label>
                    <label>
                        <input
                            type="radio"
                            name="sortOption"
                            value="popularity"
                            checked={sortOption === 'popularity'}
                            onChange={(e) => setSortOption(e.target.value)}
                        />
                        <span>Popularity</span>
                    </label>
                </div>

                <div className={styles.filterSection}>
                    <label>
                        <input
                            type="checkbox"
                            checked={showLiked}
                            onChange={() => setShowLiked(!showLiked)}
                        />
                        <span>Show liked</span>
                    </label>
                </div>

                <div className={styles.resultSection}>
                    <Link className={styles.showResults} to="/browsing/feed">
                        Show Results
                    </Link>
                </div>
            </div>
        </Card>
    )
}