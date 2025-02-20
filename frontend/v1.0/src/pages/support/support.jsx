import Card from "../../components/card/card"
import styles from './support.module.css'

export default function Support() {
   return (
       <>
           <Card>
               <div className={styles.container}>
                   <h2 className={styles.title}>Support Coinder</h2>
                   <p className={styles.description}>
                       Help us make Coinder better by supporting our development. Your contribution will help us add new features and improve existing ones.
                   </p>
                    <div className={styles.telegramStars}>
                        <button className={styles.starsButton}>
                            Support with Telegram Stars
                        </button>
                    </div>
                   
               </div>
           </Card>
       </>
   )
}