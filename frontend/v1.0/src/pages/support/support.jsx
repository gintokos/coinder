import { useCallback, useEffect, useState } from "react"
import styles from './support.module.css'
import Card from "../../components/card/card"
import Modal from "../../components/modal/modal"
import { useModal } from "../../hooks/modal"
import { coinderApi } from "../../api/api"


// const isMiniApp = true
export default function Support() {
    const isMiniApp = window?.Telegram?.WebApp?.initData !== undefined && window.Telegram.WebApp.initData !== ""
    const miniAppLink = import.meta.env.VITE_MINI_APP_LINK
    const tg = window.Telegram?.WebApp
    
    const [isOpen, open, close] = useModal()
    const [starAmount, setStarAmount] = useState("")
    const [isLoading, setIsLoading] = useState(false)

    const handleStarAmountChange = (e) => {
        const value = e.target.value.replace(/[^0-9]/g, '')
        setStarAmount(value)
    }

    useEffect(() => {
        if (isMiniApp && tg) {
            const inputElement = document.getElementById('stars');
            if (inputElement) {
                inputElement.autocomplete = 'off';
                inputElement.spellcheck = false;
                inputElement.autocorrect = 'off';
                inputElement.autocapitalize = 'off';
            }
        }
    }, []);


    const isButtonDisabled = !starAmount || parseInt(starAmount) <= 0 || isLoading
    const handleDonate = useCallback(() => {
        if (isButtonDisabled) return;
        
        setIsLoading(true);
        
        coinderApi.createInvoice(starAmount).then((response) => {
            if (response && response.data) {
                tg.openInvoice(response.data);
            }
        }).catch(error => {
            console.error("Error creating invoice:", error);
        }).finally(() => {
            setIsLoading(false);
        });
        
    }, [isButtonDisabled, starAmount, tg, setIsLoading])

    return (
        <>
            <Card>
                <div className={styles.container}>
                    <h2 className={styles.title}>Support Coinder</h2>
                    <p className={styles.description}>
                        Help us make Coinder better by supporting our development. Your contribution will help us add new features and improve existing ones.
                    </p>
                    <div className={styles.telegramStars}>
                        <button onClick={open} className={styles.starsButton}>
                            Support with Telegram Stars
                        </button>
                    </div>
                    <Modal isOpen={isOpen} onClose={close}>
                        {isMiniApp ? (
                            <div className={styles.modal_container}>
                                <h5 className={styles.modal_title}>
                                    Support with Telegram Stars
                                </h5>
                                
                                <div className={styles.input_container}>
                                    <input 
                                        id="stars"
                                        type="text" 
                                        value={starAmount}
                                        onChange={handleStarAmountChange}
                                        className={styles.stars_input}
                                        placeholder=""
                                        inputMode="numeric"
                                        autoComplete="off"
                                        autoCorrect="off"
                                        spellCheck="false"
                                        autoCapitalize="off"
                                    />
                                </div>
                                
                                <button 
                                    className={`${styles.donate_button} ${isButtonDisabled ? styles.donate_button_disabled : ''}`}
                                    onClick={handleDonate}
                                    disabled={isButtonDisabled}
                                >
                                    {isLoading ? 'Processing...' : (
                                        <span className={styles.donate_text}>
                                            Donate {starAmount && <span>{starAmount}</span>} <i className={styles.star_icon}></i>
                                        </span>
                                    )}
                                </button>
                            </div>
                        ) : (
                            <div className={styles.modal_container}>
                                <h5 className={styles.modal_title}>
                                    To support our Project with stars use miniapp
                                </h5>
                                <a href={miniAppLink} className={styles.modal_link} target="_blank">
                                    <span>Open miniapp</span>
                                </a>
                            </div>
                        )}
                    </Modal>
                </div>
            </Card>
        </>
    )
}