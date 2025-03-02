import Card from "../../components/card/card";
import styles from "./profile.module.css";
import { useSelector } from "react-redux";
import { useCallback } from "react";
import { useNavigate } from "react-router";

export default function Profile() {
    const auth = useSelector(state => state.auth)

    const navigate = useNavigate()
    
    const handleLikedBnt = useCallback((islikedToday) => {
        return () => {
            const params = new URLSearchParams()
            params.append('user_id_target', auth.id)
            params.append('liked_by_user', true)
            if (islikedToday) {
                params.append('liked_today', true)
            }

            navigate(`/browsing/feed?${params.toString()}`)
        }
    },[])
    

    return (
        <Card>
            <div className={styles.profileContainer}>
                <div className={styles.userInfo}>
                    <div className={styles.avatarWrapper}>
                        <img 
                            src={auth.photoUrl}
                            alt="User avatar" 
                            className={styles.avatar}
                        />
                    </div>
                    <span className={styles.userName}>
                        { auth.username ? (auth.username) : (auth.firstName) }
                    </span>
                </div>

                <div className={styles.viewSection}>
                    <p className={styles.viewTitle}>
                        View liked posts
                    </p>
                    <p className={styles.viewDescription}>
                        Choose the time period
                    </p>
                    <div className={styles.buttonsContainer}>
                        <button onClick={handleLikedBnt(true)} className={styles.linkButton}>
                            <span className={styles.buttonText}>
                                Liked today
                            </span>
                        </button>
                        
                        <button onClick={handleLikedBnt(false)} className={styles.linkButton}>
                            <span className={styles.buttonText}>
                                Alltime liked
                            </span>
                        </button>
                    </div>
                </div>
            </div>
        </Card>
    )
}