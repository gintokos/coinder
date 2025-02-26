import NNumber from '../number/number'
import { CommentOutlined } from '@ant-design/icons'
import styles from './commentslogo.module.css'



export default function Likes({coin}) {
    const count = 2122
    
    return (
        <div className={styles.container}>
            <button className={styles.button}>
                <CommentOutlined className={styles.comment} />
                <NNumber count={count} />
            </button>
        </div>
    )
}