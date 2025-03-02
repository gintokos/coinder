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
                       Coinder is a platform that helps you manage and track your cryptocurrency portfolio efficiently. 
                       Monitor your investments, analyze market trends, and make informed decisions all in one place.
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