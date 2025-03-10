import { useSelector } from 'react-redux'
import Card from '../../components/card/card'
import styles from './main.module.css'
import { Link } from 'react-router'

export default function Main() {
    const auth = useSelector((state) => state.auth)

   return (
       <Card>
           <div className={styles.container}>
               <h1 className={styles.title}>Welcome to Coinder</h1>
               <div className={styles.description}>
                   <p className={styles.text}>
                        Coinder– Your elegant gateway to the world of cryptocurrencies. Discover trending coins with a simple scroll, mark your favorites with a tap, and build your personalized collection of digital assets. Clean, intuitive, and engaging – explore the crypto universe at your fingertips.
                   </p>
               </div>
                {
                    !auth.isAuth && !auth.loading ? (
                        <>
                        <div className={styles.link_container}>
                            <p className={styles.link_text}>
                                To access all features, please sign in
                            </p>
                            <Link to="/auth" className={styles.link}>
                                Sign In
                            </Link>
                        </div>
                        </>
                    ): (
                        <>
                            <div className={styles.link_container}>
                                <p className={styles.link_text}>
                                    Go to Browsing
                                </p>
                                <Link to="/browsing" className={styles.link}>
                                    Browsing
                                </Link>
                            </div>
                        </>
                    )
                }
               <div className={styles.support}>
                   <p className={styles.supportText}>
                       If you like our project and want to support its development
                   </p>
                   <Link to="/support_project" className={styles.supportlink}>
                       Support Project
                   </Link>
               </div>
           </div>
       </Card>
   )
}