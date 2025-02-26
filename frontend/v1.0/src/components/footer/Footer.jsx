import styles from './footer.module.css'
import { HouseIcon, TVIcon, ProfileIcon } from '../icons/icons'
import { Link } from 'react-router'
import { useSelector } from 'react-redux'

export default function Footer() {
    const auth = useSelector(state => state.auth)

    return (
        <>
            <div className={styles.container}>
                <Link to="/">
                    <HouseIcon className={`${styles.icon} ${styles.house}`} />
                </Link>
                <Link to="browsing">
                    <TVIcon className={`${styles.icon} ${styles.tv}`} />
                </Link>
                <Link to={`/profile/${auth.id}`}>
                    <ProfileIcon className={`${styles.icon} ${styles.profile}`}/>
                </Link>
            </div>
        </>
    )
}