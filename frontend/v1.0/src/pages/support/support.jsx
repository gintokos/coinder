import { useCallback } from "react"
import Card from "../../components/card/card"
import Modal from "../../components/modal/modal"
import { useModal } from "../../hooks/modal"
import styles from './support.module.css'

export default function Support() {
    const isMiniApp = window?.Telegram?.WebApp?.initData !== undefined && window.Telegram.WebApp.initData !== ""
    const miniAppLink = import.meta.env.VITE_MINI_APP_LINK
    
    const [isOpen, open, close] = useModal()
    const handleSupport = useCallback(()=>{
        if (isMiniApp) {

        } else {
            open()
        }
    },[])
   return (
       <>
           <Card>
               <div className={styles.container}>
                   <h2 className={styles.title}>Support Coinder</h2>
                   <p className={styles.description}>
                       Help us make Coinder better by supporting our development. Your contribution will help us add new features and improve existing ones.
                   </p>
                    <div className={styles.telegramStars}>
                        <button onClick={handleSupport} className={styles.starsButton}>
                            Support with Telegram Stars
                        </button>
                    </div>
                   <Modal isOpen={isOpen} onClose={close}>
                        <h5 style={{textAlign: 'center', color: 'var(--accent-blue)', fontSize: '2rem', marginTop: '2rem', marginBottom:"1rem"}}>
                            To support our Project with stars use miniapp
                        </h5>
                        <a href={miniAppLink} style={{ textAlign:"center", textDecoration: 'none'}} target="_blank">
                            <span style={{color: 'var(--neutral-100)',}}>Open miniapp</span>
                        </a>
                   </Modal>
               </div>
           </Card>
       </>
   )
}