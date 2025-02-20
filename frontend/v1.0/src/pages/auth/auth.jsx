import { useDispatch, useSelector } from "react-redux"
import { webAuth } from "../../thunk/auth"
import Card from "../../components/card/card"
import { Navigate, useLocation, useNavigate } from "react-router"
import styles from './auth.module.css'
import { useEffect } from "react"

export default function Auth() {
    const dispatch = useDispatch()  
    const location = useLocation()
    const navigate = useNavigate()
    const auth = useSelector((state) => state.auth)
    const from = location?.state?.from?.pathname || "/"
 
    useEffect(() => {
        if (auth.isAuth && !auth.loading) {
            navigate(from)
        }
    }, [auth.isAuth, auth.loading, from, navigate])
 
    const onClick = () => {
        dispatch(webAuth())
    }
 
    return (
        <Card>
            <div className={styles.container}>
                <h2 className={styles.title}>Sign in with Telegram</h2>
                <p className={styles.description}>
                    Connect your Telegram account to access Coinder
                </p>
                <button 
                    className={styles.button} 
                    onClick={onClick}
                    disabled={auth.loading}
                >
                    Sign in with Telegram
                </button>
            </div>
        </Card>
    )
 }